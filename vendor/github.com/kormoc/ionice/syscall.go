// +build !linux

package ionice

func GetIOPriority(which int, who int) (prio uint32, err error) {
    return 0, NotImplementedError
}

func SetIOPriority(which int, who int, prio uint32) (err error) {
    return NotImplementedError
}
