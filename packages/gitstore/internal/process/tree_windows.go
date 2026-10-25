package process

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"unsafe"
)

var kernel = syscall.NewLazyDLL("kernel32.dll")
var createJob = kernel.NewProc("CreateJobObjectW")
var configureJob = kernel.NewProc("SetInformationJobObject")
var assignJob = kernel.NewProc("AssignProcessToJobObject")
var terminateJob = kernel.NewProc("TerminateJobObject")
var firstThread = kernel.NewProc("Thread32First")
var nextThread = kernel.NewProc("Thread32Next")
var openThread = kernel.NewProc("OpenThread")
var resumeThread = kernel.NewProc("ResumeThread")

type basicLimits struct {
	ProcessTime, JobTime         int64
	Flags                        uint32
	MinWorkingSet, MaxWorkingSet uintptr
	ActiveProcesses              uint32
	Affinity                     uintptr
	Priority, Scheduling         uint32
}
type jobLimits struct {
	Basic                                                      basicLimits
	IO                                                         [6]uint64
	ProcessMemory, JobMemory, PeakProcessMemory, PeakJobMemory uintptr
}
type threadEntry struct {
	Size, Usage, ID, Owner uint32
	Priority, Delta        int32
	Flags                  uint32
}
type jobControl struct{ handle syscall.Handle }

func windowsError(operation string, err error) error {
	if err == nil || err == syscall.Errno(0) {
		return fmt.Errorf("%s failed", operation)
	}
	return fmt.Errorf("%s: %w", operation, err)
}
func newTree(command *exec.Cmd) (treeControl, error) {
	handle, _, err := createJob.Call(0, 0)
	if handle == 0 {
		return nil, windowsError("CreateJobObject", err)
	}
	control := &jobControl{handle: syscall.Handle(handle)}
	limits := jobLimits{Basic: basicLimits{Flags: 0x2000}} // JOB_OBJECT_LIMIT_KILL_ON_JOB_CLOSE
	ok, _, err := configureJob.Call(handle, 9, uintptr(unsafe.Pointer(&limits)), unsafe.Sizeof(limits))
	if ok == 0 {
		control.close()
		return nil, windowsError("SetInformationJobObject", err)
	}
	// Assign before the initial thread runs, so all descendants inherit the job.
	command.SysProcAttr = &syscall.SysProcAttr{CreationFlags: 0x4} // CREATE_SUSPENDED
	return control, nil
}
func (control *jobControl) started(_ context.Context, command *exec.Cmd) error {
	process, err := syscall.OpenProcess(0x101, false, uint32(command.Process.Pid)) // SET_QUOTA | TERMINATE
	if err != nil {
		return err
	}
	defer syscall.CloseHandle(process)
	ok, _, err := assignJob.Call(uintptr(control.handle), uintptr(process))
	if ok == 0 {
		return windowsError("AssignProcessToJobObject", err)
	}
	snapshot, err := syscall.CreateToolhelp32Snapshot(4, 0) // TH32CS_SNAPTHREAD
	if err != nil {
		return err
	}
	defer syscall.CloseHandle(snapshot)
	entry := threadEntry{}
	entry.Size = uint32(unsafe.Sizeof(entry))
	ok, _, err = firstThread.Call(uintptr(snapshot), uintptr(unsafe.Pointer(&entry)))
	resumed := false
	for ok != 0 {
		if entry.Owner == uint32(command.Process.Pid) {
			thread, _, openErr := openThread.Call(2, 0, uintptr(entry.ID)) // THREAD_SUSPEND_RESUME
			if thread == 0 {
				return windowsError("OpenThread", openErr)
			}
			count, _, resumeErr := resumeThread.Call(thread)
			_ = syscall.CloseHandle(syscall.Handle(thread))
			if uint32(count) == ^uint32(0) {
				return windowsError("ResumeThread", resumeErr)
			}
			resumed = true
		}
		entry.Size = uint32(unsafe.Sizeof(entry))
		ok, _, err = nextThread.Call(uintptr(snapshot), uintptr(unsafe.Pointer(&entry)))
	}
	if !resumed {
		return windowsError("find suspended process thread", err)
	}
	return nil
}
func (control *jobControl) stop(command *exec.Cmd) error {
	if command.Process == nil {
		return os.ErrProcessDone
	}
	ok, _, err := terminateJob.Call(uintptr(control.handle), 1)
	if ok == 0 {
		return windowsError("TerminateJobObject", err)
	}
	return nil
}
func (control *jobControl) close() {
	if control.handle != 0 {
		_ = syscall.CloseHandle(control.handle)
		control.handle = 0
	}
}
