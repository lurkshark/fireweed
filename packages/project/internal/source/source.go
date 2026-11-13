package source

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
)

func Failure(code, subject string, cause error) error {
	return contracts.CodedError{Code: code, Subject: subject, Cause: cause}
}

func FileError(code, path string, cause error) error {
	return Failure(code, path, fmt.Errorf("%s: %w", path, cause))
}

func Path(root, relative string) string {
	return filepath.Join(root, filepath.FromSlash(relative))
}

var changeName = regexp.MustCompile(`^([0-9]{4,})-([a-z0-9-]+)\.md$`)

func ChangeName(name string) (number int, slug string, ok bool) {
	fields := changeName.FindStringSubmatch(name)
	if fields == nil {
		return 0, "", false
	}
	number, err := strconv.Atoi(fields[1])
	return number, fields[2], err == nil
}

func Files(root, dir string) ([]string, error) {
	entries, err := os.ReadDir(Path(root, dir))
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, FileError("C03", dir, err)
	}
	var files []string
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}
	return files, nil
}

func Items(doc contracts.Document, section string) ([]contracts.Item, error) {
	body, _ := contracts.Section(doc.Body, section)
	return contracts.ParseItems(body)
}

func DocumentError(path string, cause error) error {
	code, _, ok := contracts.ErrorCode(cause)
	if !ok {
		code = "C03"
	}
	return FileError(code, path, cause)
}

func Git(root string, args ...string) (string, error) {
	result, err := gitstore.Run(context.Background(), gitstore.Request{
		Args:    append([]string{"git", "--no-optional-locks", "-C", root}, args...),
		Env:     map[string]string{"GIT_NO_LAZY_FETCH": "1", "GIT_TERMINAL_PROMPT": "0"},
		Timeout: 30 * time.Second,
	})
	if err != nil {
		return "", fmt.Errorf("git %s: %w: %s", args[0], err, strings.TrimSpace(result.Stderr))
	}
	return result.Stdout, nil
}
