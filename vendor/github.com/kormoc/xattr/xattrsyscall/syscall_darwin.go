package xattrsyscall

import (
    "syscall"
    "unsafe"
)

var _zero uintptr

func Getxattr(path string, attr string, dest []byte) (int, error) {
    var destPtr *byte
    var size int

    if dest != nil {
        destPtr = &dest[0]
        size = len(dest)
    }

    r0, _, e1 := syscall.Syscall6(syscall.SYS_GETXATTR, uintptr(unsafe.Pointer(syscall.StringBytePtr(path))), uintptr(unsafe.Pointer(syscall.StringBytePtr(attr))), uintptr(unsafe.Pointer(destPtr)), uintptr(size), uintptr(_zero), uintptr(_zero))
    return int(r0), e1
}

func Listxattr(path string, dest []byte) (int, error) {
    var destPtr *byte
    var size int

    if dest != nil {
        destPtr = &dest[0]
        size = len(dest)
    }

    r0, _, e1 := syscall.Syscall6(syscall.SYS_LISTXATTR, uintptr(unsafe.Pointer(syscall.StringBytePtr(path))), uintptr(unsafe.Pointer(destPtr)), uintptr(size), uintptr(_zero), 0, 0)
    return int(r0), e1
}

func Setxattr(path string, attr string, data []byte) error {
    _, _, e1 := syscall.Syscall6(syscall.SYS_SETXATTR, uintptr(unsafe.Pointer(syscall.StringBytePtr(path))), uintptr(unsafe.Pointer(syscall.StringBytePtr(attr))), uintptr(unsafe.Pointer(&data[0])), uintptr(len(data)), uintptr(_zero), uintptr(_zero))
    return e1
}

func Removexattr(path string, attr string) error {
    _, _, e1 := syscall.Syscall(syscall.SYS_REMOVEXATTR, uintptr(unsafe.Pointer(syscall.StringBytePtr(path))), uintptr(unsafe.Pointer(syscall.StringBytePtr(attr))), uintptr(_zero))
    return e1
}
