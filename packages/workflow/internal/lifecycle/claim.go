package lifecycle

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"os"
	"path"
	"reflect"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/model"
	"gopkg.in/yaml.v3"
)

type claim struct {
	Worktree   string         `json:"worktree"`
	Slug       string         `json:"change"`
	Path       string         `json:"path"`
	Base       string         `json:"base"`
	Approved   string         `json:"approved"`
	Projection snapshotRecord `json:"projection"`
}

func encodeClaim(root string, c claim) (string, error) {
	data, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return gitstore.WriteBlob(root, data)
}

func readClaim(root string, r ref) (claim, error) {
	var c claim
	data, err := gitstore.ReadBlob(root, r.oid)
	if err == nil {
		err = json.Unmarshal(data, &c)
	}
	if err == nil && (c.Worktree == "" || c.Slug == "" || c.Slug != strings.TrimPrefix(r.name, claimsPrefix)) {
		err = errors.New("invalid claim identity")
	}
	if err != nil {
		return c, coded("P02", r.name, fmt.Errorf("read claim %s: %w", r.name, err))
	}
	return c, nil
}

// YAML preserves the types of configuration and document values. JSON
// preserves the typed module fields, including absent and empty zone lists.
type snapshotRecord struct {
	Config  string
	Modules []moduleRecord
}

type moduleRecord struct {
	Module      project.Module
	Frontmatter string
}

func recordSnapshot(snapshot project.Snapshot) (snapshotRecord, error) {
	config, err := yaml.Marshal(snapshot.Config)
	if err != nil {
		return snapshotRecord{}, err
	}
	record := snapshotRecord{Config: string(config), Modules: make([]moduleRecord, 0, len(snapshot.Modules))}
	for _, module := range snapshot.Modules {
		fields, err := yaml.Marshal(module.Document.Frontmatter)
		if err != nil {
			return record, err
		}
		module.Document.Frontmatter = nil
		record.Modules = append(record.Modules, moduleRecord{Module: module, Frontmatter: string(fields)})
	}
	return record, nil
}

func (c claim) snapshot(root string) (project.Snapshot, error) {
	snapshot := project.Snapshot{Root: root}
	if len(c.Projection.Modules) == 0 || c.Base == "" ||
		changeIdentity(c.Path).Slug != c.Slug || path.Dir(c.Path) != pendingDir {
		return snapshot, failure("P02", c.Slug, "invalid approved snapshot or phase base")
	}
	if err := yaml.Unmarshal([]byte(c.Projection.Config), &snapshot.Config); err != nil {
		return snapshot, coded("P02", c.Slug, err)
	}
	if snapshot.Config == nil {
		return snapshot, failure("P02", c.Slug, "missing approved configuration")
	}
	for _, record := range c.Projection.Modules {
		module := record.Module
		if err := yaml.Unmarshal([]byte(record.Frontmatter), &module.Document.Frontmatter); err != nil {
			return snapshot, coded("P02", c.Slug, err)
		}
		snapshot.Modules = append(snapshot.Modules, module)
	}
	return snapshot, nil
}

func (c claim) approvedDocument() (contracts.Document, error) {
	doc, err := contracts.ParseDocument(c.Approved)
	if err != nil {
		return doc, coded("P02", c.Slug, fmt.Errorf("read approved change: %w", err))
	}
	return doc, nil
}

func ownedClaim(root string) (claim, string, error) {
	all, err := refs(root, claimsPrefix)
	if err != nil {
		return claim{}, "", coded("P02", "", err)
	}
	for _, r := range all {
		c, err := readClaim(root, r)
		if err != nil {
			return c, "", err
		}
		if c.Worktree == root {
			return c, r.oid, nil
		}
	}
	return claim{}, "", nil
}

func Approved(root string) (model.ApprovedSnapshot, bool, error) {
	repo, err := gitstore.Inspect(root)
	if err != nil {
		return model.ApprovedSnapshot{}, false, err
	}
	c, oid, err := ownedClaim(repo.Root)
	if err != nil || oid == "" {
		return model.ApprovedSnapshot{}, false, err
	}
	snapshot, err := c.snapshot(repo.Root)
	if err == nil {
		_, err = c.approvedDocument()
	}
	if err != nil {
		return model.ApprovedSnapshot{}, false, err
	}
	return model.ApprovedSnapshot{Snapshot: snapshot, Slug: c.Slug, Path: c.Path}, true, nil
}

func loadHeld(root, selector string) (model.State, claim, error) {
	var state model.State
	var held claim
	repo, err := gitstore.Inspect(root)
	if err != nil {
		return state, held, err
	}
	if selector == "" {
		held, state.ClaimOID, err = ownedClaim(repo.Root)
		if err != nil {
			return state, held, err
		}
		if state.ClaimOID == "" {
			return state, held, failure("P02", "", "worktree holds no change")
		}
	} else {
		paths, err := changePaths(repo.Root, "", pendingDir)
		if err != nil {
			return state, held, err
		}
		ch, err := selectChange(paths, selector)
		if err != nil {
			return state, held, err
		}
		name := claimsPrefix + ch.Slug
		oid, found, err := gitstore.Ref(repo.Root, name)
		if err != nil {
			return state, held, coded("P02", ch.Slug, err)
		}
		if !found {
			return state, held, failure("P02", ch.Slug, "change has no claim: %s", ch.Slug)
		}
		held, err = readClaim(repo.Root, ref{name, oid})
		if err != nil {
			return state, held, err
		}
		if held.Worktree != repo.Root {
			return state, held, failure("P02", ch.Slug, "claim belongs to another worktree: %s", held.Worktree)
		}
		if held.Path != ch.Path {
			return state, held, failure("P02", ch.Slug, "claim belongs to a different change path: %s", held.Path)
		}
		state.ClaimOID = oid
	}
	state.Head, state.PhaseBase = repo.Head, held.Base
	state.Snapshot, err = held.snapshot(repo.Root)
	if err != nil {
		return state, held, err
	}
	approved, err := held.approvedDocument()
	if err != nil {
		return state, held, err
	}
	state.Run, state.Phase = str(approved.Frontmatter["run"]), str(approved.Frontmatter["phase"])
	state.Change, err = readChange(repo.Root, "", changeIdentity(held.Path))
	if err != nil {
		return state, held, coded("P03", held.Path, fmt.Errorf("read active change: %w", err))
	}
	return state, held, nil
}

func Current(root, selector string) (model.State, error) {
	state, _, err := loadHeld(root, selector)
	return state, err
}

func Release(root, selector string) error {
	repo, err := gitstore.Inspect(root)
	if err != nil {
		return err
	}
	if selector == "" {
		held, oid, err := ownedClaim(repo.Root)
		if err != nil {
			return err
		}
		if oid == "" {
			return failure("P02", "", "worktree holds no change")
		}
		return releaseClaim(repo.Root, held.Slug, oid)
	}
	paths, err := changePaths(repo.Root, "", pendingDir)
	if err != nil {
		return err
	}
	ch, err := selectChange(paths, selector)
	if err != nil {
		return err
	}
	name := claimsPrefix + ch.Slug
	oid, found, err := gitstore.Ref(repo.Root, name)
	if err != nil {
		return coded("P02", ch.Slug, err)
	}
	if !found {
		return failure("P02", ch.Slug, "change has no claim: %s", ch.Slug)
	}
	held, err := readClaim(repo.Root, ref{name, oid})
	if err != nil {
		return err
	}
	if held.Worktree != repo.Root || held.Path != ch.Path {
		return failure("P02", ch.Slug, "claim belongs to another worktree or change")
	}
	return releaseClaim(repo.Root, ch.Slug, oid)
}

func releaseClaim(root, slug, oid string) error {
	if err := gitstore.CASRef(root, claimsPrefix+slug, oid, ""); err != nil {
		return coded("P02", slug, err)
	}
	return nil
}

func queuedStatus(root string, queue []project.Change) ([]model.StatusEntry, error) {
	var blocked []project.Change
	result := make([]model.StatusEntry, 0, len(queue))
	for _, change := range queue {
		fields := change.Document.Frontmatter
		entry := model.StatusEntry{Slug: change.Slug, Path: change.Path, Module: change.Module,
			Phase: str(fields["phase"]), Blocked: str(fields["blocked"]), AuditRounds: integer(fields["audit_rounds"], 0)}
		if entry.Phase == "" {
			entry.Phase = "queued"
		}
		if entry.Blocked != "" {
			blocked = append(blocked, change)
		} else if len(blocked) != 0 {
			snapshot, err := project.Project(root, change.Path)
			if err != nil {
				return nil, err
			}
			for _, earlier := range blocked {
				switch {
				case earlier.Module == change.Module:
					entry.WaitingRule = "same_module"
				case dependsOn(snapshot, change.Module, earlier.Module, true):
					entry.WaitingRule = "depends_on_blocked"
				case strings.TrimSpace(section(change.Document, "Interface")) != "" &&
					dependsOn(snapshot, earlier.Module, change.Module, true):
					entry.WaitingRule = "interface_dependency"
				}
				if entry.WaitingRule != "" {
					entry.WaitingOn = numberedName(earlier)
					break
				}
			}
		}
		result = append(result, entry)
	}
	return result, nil
}

func Eligible(root string) ([]project.Change, error) {
	queue, err := project.Queue(root)
	if err != nil {
		return nil, err
	}
	return eligibleChanges(root, queue)
}

func eligibleChanges(root string, queue []project.Change) ([]project.Change, error) {
	entries, err := queuedStatus(root, queue)
	if err != nil {
		return nil, err
	}
	var result []project.Change
	for i, entry := range entries {
		if entry.Blocked == "" && entry.WaitingOn == "" {
			result = append(result, queue[i])
		}
	}
	return result, nil
}

func Start(root, selector string) (model.State, error) { return start(root, selector, 0) }

func StartFromAttempt(root, selector string, attempt int) (model.State, error) {
	if attempt <= 0 {
		return model.State{}, failure("P01", selector, "attempt must be positive: %d", attempt)
	}
	return start(root, selector, attempt)
}

func start(root, selector string, attempt int) (model.State, error) {
	var state model.State
	repo, err := gitstore.Inspect(root)
	if err != nil {
		return state, err
	}
	root = repo.Root
	if err := clean(root, "start"); err != nil {
		return state, err
	}
	if selector != "" {
		paths, err := changePaths(root, "", pendingDir)
		if err != nil {
			return state, err
		}
		if _, err := selectChange(paths, selector); err != nil {
			return state, err
		}
	}
	queue, err := project.Queue(root)
	if err != nil {
		return state, err
	}
	eligible, err := eligibleChanges(root, queue)
	if err != nil {
		return state, err
	}
	if len(eligible) == 0 {
		return state, failure("P01", selector, "no eligible change")
	}
	change := eligible[0]
	if selector != "" && selector != change.Path && selector != path.Base(change.Path) && selector != change.Slug {
		return state, failure("P01", selector, "change is not eligible queue head: %s", selector)
	}

	state.Head, state.PhaseBase, state.Change = repo.Head, repo.Head, change
	fields := change.Document.Frontmatter
	state.Phase, state.Run = str(fields["phase"]), str(fields["run"])
	regenerate, _ := fields["regenerate"].(bool)
	unstarted := state.Phase == ""
	var saved savedAttempt
	if attempt != 0 {
		if unstarted {
			return state, failure("P01", change.Path, "attempt requires a committed phase")
		}
		saved, err = loadAttempt(root, change, attempt)
		if err != nil {
			return state, err
		}
	}

	if err := staticFailure(project.Check(root)); err != nil {
		return state, err
	}
	state.Snapshot, err = project.Project(root, change.Path)
	if err != nil {
		return state, err
	}
	local, shared, err := baseRefs(state.Snapshot)
	if err != nil {
		return state, err
	}
	count, err := unapprovedCount(root, repo.Head, shared)
	if err != nil {
		return state, err
	}
	limit := integer(mapping(state.Snapshot.Config["authoring"])["max_unapproved"], 5)
	if count >= limit {
		d := project.Diagnostic{Code: "P07", Path: change.Path, Message: fmt.Sprintf("%d unapproved applied changes reaches the limit of %d", count, limit)}
		state.Diagnostics = append(state.Diagnostics, d)
		return state, diagnosticError(d)
	}
	if !atRevision(root, local, change.Path) && !atRevision(root, shared, change.Path) {
		state.Diagnostics = append(state.Diagnostics, warning("P04", change.Path, "change is absent from the local and shared base"))
	}
	if regenerate {
		latest, err := latestImplementation(root, change.Module)
		if err != nil {
			return state, err
		}
		date, result := drillDate(root, change.Module)
		if result != "pass" || !date.After(latest) {
			state.Diagnostics = append(state.Diagnostics, warning("P10", change.Path, "%s has no passing drill newer than its latest applied implementation", change.Module))
		}
	}

	// Prepare the claim and complete document before making any working-tree edits.
	original, err := os.ReadFile(disk(root, change.Path))
	if err != nil {
		return state, err
	}
	approved := change.Document
	if unstarted {
		approved, err = assignIDs(root, change, &state.Snapshot)
		if err != nil {
			return state, err
		}
		approved.Frontmatter = maps.Clone(approved.Frontmatter)
		state.Phase = initialPhase(change)
		if state.Run == "" {
			var nonce [12]byte
			if _, err := rand.Read(nonce[:]); err != nil {
				return state, err
			}
			state.Run = fmt.Sprintf("%s-%x", time.Now().UTC().Format("20060102T150405"), nonce)
		}
		approved.Frontmatter["phase"], approved.Frontmatter["run"] = state.Phase, state.Run
	}
	approvedText, err := contracts.RenderDocument(approved)
	if err != nil {
		return state, err
	}
	projection, err := recordSnapshot(state.Snapshot)
	if err != nil {
		return state, err
	}
	if held, oid, err := ownedClaim(root); err != nil {
		return state, err
	} else if oid != "" && held.Slug != change.Slug {
		return state, failure("P02", held.Slug, "worktree already holds change %s", held.Slug)
	}

	c := claim{Worktree: root, Slug: change.Slug, Path: change.Path, Base: repo.Head, Approved: approvedText, Projection: projection}
	oid, err := encodeClaim(root, c)
	if err != nil {
		return state, coded("P02", change.Slug, err)
	}
	if err := gitstore.CASRef(root, claimsPrefix+change.Slug, "", oid); err != nil {
		return state, coded("P02", change.Slug, fmt.Errorf("claim already held: %w", err))
	}
	state.ClaimOID = oid

	head, err := git(root, "rev-parse", "--verify", "HEAD")
	if err != nil || strings.TrimSpace(head) != state.PhaseBase {
		if err == nil {
			err = fmt.Errorf("phase base %s differs from HEAD %s", state.PhaseBase, strings.TrimSpace(head))
		}
		return state, errors.Join(coded("P05", change.Path, err), releaseClaim(root, change.Slug, oid))
	}
	if unstarted {
		err = writeFile(disk(root, change.Path), []byte(approvedText))
	}
	if err == nil && unstarted && regenerate {
		err = removeDisposable(state)
	}
	if err == nil && attempt != 0 {
		err = restoreAttempt(root, change, saved)
	}
	if err != nil {
		restoreErr := writeFile(disk(root, change.Path), original)
		releaseErr := releaseClaim(root, change.Slug, oid)
		return state, errors.Join(err, restoreErr, releaseErr)
	}
	state.Change, err = readChange(root, "", change)
	return state, err
}

func initialPhase(change project.Change) string {
	fields := change.Document.Frontmatter
	regenerate, _ := fields["regenerate"].(bool)
	if !regenerate && (len(change.Added)+len(change.Modified)+len(change.Removed) > 0 ||
		len(stringsOf(fields["fixes"]))+len(stringsOf(fields["eval_fix"])) > 0) {
		return "evals"
	}
	return "impl"
}

var handle = regexp.MustCompile(`new:[a-z0-9-]+`)

func assignIDs(root string, change project.Change, snapshot *project.Snapshot) (contracts.Document, error) {
	doc := change.Document
	if len(change.Added) == 0 {
		return doc, nil
	}
	live, err := project.Load(root)
	if err != nil {
		return doc, err
	}
	module := moduleNamed(live, change.Module)
	current, err := items(module.Document, "Specs")
	if err != nil {
		return doc, err
	}
	maxID := 0
	reserve := func(list []contracts.Item) {
		for _, item := range list {
			if suffix, ok := strings.CutPrefix(item.ID, module.IDPrefix+"-"); ok {
				if number, err := strconv.Atoi(suffix); err == nil && number > maxID {
					maxID = number
				}
			}
		}
	}
	reserve(current)
	reserve(change.Added)
	archived, err := changePaths(root, "", archiveDir)
	if err != nil {
		return doc, err
	}
	for _, identity := range archived {
		archived, err := readChange(root, "", identity)
		if err != nil {
			return doc, err
		}
		if archived.Module == change.Module {
			reserve(archived.Removed)
		}
	}
	replacements := make(map[string]string)
	for _, item := range change.Added {
		if strings.HasPrefix(item.ID, "new:") {
			maxID++
			replacements[item.ID] = fmt.Sprintf("%s-%03d", module.IDPrefix, maxID)
		}
	}
	replace := func(text string) string {
		return handle.ReplaceAllStringFunc(text, func(token string) string {
			if id, ok := replacements[token]; ok {
				return id
			}
			return token
		})
	}
	raw, err := contracts.RenderDocument(doc)
	if err != nil {
		return doc, err
	}
	doc, err = contracts.ParseDocument(replace(raw))
	for i, candidate := range snapshot.Modules {
		if candidate.Name == change.Module {
			snapshot.Modules[i].Document.Body = replace(candidate.Document.Body)
		}
	}
	return doc, err
}

func atRevision(root, revision, file string) bool {
	out, err := git(root, "ls-tree", "--name-only", revision, "--", file)
	return err == nil && out != ""
}

func removeDisposable(state model.State) error {
	root := state.Snapshot.Root
	module := moduleNamed(state.Snapshot, state.Change.Module)
	out, err := git(root, "ls-files", "-z", "--", module.Dir)
	if err != nil {
		return err
	}
	var removed []string
	for _, file := range gitFields(out) {
		if fileZone(state.Snapshot, module, file) != "disposable" {
			continue
		}
		if err := os.Remove(disk(root, file)); err != nil {
			var restoreErr error
			if len(removed) > 0 {
				_, restoreErr = git(root, append([]string{"restore", "--worktree", "--source", state.PhaseBase, "--"}, removed...)...)
			}
			return errors.Join(err, restoreErr)
		}
		removed = append(removed, file)
	}
	return nil
}

func verifyApproval(state model.State, c claim, declaringBlock bool) error {
	approved, err := c.approvedDocument()
	if err != nil {
		return err
	}
	actual := state.Change.Document
	for _, key := range []string{"run", "phase"} {
		if !reflect.DeepEqual(actual.Frontmatter[key], approved.Frontmatter[key]) {
			return failure("P03", c.Path, "approved run or phase changed")
		}
	}
	if !reflect.DeepEqual(actual.Frontmatter["audit_rounds"], approved.Frontmatter["audit_rounds"]) ||
		state.Phase != "audit" && !sameSection(approved, actual, "Audit") {
		return failure("P03", c.Path, "approved audit state changed")
	}
	a, b := maps.Clone(approved.Frontmatter), maps.Clone(actual.Frontmatter)
	for _, key := range []string{"run", "phase", "audit_rounds"} {
		delete(a, key)
		delete(b, key)
	}
	if declaringBlock {
		delete(b, "blocked")
	}
	if !reflect.DeepEqual(a, b) {
		return failure("P03", c.Path, "approved author metadata changed")
	}

	// Compare the rest of the body after removing only worker-editable sections.
	editable := []string{"Decisions", "Questions", "Notes"}
	if state.Phase == "audit" {
		editable = []string{"Audit"}
	}
	if immutableBody(approved, editable) != immutableBody(actual, editable) {
		return failure("P03", c.Path, "approved author metadata changed")
	}
	if state.Head != c.Base {
		return coded("P05", c.Path, fmt.Errorf("phase base %s differs from HEAD %s", c.Base, state.Head))
	}
	return nil
}

func sameSection(a, b contracts.Document, name string) bool {
	return strings.TrimSpace(section(a, name)) == strings.TrimSpace(section(b, name))
}

func immutableBody(doc contracts.Document, editable []string) string {
	body := doc.Body
	for _, name := range editable {
		_, exists := contracts.Section(body, name)
		if !exists {
			continue
		}
		// Replace the grammar-recognized section with a marker, then remove its heading.
		const marker = "\x00fireweed-section\x00\n"
		body = contracts.ReplaceSection(body, name, marker)
		index := strings.Index(body, marker)
		start := strings.LastIndex(strings.TrimSuffix(body[:index], "\n"), "\n") + 1
		body = body[:start] + body[index+len(marker):]
	}
	lines := strings.Split(body, "\n")
	lines = slices.DeleteFunc(lines, func(line string) bool { return strings.TrimSpace(line) == "" })
	return strings.Join(lines, "\n")
}
