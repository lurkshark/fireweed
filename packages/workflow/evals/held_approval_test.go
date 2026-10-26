package evals

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	c "github.com/lurkshark/fireweed/packages/contracts"
	p "github.com/lurkshark/fireweed/packages/project"
	w "github.com/lurkshark/fireweed/packages/workflow"
)

func heldSourceFixture(t *testing.T, phase, format string, late bool) approvalFixtureData {
	t.Helper()
	f := approvalFixture(t, phase, format, late)
	// Exact LoadApproval intentionally ignores this payload. Implicit selection
	// cannot classify it, so remove it when establishing valid source controls.
	claimGit(t, f.Root, "update-ref", "-d", "refs/fireweed/claims/unrelated-malformed")
	return f
}

func heldPeer(t *testing.T, root string) string {
	t.Helper()
	peer := filepath.Join(t.TempDir(), "detached held peer")
	claimGit(t, root, "worktree", "add", "--detach", peer)
	return stateCanonicalDirectory(t, peer)
}
func heldPoisonCurrent(t *testing.T, root string) {
	t.Helper()
	for _, name := range []string{"fireweed.yaml", "units/a/SPEC.md", "changes/0001-before.md", "changes/0002-current.md", "changes/0003-later.md", "changes/archive/0009-earlier.md"} {
		worktreePut(t, root, name, "[: invalid worker metadata "+name)
	}
	if _, err := p.Load(root); err == nil {
		t.Fatal("poisoned current-source control was accepted")
	}
}
func heldSourceResult(t *testing.T, start string, f approvalFixtureData) p.Snapshot {
	t.Helper()
	got, claim, present, err := w.LoadHeldApproval(start)
	if err != nil || !present || claim != f.Claim || !reflect.DeepEqual(got, f.Want) {
		t.Errorf("held result present=%t error=%v claim=%+v; want=%+v\nsource=%s\nwant=%s", present, err, claim, f.Claim, policyJSON(t, got), policyJSON(t, f.Want))
	}
	return got
}
func heldAbsentResult(t *testing.T, start string) {
	t.Helper()
	got, claim, present, err := w.LoadHeldApproval(start)
	if err != nil || present || claim != (w.Claim{}) || !reflect.DeepEqual(got, p.Snapshot{}) {
		t.Errorf("absence returned partial authority: %t %+v %+v %v", present, got, claim, err)
	}
}
func heldRejectedResult(t *testing.T, start string) error {
	t.Helper()
	got, claim, present, err := w.LoadHeldApproval(start)
	if err == nil || present || claim != (w.Claim{}) || !reflect.DeepEqual(got, p.Snapshot{}) || !policyCorrective(err.Error()) {
		t.Errorf("unsafe/partial held authority: %t %+v %+v %v", present, got, claim, err)
	}
	return err
}
func Test_FLOW048_FLOW049HeldApprovedSources(t *testing.T) {
	for _, format := range []string{"sha1", "sha256"} {
		for _, mode := range []string{"ordinary", "nested", "linked", "linked-nested", "late", "evals", "audit", "impl", "moved-head"} {
			t.Run(format+"/"+mode, func(t *testing.T) {
				phase := "impl"
				if mode == "evals" || mode == "audit" {
					phase = mode
				}
				f := heldSourceFixture(t, phase, format, mode == "late")
				if mode == "moved-head" {
					configuration, err := os.ReadFile(filepath.Join(f.Root, "fireweed.yaml"))
					if err != nil {
						t.Fatal(err)
					}
					worktreePut(t, f.Root, "fireweed.yaml", string(configuration)+"\nverify_dependents: none\n")
					worktreePut(t, f.Root, "units/a/SPEC.md", strings.Replace(f.Spec, "retained behavior MUST stay", "new HEAD behavior MUST be distinct", 1))
					newHead := approvalCommit(t, f.Root)
					if newHead == f.Claim.Base {
						t.Fatal("distinct committed HEAD premise")
					}
					current, err := p.Load(f.Root)
					if err != nil || reflect.DeepEqual(current, f.Want) {
						t.Fatal("new HEAD metadata does not distinguish approval", err)
					}
				}
				start := f.Root
				if strings.HasPrefix(mode, "linked") {
					peer := heldPeer(t, f.Root)
					f.Claim.Worktree = peer
					f.Claim.Object = claimRaw(t, peer, "current", claimPayload(t, f.Claim))
					f.Want.Root = peer
					start = peer
				}
				if strings.Contains(mode, "nested") {
					start = filepath.Join(start, "units/a/e")
				}
				// Keep a fully valid foreign ownership record: implicit selection must
				// ignore it, while retaining the exact locally owned claim and token.
				foreign := f.Claim
				foreign.Change = "foreign"
				foreign.Worktree = heldPeer(t, f.Root)
				claimRaw(t, f.Root, foreign.Change, claimPayload(t, foreign))
				want, cl, err := w.LoadApproval(start, "current")
				if err != nil || cl != f.Claim || !reflect.DeepEqual(want, f.Want) {
					t.Fatal("existing public approved-source control", cl, err)
				}
				heldPoisonCurrent(t, f.Want.Root)
				got := heldSourceResult(t, start, f)
				if len(got.Modules) == 0 || len(got.Changes) < 2 {
					return
				}
				if len(got.Modules[0].Specs) != 1 {
					t.Error("snapshot prematurely projected the selected change")
				}
				before, err := p.Project(got, "current")
				if err != nil {
					t.Error(err)
					return
				}
				selected := got.Changes[1]
				inclusive, err := p.ApplyChange(before.Modules[0], selected)
				if err != nil || len(inclusive.Specs) != 2 || len(before.Modules[1].Specs) != 2 || !reflect.DeepEqual(inclusive.Interface, []string{"pub/**"}) || !reflect.DeepEqual(inclusive.DependsOn, []string{"gamma", "beta"}) {
					t.Error("once-only preceding and selected projection", inclusive, before, err)
				}
				for _, spec := range inclusive.Specs {
					if spec.ID == "ALP-003" {
						t.Error("later queue effect included")
					}
				}
			})
		}
	}
}

func Test_FLOW050HeldCommittedSourcePaths(t *testing.T) {
	if os.PathSeparator == '\\' {
		t.Skip("POSIX committed link fixtures; portable object/source cases run separately")
	}
	for _, format := range []string{"sha1", "sha256"} {
		for _, mode := range []string{"config-escape", "module-escape", "selected-change-escape", "module-dangling", "module-loop"} {
			t.Run(format+"/"+mode, func(t *testing.T) {
				f := heldSourceFixture(t, "impl", format, false)
				selected := "units/a/SPEC.md"
				if mode == "config-escape" {
					selected = "fireweed.yaml"
				} else if mode == "selected-change-escape" {
					selected = f.Path
				}
				original, err := os.ReadFile(filepath.Join(f.Root, selected))
				if err != nil {
					t.Fatal(err)
				}
				outside := filepath.Join(stateCanonicalDirectory(t, t.TempDir()), "outside-authority")
				if err := os.WriteFile(outside, original, 0600); err != nil {
					t.Fatal(err)
				}
				target := outside
				if mode == "module-dangling" {
					target = "missing-contained-module"
				} else if mode == "module-loop" {
					target = "SPEC.md"
				}
				if err := os.Remove(filepath.Join(f.Root, selected)); err != nil {
					t.Fatal(err)
				}
				if err := os.Symlink(target, filepath.Join(f.Root, selected)); err != nil {
					t.Fatal(err)
				}
				f.Claim.Base = approvalCommit(t, f.Root)
				f.Claim.Object = claimRaw(t, f.Root, "current", claimPayload(t, f.Claim))
				if source, cl, err := w.LoadApproval(f.Root, "current"); err == nil || cl != (w.Claim{}) || !reflect.DeepEqual(source, p.Snapshot{}) {
					t.Fatal("unsafe committed authority control", err, cl)
				}
				// Current contained bytes are a decoy, never approval authority.
				if err := os.Remove(filepath.Join(f.Root, selected)); err != nil {
					t.Fatal(err)
				}
				worktreePut(t, f.Root, selected, string(original))
				heldRejectedResult(t, f.Root)
				if after, err := os.ReadFile(outside); err != nil || string(after) != string(original) {
					t.Fatal("outside authority bytes changed", err)
				}
			})
		}
	}
}
func Test_FLOW048HeldAbsenceAndForeignOwnership(t *testing.T) {
	for _, format := range []string{"sha1", "sha256"} {
		for _, mode := range []string{"none", "foreign", "linked-foreign"} {
			t.Run(format+"/"+mode, func(t *testing.T) {
				f := heldSourceFixture(t, "impl", format, false)
				start := filepath.Join(f.Root, "units/a/e")
				if mode == "none" {
					claimGit(t, f.Root, "update-ref", "-d", "refs/fireweed/claims/current", f.Claim.Object)
				} else {
					if mode == "linked-foreign" {
						start = filepath.Join(heldPeer(t, f.Root), "units/a/e")
					} else {
						f.Claim.Worktree = heldPeer(t, f.Root)
						f.Claim.Object = claimRaw(t, f.Root, "current", claimPayload(t, f.Claim))
					}
				}
				heldPoisonCurrent(t, filepath.Clean(filepath.Join(start, "../../..")))
				heldAbsentResult(t, start)
			})
		}
	}
}

func Test_FLOW050HeldMalformedCommittedContent(t *testing.T) {
	for _, format := range []string{"sha1", "sha256"} {
		for _, path := range []string{"fireweed.yaml", "units/a/SPEC.md", "changes/0002-current.md"} {
			t.Run(format+"/"+path, func(t *testing.T) {
				f := heldSourceFixture(t, "impl", format, false)
				original, err := os.ReadFile(filepath.Join(f.Root, path))
				if err != nil {
					t.Fatal(err)
				}
				worktreePut(t, f.Root, path, "---\n[malformed committed approval\n")
				f.Claim.Base = approvalCommit(t, f.Root)
				f.Claim.Object = claimRaw(t, f.Root, "current", claimPayload(t, f.Claim))
				if source, claim, err := w.LoadApproval(f.Root, "current"); err == nil || claim != (w.Claim{}) || !reflect.DeepEqual(source, p.Snapshot{}) {
					t.Fatal("malformed selected approval control", err, claim)
				}
				worktreePut(t, f.Root, path, string(original))
				heldRejectedResult(t, f.Root)
			})
		}
	}
}
func Test_FLOW048_FLOW050HeldMalformedSelection(t *testing.T) {
	for _, mode := range []string{"malformed-selected", "malformed-unclassified", "unreadable-unclassified", "invalid-owner-type", "nonblob-unclassified", "empty-unclassified-owner", "ambiguous-local", "symbolic-selected", "missing-payload", "invalid-state", "wrong-run", "reservation"} {
		t.Run(mode, func(t *testing.T) {
			f := heldSourceFixture(t, "impl", "sha1", false)
			switch mode {
			case "malformed-selected":
				claimRaw(t, f.Root, "current", "{broken JSON")
			case "malformed-unclassified":
				claimRaw(t, f.Root, "cannot-classify", "{broken JSON")
			case "unreadable-unclassified":
				other := f.Claim
				other.Change = "cannot-classify"
				object := claimRaw(t, f.Root, other.Change, claimPayload(t, other))
				if err := os.Remove(filepath.Join(claimCommon(t, f.Root), "objects", object[:2], object[2:])); err != nil {
					t.Fatal(err)
				}
				if _, err := claimCommand(t, f.Root, "", "cat-file", "-e", object); err == nil {
					t.Fatal("unclassified payload is still available")
				}
			case "invalid-owner-type":
				other := f.Claim
				other.Change = "cannot-classify"
				payload := claimPayload(t, other)
				raw := strings.Replace(payload, `"worktree":"`+f.Root+`"`, `"worktree":42`, 1)
				if raw == payload {
					t.Fatal("ownership-type mutation did not match payload")
				}
				claimRaw(t, f.Root, "cannot-classify", raw)
			case "nonblob-unclassified":
				claimGit(t, f.Root, "update-ref", "refs/fireweed/claims/cannot-classify", f.Claim.Base)
				if claimGit(t, f.Root, "cat-file", "-t", "refs/fireweed/claims/cannot-classify") != "commit" {
					t.Fatal("nonblob claim object control")
				}
			case "empty-unclassified-owner":
				other := f.Claim
				other.Change, other.Worktree = "cannot-classify", ""
				claimRaw(t, f.Root, other.Change, claimPayload(t, other))
			case "ambiguous-local":
				other := f.Claim
				other.Change = "second"
				claimRaw(t, f.Root, "second", claimPayload(t, other))
			case "symbolic-selected":
				claimGit(t, f.Root, "symbolic-ref", "refs/fireweed/claims/current", "refs/fireweed/claims/alias")
			case "missing-payload":
				claimGit(t, f.Root, "update-ref", "-d", "refs/fireweed/claims/current", f.Claim.Object)
				if err := os.Remove(filepath.Join(claimCommon(t, f.Root), "objects", f.Claim.Object[:2], f.Claim.Object[2:])); err != nil {
					t.Fatal(err)
				}
				// Git refuses to install a missing object. Preserve the exact native
				// ref syntax directly, then independently prove the object is absent.
				worktreePut(t, claimCommon(t, f.Root), "refs/fireweed/claims/current", f.Claim.Object+"\n")
				if _, err := claimCommand(t, f.Root, "", "cat-file", "-e", f.Claim.Object); err == nil {
					t.Fatal("selected payload is not genuinely missing")
				}
			case "invalid-state":
				f.Claim.State = "---\n[broken"
				claimRaw(t, f.Root, "current", claimPayload(t, f.Claim))
			case "wrong-run":
				f.Claim.Run = "different"
				claimRaw(t, f.Root, "current", claimPayload(t, f.Claim))
			case "reservation":
				f.Claim.State = ""
				claimRaw(t, f.Root, "current", claimPayload(t, f.Claim))
			}
			heldRejectedResult(t, f.Root)
		})
	}
}

func heldMutateContainers(value reflect.Value, marker string, kinds map[reflect.Kind]int) int {
	if !value.IsValid() {
		return 0
	}
	switch value.Kind() {
	case reflect.Interface:
		if value.IsNil() || !value.CanSet() {
			return 0
		}
		copy := reflect.New(value.Elem().Type()).Elem()
		copy.Set(value.Elem())
		count := heldMutateContainers(copy, marker, kinds)
		value.Set(copy)
		return count
	case reflect.Struct:
		count := 0
		for index := 0; index < value.NumField(); index++ {
			if value.Type().Field(index).PkgPath == "" {
				count += heldMutateContainers(value.Field(index), marker, kinds)
			}
		}
		return count
	case reflect.Slice:
		if value.Len() != 0 {
			kinds[reflect.Slice]++
		}
		count := 0
		for index := 0; index < value.Len(); index++ {
			count += heldMutateContainers(value.Index(index), marker, kinds)
		}
		return count
	case reflect.Map:
		if value.IsNil() {
			return 0
		}
		if value.Len() != 0 {
			kinds[reflect.Map]++
		}
		count := 0
		for _, key := range value.MapKeys() {
			copy := reflect.New(value.Type().Elem()).Elem()
			copy.Set(value.MapIndex(key))
			count += heldMutateContainers(copy, marker, kinds)
			value.SetMapIndex(key, copy)
		}
		return count
	case reflect.String:
		if value.CanSet() {
			value.SetString(marker)
			return 1
		}
	case reflect.Bool:
		if value.CanSet() {
			value.SetBool(!value.Bool())
			return 1
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if value.CanSet() {
			value.SetInt(value.Int() + 1)
			return 1
		}
	}
	return 0
}

func heldEmptyContainers(t *testing.T, a, b reflect.Value, counts map[reflect.Kind]int) {
	t.Helper()
	if !a.IsValid() || !b.IsValid() || a.Type() != b.Type() {
		return
	}
	marked := func(typ reflect.Type, marker string, second bool) reflect.Value {
		value := reflect.New(typ).Elem()
		heldMutateContainers(value, marker, map[reflect.Kind]int{})
		switch typ.Kind() {
		case reflect.Interface:
			value.Set(reflect.ValueOf(marker))
		case reflect.Bool:
			value.SetBool(second)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if second {
				value.SetInt(2)
			} else {
				value.SetInt(1)
			}
		}
		return value
	}
	switch a.Kind() {
	case reflect.Struct:
		for index := 0; index < a.NumField(); index++ {
			if a.Type().Field(index).PkgPath == "" {
				heldEmptyContainers(t, a.Field(index), b.Field(index), counts)
			}
		}
	case reflect.Interface:
		if !a.IsNil() && !b.IsNil() && a.CanSet() && b.CanSet() {
			av, bv := reflect.New(a.Elem().Type()).Elem(), reflect.New(b.Elem().Type()).Elem()
			av.Set(a.Elem())
			bv.Set(b.Elem())
			heldEmptyContainers(t, av, bv, counts)
			a.Set(av)
			b.Set(bv)
		}
	case reflect.Slice:
		if a.Len() == 0 && b.Len() == 0 && a.CanSet() && b.CanSet() {
			a.Set(reflect.Append(a, marked(a.Type().Elem(), "first-empty-append", false)))
			before := policyJSON(t, a.Interface())
			b.Set(reflect.Append(b, marked(b.Type().Elem(), "second-empty-append", true)))
			if policyJSON(t, a.Interface()) != before {
				t.Error("empty slices share writable append backing storage", a.Type())
			}
			counts[reflect.Slice]++
		} else {
			for index := 0; index < a.Len() && index < b.Len(); index++ {
				heldEmptyContainers(t, a.Index(index), b.Index(index), counts)
			}
		}
	case reflect.Map:
		if a.IsNil() || b.IsNil() {
			return
		}
		if a.Len() == 0 && b.Len() == 0 && a.Type().Key().Kind() == reflect.String {
			key := reflect.New(a.Type().Key()).Elem()
			key.SetString("new-empty-map-entry")
			a.SetMapIndex(key, marked(a.Type().Elem(), "first-empty-map", false))
			if b.MapIndex(key).IsValid() {
				t.Error("empty writable maps share inserted entries", a.Type())
			}
			counts[reflect.Map]++
		} else {
			for _, key := range a.MapKeys() {
				if !b.MapIndex(key).IsValid() {
					continue
				}
				av, bv := reflect.New(a.Type().Elem()).Elem(), reflect.New(b.Type().Elem()).Elem()
				av.Set(a.MapIndex(key))
				bv.Set(b.MapIndex(key))
				heldEmptyContainers(t, av, bv, counts)
				a.SetMapIndex(key, av)
				b.SetMapIndex(key, bv)
			}
		}
	}
}

func Test_FLOW049_FLOW051HeldResultOwnership(t *testing.T) {
	for _, format := range []string{"sha1", "sha256"} {
		t.Run(format, func(t *testing.T) {
			f := heldSourceFixture(t, "impl", format, true)
			configuration, err := os.ReadFile(filepath.Join(f.Root, "fireweed.yaml"))
			if err != nil {
				t.Fatal(err)
			}
			configuration = []byte(strings.Replace(string(configuration), "  command: [context-agent-trap]", "  command: [context-agent-trap]\n  wrap: [wrapper, '{prompt}']", 1))
			worktreePut(t, f.Root, "fireweed.yaml", string(configuration))
			parsed, err := c.ParseChange(f.Claim.State, f.Path, "ALP")
			if err != nil {
				t.Fatal(err)
			}
			document := parsed.Document
			document.Metadata["bases"] = map[string]string{"ALP-001": c.SpecHash(f.Want.Modules[0].Specs[0])}
			document.Sections = append(document.Sections, c.Section{Name: "Modified", Body: "### ALP-001\nThe retained behavior MUST stay.\n"})
			f.Claim.State, err = c.RenderMarkdown(document)
			if err != nil {
				t.Fatal(err)
			}
			worktreePut(t, f.Root, f.Path, f.Claim.State)
			worktreePut(t, f.Root, "changes/0004-empty.md", "---\nmodule: alpha\nset: {}\n---\n\n## Motivation\n\nPopulate an independent writable empty map.\n\n## Decisions\n\nAn empty config edit has no projected effect.\n")
			f.Claim.Base = approvalCommit(t, f.Root)
			f.Claim.Object = claimRaw(t, f.Root, "current", claimPayload(t, f.Claim))
			f.Want, _, err = w.LoadApproval(f.Root, "current")
			if err != nil || len(f.Want.Config.Agent.Wrap) != 2 || len(f.Want.Changes[1].Bases) != 1 || len(f.Want.Changes[1].Modified) != 1 {
				t.Fatal("rich writable source-container control", err)
			}
			control, _, err := w.LoadApproval(f.Root, "current")
			if err != nil {
				t.Fatal(err)
			}
			controlBefore := policyJSON(t, f.Want)
			controlKinds := map[reflect.Kind]int{}
			if count := heldMutateContainers(reflect.ValueOf(&control).Elem(), "container-calibration", controlKinds); count < 50 || controlKinds[reflect.Map] < 3 || controlKinds[reflect.Slice] < 8 {
				t.Fatal("nested mutation observer has insufficient populated controls", count, controlKinds)
			}
			if policyJSON(t, f.Want) != controlBefore {
				t.Fatal("existing independent approved-source control aliases its peer")
			}
			emptyA, _, _ := w.LoadApproval(f.Root, "current")
			emptyB, _, _ := w.LoadApproval(f.Root, "current")
			emptyKinds := map[reflect.Kind]int{}
			heldEmptyContainers(t, reflect.ValueOf(&emptyA).Elem(), reflect.ValueOf(&emptyB).Elem(), emptyKinds)
			if emptyKinds[reflect.Slice] == 0 || emptyKinds[reflect.Map] == 0 {
				t.Fatal("empty writable container control absent", emptyKinds)
			}
			heldPoisonCurrent(t, f.Root)
			a, ac, ap, ae := w.LoadHeldApproval(f.Root)
			b, bc, bp, be := w.LoadHeldApproval(f.Root)
			if ae != nil || be != nil || !ap || !bp || ac != f.Claim || bc != f.Claim || !reflect.DeepEqual(a, f.Want) || !reflect.DeepEqual(b, f.Want) {
				t.Fatal("repeat held-source values", ae, be, ap, bp)
			}
			before := policyJSON(t, b)
			a.Modules[0].Name = "mutated"
			a.Modules[1].Document.Metadata["name"] = "mutated"
			a.Config.Protected[0] = "mutated"
			a.Changes[1].Set["depends_on"].([]any)[0] = "mutated"
			a.Changes[1].Document.Sections[0].Body = "mutated"
			a.Archive[0].Added[0].Body = "mutated"
			a.Diagnostics[0].Message = "mutated"
			ac.State = "mutated"
			if policyJSON(t, b) != before || bc != f.Claim {
				t.Fatal("peer result aliases writable source fields")
			}
			kinds := map[reflect.Kind]int{}
			if count := heldMutateContainers(reflect.ValueOf(&a).Elem(), "changed-every-container", kinds); count < 50 || kinds[reflect.Map] < 3 || kinds[reflect.Slice] < 8 {
				t.Fatal("rich nested mutation control", count, kinds)
			}
			if policyJSON(t, b) != before {
				t.Fatal("nested specs/zones/dependencies/config/document/change containers alias peer")
			}
			x, _, xp, xe := w.LoadHeldApproval(f.Root)
			y, _, yp, ye := w.LoadHeldApproval(f.Root)
			if xe != nil || ye != nil || !xp || !yp {
				t.Fatal("empty writable pair", xe, ye, xp, yp)
			}
			emptyKinds = map[reflect.Kind]int{}
			heldEmptyContainers(t, reflect.ValueOf(&x).Elem(), reflect.ValueOf(&y).Elem(), emptyKinds)
			if emptyKinds[reflect.Slice] == 0 || emptyKinds[reflect.Map] == 0 {
				t.Fatal("held empty writable container shape changed", emptyKinds)
			}
			again, next, present, err := w.LoadHeldApproval(f.Root)
			if err != nil || !present || next != f.Claim || policyJSON(t, again) != before {
				t.Fatal("mutation poisoned later held source", err, present, next)
			}
		})
	}
}

func Test_FLOW049_FLOW050HeldLocalObjectAvailability(t *testing.T) {
	probe := heldBuildColdProbe(t)
	for _, format := range []string{"sha1", "sha256"} {
		for _, mode := range []string{"unrelated-body", "unneeded-parent", "base", "root-tree", "config", "module", "selected-change", "archive"} {
			t.Run(format+"/"+mode, func(t *testing.T) {
				f := heldSourceFixture(t, "impl", format, true)
				selector := f.Claim.Base
				switch mode {
				case "unrelated-body":
					selector += ":units/a/impl/private.txt"
				case "unneeded-parent":
					selector += "^"
				case "root-tree":
					selector += "^{tree}"
				case "config":
					selector += ":fireweed.yaml"
				case "module":
					selector += ":units/a/SPEC.md"
				case "selected-change":
					selector += ":" + f.Path
				case "archive":
					selector += ":changes/archive/0009-earlier.md"
				}
				object := claimGit(t, f.Root, "rev-parse", "--verify", selector)
				if len(object) != len(f.Claim.Base) {
					t.Fatal("full native object identity premise", object)
				}
				if err := os.Remove(filepath.Join(claimCommon(t, f.Root), "objects", object[:2], object[2:])); err != nil {
					t.Fatal("remove owned loose fixture object", err)
				}
				if _, err := claimCommand(t, f.Root, "", "cat-file", "-e", object); err == nil {
					t.Fatal("selected native object is not actually unavailable")
				}
				available := mode == "unrelated-body" || mode == "unneeded-parent" || mode == "archive"
				control := heldColdCall(t, probe, "approval", f.Root)
				before, cl := control.Snapshot, control.Claim
				if mode == "archive" {
					// PRJ-025 retains categorized nonfatal document diagnostics;
					// LoadHeldApproval must preserve that public source behavior.
					found := false
					for _, diagnostic := range before.Diagnostics {
						found = found || diagnostic.Code == "C05" && diagnostic.Path == "changes/archive/0009-earlier.md"
					}
					if control.Error != "" || cl != f.Claim || len(before.Archive) != 0 || !found {
						t.Fatal("categorized nonfatal archive source control", control.Error, before.Diagnostics)
					}
					f.Want = before
				}
				if available {
					if control.Error != "" || cl != f.Claim || !reflect.DeepEqual(before, f.Want) {
						t.Fatal("locally complete existing public source control", control.Error, cl)
					}
				} else if control.Error == "" || cl != (w.Claim{}) || !reflect.DeepEqual(before, p.Snapshot{}) {
					t.Fatal("required-object rejection control", control.Error, cl)
				}
				heldPoisonCurrent(t, f.Root)
				heldColdAssert(t, heldColdCall(t, probe, "held", f.Root), f, available)
			})
		}
	}
}

func Test_FLOW048_FLOW050_FLOW051HeldReadOnlyState(t *testing.T) {
	for _, format := range []string{"sha1", "sha256"} {
		for _, linked := range []bool{false, true} {
			for _, mode := range []string{"held", "none", "foreign", "corrupt"} {
				t.Run(format+"/"+fmt.Sprint(linked)+"/"+mode, func(t *testing.T) {
					f := heldSourceFixture(t, "audit", format, false)
					originalRoot := f.Root
					if linked {
						f.Root = heldPeer(t, f.Root)
						f.Claim.Worktree = f.Root
						f.Claim.Object = claimRaw(t, f.Root, "current", claimPayload(t, f.Claim))
						f.Want.Root = f.Root
					}
					switch mode {
					case "none":
						claimGit(t, f.Root, "update-ref", "-d", "refs/fireweed/claims/current", f.Claim.Object)
					case "foreign":
						f.Claim.Worktree = heldPeer(t, originalRoot)
						claimRaw(t, f.Root, "current", claimPayload(t, f.Claim))
					case "corrupt":
						claimRaw(t, f.Root, "current", "{broken implicit ownership")
					}
					heldPoisonCurrent(t, f.Root)
					worktreePut(t, f.Root, "untracked-sentinel", "retain untracked bytes\n")
					worktreePut(t, f.Root, "ignored/retained", "retain ignored bytes\n")
					common := claimCommon(t, f.Root)
					worktreePut(t, common, "fireweed/attempts-sentinel", "retain attempt metadata\n")
					worktreePut(t, f.Root, ".fireweed/drills/alpha.md", "retain drill report\n")
					// Native fixture setup and any transport/hook calibration must
					// precede these exact images; never exclude loader effects.
					tmp := stateCanonicalDirectory(t, t.TempDir())
					t.Setenv("TMPDIR", tmp)
					t.Setenv("TMP", tmp)
					t.Setenv("TEMP", tmp)
					files := claimTree(t, f.Root, func(path string) bool { return path == ".git" })
					original := claimTree(t, originalRoot, func(path string) bool { return path == ".git" })
					metadata := claimTree(t, common, func(string) bool { return false })
					temps := claimTree(t, tmp, func(string) bool { return false })
					check := protocolParent(t)
					defer check()
					switch mode {
					case "held":
						heldSourceResult(t, f.Root, f)
					case "none", "foreign":
						heldAbsentResult(t, f.Root)
					case "corrupt":
						heldRejectedResult(t, f.Root)
					}
					if !reflect.DeepEqual(files, claimTree(t, f.Root, func(path string) bool { return path == ".git" })) || !reflect.DeepEqual(original, claimTree(t, originalRoot, func(path string) bool { return path == ".git" })) || !reflect.DeepEqual(metadata, claimTree(t, common, func(string) bool { return false })) || !reflect.DeepEqual(temps, claimTree(t, tmp, func(string) bool { return false })) {
						t.Fatal("held selection/source mutated files, modes, index, refs, attempts, drills or temporary storage")
					}
				})
			}
		}
	}
}
