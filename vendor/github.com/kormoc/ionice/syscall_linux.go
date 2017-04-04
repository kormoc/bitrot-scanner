// +build linux

package ionice

import "errors"
import "fmt"
import "syscall"

func GetIOPriority(which int, who int) (prio uint32, err error) {
    r0, _, e1 := syscall.Syscall(syscall.SYS_IOPRIO_GET, uintptr(which), uintptr(who), 0)
    prio = uint32(r0)
    if e1 != 0 {
        err = errors.New(fmt.Sprintf("Received error number %v", e1))
    }
    return
}

func SetIOPriority(which int, who int, prio uint32) (err error) {
    _, _, e1 := syscall.Syscall(syscall.SYS_IOPRIO_SET, uintptr(which), uintptr(who), uintptr(prio))
    if e1 != 0 {
        err = errors.New(fmt.Sprintf("Received error number %v", e1))
    }
    return
}
