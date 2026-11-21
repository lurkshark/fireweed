package lifecycle

import (
	"errors"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
)

type auditSection struct {
	before    string
	edit      string
	separator string
	after     string
}

func ApplyAuditEdit(original, submitted string) (string, error) {
	old, err := splitAudit(original)
	if err != nil {
		return "", err
	}
	edit, err := splitAudit(submitted)
	if err != nil {
		return "", err
	}
	oldBefore, _ := splitTrailingBlankLines(old.before)
	newBefore, _ := splitTrailingBlankLines(edit.before)
	switch {
	case old.edit != "" && edit.edit != "":
		if oldBefore == newBefore && old.after == edit.after {
			return old.before + joinAtLineBoundary(edit.edit, old.separator) + old.after, nil
		}
	case edit.edit != "":
		// At EOF, a new heading can require terminating the preceding line.
		if edit.after == "" && !strings.HasSuffix(original, "\n") &&
			(newBefore == original+"\n" || newBefore == original+"\r\n") {
			original = newBefore
		}
		if gap, ok := blankGap(original, newBefore, edit.after); ok {
			return joinAtLineBoundary(newBefore+gap, edit.edit) + edit.after, nil
		}
	case old.edit != "":
		if _, ok := blankGap(submitted, oldBefore, old.after); ok {
			return old.before + old.separator + old.after, nil
		}
	case original == submitted:
		return original, nil
	}
	return "", errors.New("only the Audit section may change")
}

func splitAudit(body string) (auditSection, error) {
	content, found := contracts.Section(body, "Audit")
	if !found {
		return auditSection{}, nil
	}

	// Locate the heading and body using the same Markdown grammar as Section.
	marker := "\x00"
	for strings.Contains(body, marker) {
		marker += "\x00"
	}
	prefix, after, _ := strings.Cut(contracts.ReplaceSection(body, "Audit", marker), marker)
	if _, duplicate := contracts.Section(after, "Audit"); duplicate {
		return auditSection{}, errors.New("multiple Audit sections")
	}
	heading := strings.LastIndex(strings.TrimSuffix(prefix, "\n"), "\n") + 1
	content, separator := splitTrailingBlankLines(content)
	return auditSection{before: prefix[:heading], edit: prefix[heading:] + content, separator: separator, after: after}, nil
}

func splitTrailingBlankLines(text string) (string, string) {
	end := len(text)
	for end > 0 {
		line := strings.TrimSuffix(text[:end], "\n")
		start := strings.LastIndexByte(line, '\n') + 1
		if strings.Trim(strings.TrimSuffix(line[start:], "\r"), " \t") != "" {
			break
		}
		end = start
	}
	return text[:end], text[end:]
}

func blankGap(body, before, after string) (string, bool) {
	rest, hasBefore := strings.CutPrefix(body, before)
	gap, hasAfter := strings.CutSuffix(rest, after)
	// The suffix must start on its own line. Its indentation is protected.
	if !hasBefore || !hasAfter || after != "" && gap != "" && !strings.HasSuffix(gap, "\n") {
		return "", false
	}
	nonblank, _ := splitTrailingBlankLines(gap)
	return gap, nonblank == ""
}

func joinAtLineBoundary(before, after string) string {
	if before != "" && after != "" && !strings.HasSuffix(before, "\n") {
		return before + "\n" + after
	}
	return before + after
}
