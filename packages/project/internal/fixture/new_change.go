package fixture

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"slices"
	"syscall"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/internal/source"
	"github.com/lurkshark/fireweed/packages/project/model"
)

var validSlug = regexp.MustCompile(`^[a-z0-9-]+$`)

func NewChange(root, module, slug string) (string, error) {
	if !validSlug.MatchString(slug) {
		return "", source.Failure("U01", slug, fmt.Errorf("invalid change slug %q", slug))
	}
	snapshot, err := Load(root)
	if err != nil {
		return "", err
	}
	if !slices.ContainsFunc(snapshot.Modules, func(candidate model.Module) bool { return candidate.Name == module }) {
		return "", source.Failure("U02", module, fmt.Errorf("unknown module %q", module))
	}
	next := 1
	for _, dir := range []string{"changes", "changes/archive"} {
		files, err := source.Files(root, dir)
		if dir == "changes" && errors.Is(err, syscall.ENOTDIR) {
			// There is no queue or archive to number. MkdirAll below
			// reports the obstruction as a failure to create the file.
			break
		}
		if err != nil {
			return "", err
		}
		for _, file := range files {
			number, existing, ok := source.ChangeName(file)
			if !ok {
				continue
			}
			if existing == slug {
				return "", source.Failure("C13", slug, fmt.Errorf("change slug %q already exists in %s/%s", slug, dir, file))
			}
			next = max(next, number+1)
		}
	}
	relative := fmt.Sprintf("changes/%04d-%s.md", next, slug)
	text, err := contracts.RenderDocument(contracts.Document{
		Frontmatter: map[string]any{"module": module},
		Body:        "\n## Motivation\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n",
	})
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(source.Path(root, "changes"), 0755); err != nil {
		return "", fmt.Errorf("create %s: %w", relative, err)
	}
	full := source.Path(root, relative)
	file, err := os.OpenFile(full, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return "", fmt.Errorf("create %s: %w", relative, err)
	}
	_, err = file.WriteString(text)
	closeErr := file.Close()
	if err == nil {
		err = closeErr
	}
	if err != nil {
		_ = os.Remove(full)
		return "", fmt.Errorf("create %s: %w", relative, err)
	}
	return relative, nil
}
