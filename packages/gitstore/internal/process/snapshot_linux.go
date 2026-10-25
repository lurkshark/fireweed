package process

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func nativeProcess(pid int) (processRef, error) {
	raw, err := os.ReadFile("/proc/" + strconv.Itoa(pid) + "/stat")
	if err != nil {
		return processRef{}, err
	}
	// comm may contain spaces or parentheses; the following fields do not.
	end := strings.LastIndexByte(string(raw), ')')
	if end < 0 {
		return processRef{}, fmt.Errorf("invalid process metadata for %d", pid)
	}
	fields := strings.Fields(string(raw[end+1:]))
	if len(fields) < 20 {
		return processRef{}, fmt.Errorf("incomplete process metadata for %d", pid)
	}
	parent, err := strconv.Atoi(fields[1])
	if err != nil {
		return processRef{}, err
	}
	born, err := strconv.ParseUint(fields[19], 10, 64)
	if err != nil {
		return processRef{}, err
	}
	return processRef{pid: pid, parent: parent, birth: [2]uint64{born, 0}, stopped: fields[0] == "T" || fields[0] == "t", dead: fields[0] == "Z" || fields[0] == "X"}, nil
}

func nativeChildren(pid int) ([]int, error) {
	root := "/proc/" + strconv.Itoa(pid) + "/task"
	threads, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}
	result, seen := []int{}, map[int]bool{}
	// Children can be started by any OS thread, not just the leader.
	for _, thread := range threads {
		if _, err := strconv.Atoi(thread.Name()); err != nil {
			continue
		}
		raw, err := os.ReadFile(filepath.Join(root, thread.Name(), "children"))
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return nil, err
		}
		for _, field := range strings.Fields(string(raw)) {
			child, err := strconv.Atoi(field)
			if err != nil || child <= 0 {
				return nil, fmt.Errorf("invalid child process inventory for %d", pid)
			}
			if !seen[child] {
				seen[child] = true
				result = append(result, child)
			}
		}
	}
	return result, nil
}
