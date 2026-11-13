package drill

import (
	"context"
	"fmt"
	"io/fs"
	"maps"
	"os"
	"slices"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/runner/internal/invoke"
	"github.com/lurkshark/fireweed/packages/runner/internal/workspace"
)

func trackedFiles(ctx context.Context, root, base string) ([]string, error) {
	files, err := workspace.Git(ctx, root, "ls-tree", "-r", "--name-only", "-z", base)
	if err != nil || files == "" {
		return nil, err
	}
	return strings.Split(strings.TrimSuffix(files, "\x00"), "\x00"), nil
}

func disposableFile(snapshot project.Snapshot, target project.Module, name string) bool {
	return workspace.Owner(snapshot, name).Name == target.Name && workspace.Zone(snapshot, target, name) == "disposable"
}

func prepareInput(ctx context.Context, snapshot project.Snapshot, target project.Module, dependents []project.Module, tracked []string, scratch, input string) ([]string, error) {
	if err := workspace.Init(ctx, input); err != nil {
		return nil, err
	}
	var disposable []string
	for _, name := range tracked {
		if disposableFile(snapshot, target, name) {
			disposable = append(disposable, name)
			continue
		}
		if name == "changes" || strings.HasPrefix(name, "changes/") || strings.HasPrefix(name, ".fireweed/drills/") {
			continue
		}
		owner := workspace.Owner(snapshot, name).Name
		if slices.ContainsFunc(dependents, func(module project.Module) bool { return module.Name == owner }) {
			continue
		}
		file, err := workspace.Read(scratch, name)
		if err != nil {
			return nil, err
		}
		if err := file.Put(input, name); err != nil {
			return nil, err
		}
	}
	return disposable, workspace.Seed(ctx, input)
}

func deleteDisposable(snapshot project.Snapshot, target project.Module, root string) error {
	return workspace.Walk(root, func(name string, entry fs.DirEntry) error {
		if !entry.IsDir() && disposableFile(snapshot, target, name) {
			return os.Remove(workspace.Path(root, name))
		}
		return nil
	})
}

func inputFiles(root string) (map[string]workspace.File, error) {
	files := make(map[string]workspace.File)
	err := workspace.Walk(root, func(name string, entry fs.DirEntry) error {
		if entry.IsDir() {
			return nil
		}
		file, err := workspace.Read(root, name)
		if err == nil {
			files[name] = file
		}
		return err
	})
	return files, err
}

func unchangedInputs(original, submission map[string]workspace.File) error {
	for _, name := range slices.Sorted(maps.Keys(original)) {
		if file, present := submission[name]; !present || file != original[name] {
			return fmt.Errorf("reconstruction modified or deleted input %s", name)
		}
	}
	return nil
}

func reconstructionPrompt(snapshot project.Snapshot, target project.Module, initial map[string]workspace.File, bin string) (string, error) {
	spec, err := contracts.RenderDocument(target.Document)
	if err != nil {
		return "", err
	}
	var prompt strings.Builder
	fmt.Fprintf(&prompt, "Reconstruct module %s from its spec, Decisions, retained interfaces and evals.\n", target.Name)
	prompt.WriteString("Rebuild only the target module's disposable paths. Do not read outside this input workspace.\n")
	fmt.Fprintf(&prompt, "Disposable patterns in %s: %s\n\n%s\n", target.Dir, strings.Join(target.Zones.Disposable, ", "), spec)
	prompt.WriteString("\n## Retained interfaces and evals\n")
	for _, name := range slices.Sorted(maps.Keys(initial)) {
		if workspace.Owner(snapshot, name).Name == target.Name {
			switch workspace.Zone(snapshot, target, name) {
			case "evals", "interface":
				fmt.Fprintf(&prompt, "- %s\n", name)
			}
		}
	}
	command := invoke.Mapping(snapshot.Config["hooks"])["eval"]
	if override, exists := invoke.Mapping(target.Document.Frontmatter["hooks"])["eval"]; exists {
		command = override
	}
	report := ".fireweed/drill-eval.xml"
	args, err := invoke.Args(command, "{module}", target.Name, "{module_dir}", target.Dir, "{report}", report)
	if err != nil {
		return "", fmt.Errorf("local eval command: %w", err)
	}
	for i := range args {
		args[i] = invoke.Quote(args[i])
	}
	fmt.Fprintf(&prompt, "\nRun the target eval hook from this workspace with:\n\n```sh\nmkdir -p .fireweed && FIREWEED_BIN=%s %s\n```\n",
		invoke.Quote(bin), strings.Join(args, " "))
	prompt.WriteString("\nIn DRILL_NOTES.md, write separate ## Contradicted Decisions and ## Gaps sections. Each gap must name its nearest spec ID or \"no spec\" and the target module for a fixing change.\n")
	return prompt.String(), nil
}

func accept(ctx context.Context, snapshot project.Snapshot, target project.Module, scratch, base string, tracked []string, initial, submission map[string]workspace.File) (string, error) {
	if err := deleteDisposable(snapshot, target, scratch); err != nil {
		return "", err
	}
	paths := append([]string(nil), tracked...)
	for _, name := range slices.Sorted(maps.Keys(submission)) {
		if _, existed := initial[name]; existed || !disposableFile(snapshot, target, name) {
			continue
		}
		if err := submission[name].Put(scratch, name); err != nil {
			return "", err
		}
		paths = append(paths, name)
	}

	// Stage against the original tree so each retained ref contains one
	// reconstruction, including removals from the original disposable zone.
	if _, err := workspace.Git(ctx, scratch, "read-tree", base); err != nil {
		return "", err
	}
	if len(paths) > 0 {
		if _, err := workspace.Git(ctx, scratch, append([]string{"--literal-pathspecs", "add", "--force", "--all", "--"}, paths...)...); err != nil {
			return "", err
		}
	}
	tree, err := workspace.Git(ctx, scratch, "write-tree")
	if err != nil {
		return "", err
	}
	commit, err := workspace.Git(ctx, scratch, "commit-tree", strings.TrimSpace(tree), "-p", base, "-m", "Reconstruct "+target.Name)
	return strings.TrimSpace(commit), err
}
