package start

import (
	"bytes"
	"fmt"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/internal/claims"
	"github.com/lurkshark/fireweed/packages/workflow/model"
	"os"
	"path/filepath"
	"strings"
)

type node struct {
	path  string
	bytes []byte
	mode  os.FileMode
	link  string
}

// Refuse symlink ancestors even when they currently point within the checkout.
// A disposable leaf link is an object to unlink, never a source to follow.
func safeLeaf(root, relative string) (string, error) {
	if relative == "" || filepath.IsAbs(relative) || strings.ContainsAny(relative, "\\\x00") || filepath.ToSlash(filepath.Clean(relative)) != relative || relative == ".." || strings.HasPrefix(relative, "../") {
		return "", fmt.Errorf("unsafe path %q", relative)
	}
	path := root
	parts := strings.Split(relative, "/")
	for _, part := range parts[:len(parts)-1] {
		path = filepath.Join(path, part)
		info, e := os.Lstat(path)
		if e != nil {
			return "", e
		}
		if !info.IsDir() || info.Mode()&os.ModeSymlink != 0 {
			return "", fmt.Errorf("unsafe ancestor %s", path)
		}
	}
	return filepath.Join(root, filepath.FromSlash(relative)), nil
}
func readNode(root, path string, links bool) (node, error) {
	n := node{path: path}
	full, e := safeLeaf(root, path)
	if e != nil {
		return n, e
	}
	info, e := os.Lstat(full)
	if e != nil {
		return n, e
	}
	n.mode = info.Mode()
	if info.Mode()&os.ModeSymlink != 0 && links {
		n.link, e = os.Readlink(full)
		return n, e
	}
	if !info.Mode().IsRegular() {
		return n, fmt.Errorf("restore a regular source at %s", path)
	}
	n.bytes, e = os.ReadFile(full)
	return n, e
}
func sameNode(root, path string, want node) bool {
	got, e := readNode(root, path, want.link != "")
	return e == nil && got.mode == want.mode && got.link == want.link && bytes.Equal(got.bytes, want.bytes)
}
func writeAtomic(root, path string, data []byte, mode os.FileMode) error {
	full, e := safeLeaf(root, path)
	if e != nil {
		return e
	}
	f, e := os.CreateTemp(filepath.Dir(full), ".fireweed-start-")
	if e != nil {
		return e
	}
	name := f.Name()
	defer os.Remove(name)
	if _, e = f.Write(data); e != nil {
		f.Close()
		return e
	}
	if e = f.Chmod(mode); e != nil {
		f.Close()
		return e
	}
	if e = f.Close(); e != nil {
		return e
	}
	return os.Rename(name, full)
}

type transaction struct {
	root, path, written string
	original            node
	claim               model.Claim
	didWrite            bool
	deleted             []node
}

func (tx *transaction) owned() error {
	actual, e := claims.Load(tx.root, tx.claim.Change)
	if e != nil || actual != tx.claim {
		return problem("P02: claim ownership changed during start; recovery must preserve the replacement owner: %v", e)
	}
	return nil
}
func (tx *transaction) ownsSource() bool {
	return sameNode(tx.root, tx.path, node{path: tx.path, bytes: []byte(tx.written), mode: tx.original.mode})
}
func (tx *transaction) rollback(cause error) error {
	var recovery []string
	for i := len(tx.deleted) - 1; i >= 0; i-- {
		n := tx.deleted[i]
		full, e := safeLeaf(tx.root, n.path)
		if e != nil {
			recovery = append(recovery, n.path+": "+e.Error())
			continue
		}
		if _, e = os.Lstat(full); !os.IsNotExist(e) {
			recovery = append(recovery, n.path+": recreated path preserved")
			continue
		}
		if n.link != "" {
			e = os.Symlink(n.link, full)
		} else {
			var f *os.File
			f, e = os.OpenFile(full, os.O_WRONLY|os.O_CREATE|os.O_EXCL, n.mode)
			if e == nil {
				_, e = f.Write(n.bytes)
				if e == nil {
					e = f.Chmod(n.mode)
				}
				ce := f.Close()
				if e == nil {
					e = ce
				}
			}
		}
		if e != nil {
			recovery = append(recovery, n.path+": "+e.Error())
		}
	}
	if tx.didWrite {
		if tx.ownsSource() {
			if e := writeAtomic(tx.root, tx.path, tx.original.bytes, tx.original.mode); e != nil {
				recovery = append(recovery, tx.path+": "+e.Error())
			}
		} else {
			recovery = append(recovery, tx.path+": concurrently changed source preserved")
		}
	}
	if tx.claim.Object != "" {
		if e := tx.owned(); e != nil {
			recovery = append(recovery, e.Error())
		} else if e = claims.Release(tx.root, tx.claim.Change, tx.claim.Object); e != nil {
			recovery = append(recovery, e.Error())
		}
	}
	if len(recovery) > 0 {
		return fmt.Errorf("%w; incomplete recovery: %s. Check ownership and restore only your own work before retry.", cause, strings.Join(recovery, "; "))
	}
	return cause
}
func (tx *transaction) regenerate(local localGit, s project.Snapshot, ch contracts.Change) error {
	before, e := project.Project(s, ch.Slug)
	if e != nil {
		return problem("cannot project regeneration: %v", e)
	}
	var target contracts.Module
	for _, m := range before.Modules {
		if m.Name == ch.Module {
			target, e = project.ApplyChange(m, ch)
			if e != nil {
				return problem("cannot apply regeneration zones: %v", e)
			}
		}
	}
	inventory, e := local.raw("ls-files", "-z", "--cached")
	if e != nil {
		return e
	}
	if inventory != "" && !strings.HasSuffix(inventory, "\x00") {
		return problem("incomplete tracked file inventory")
	}
	// Validate and capture the complete deletion set before removing any leaf.
	var pending []node
	for _, path := range strings.Split(inventory, "\x00") {
		if path == "" {
			continue
		}
		zone, e := contracts.ZoneFor(path, target.Dir, target.Zones, s.Config.Protected)
		if e != nil {
			return problem("cannot classify %s: %v", path, e)
		}
		if zone != contracts.DisposableZone {
			continue
		}
		n, e := readNode(tx.root, path, true)
		if e != nil {
			return problem("cannot capture disposable file %s: %v", path, e)
		}
		pending = append(pending, n)
	}
	for _, n := range pending {
		if e = tx.owned(); e != nil {
			return e
		}
		if !tx.ownsSource() {
			return problem("selected source changed during regeneration; recovery preserved replacement")
		}
		if !sameNode(tx.root, n.path, n) {
			return problem("disposable path %s changed concurrently; recover ownership before retry", n.path)
		}
		full, e := safeLeaf(tx.root, n.path)
		if e != nil {
			return problem("unsafe disposable path %s: %v", n.path, e)
		}
		if e = os.Remove(full); e != nil {
			return problem("cannot delete disposable file %s: %v", n.path, e)
		}
		tx.deleted = append(tx.deleted, n)
	}
	return nil
}
