package xattrsyscall

import (
    "syscall"
    "unsafe"
)

// Taken from https://golang.org/src/syscall/zsyscall_linux_amd64.go

var _zero uintptr

// Do the interface allocations only once for common
// Errno values.
var (
    errEAGAIN error = syscall.EAGAIN
    errEINVAL error = syscall.EINVAL
    errENOENT error = syscall.ENOENT
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
    switch e {
    case 0:
        return nil
    case syscall.EAGAIN:
        return errEAGAIN
    case syscall.EINVAL:
        return errEINVAL
    case syscall.ENOENT:
        return errENOENT
    }
    return e
}

// BSD uses namespaces. Hardcode the user namespace
const EXTATTR_NAMESPACE_USER = 1


func Getxattr(path string, attr string, dest []byte) (sz int, err error) {
    var _p0 *byte
    _p0, err = syscall.BytePtrFromString(path)
    if err != nil {
        return
    }
    var _p1 *byte
    _p1, err = syscall.BytePtrFromString(attr)
    if err != nil {
        return
    }
    var _p2 unsafe.Pointer
    if len(dest) > 0 {
        _p2 = unsafe.Pointer(&dest[0])
    } else {
        _p2 = unsafe.Pointer(&_zero)
    }
    r0, _, e1 := syscall.Syscall6(syscall.SYS_EXTATTR_GET_FILE, uintptr(unsafe.Pointer(_p0)), uintptr(EXTATTR_NAMESPACE_USER), uintptr(unsafe.Pointer(_p1)), uintptr(_p2), uintptr(len(dest)), 0)
    sz = int(r0)
    if e1 != 0 {
        err = errnoErr(e1)
    }
    return
}

func Listxattr(path string, dest []byte) (sz int, err error) {
    var _p0 *byte
    _p0, err = syscall.BytePtrFromString(path)
    if err != nil {
        return
    }
    var _p1 unsafe.Pointer
    if len(dest) > 0 {
        _p1 = unsafe.Pointer(&dest[0])
    } else {
        _p1 = unsafe.Pointer(&_zero)
    }
    r0, _, e1 := syscall.Syscall6(syscall.SYS_EXTATTR_LIST_FILE, uintptr(unsafe.Pointer(_p0)), uintptr(EXTATTR_NAMESPACE_USER), uintptr(_p1), uintptr(len(dest)), 0, 0)
    sz = int(r0)
    if e1 != 0 {
        err = errnoErr(e1)
    }
    return
}

func Setxattr(path string, attr string, data []byte, flags int) (err error) {
    var _p0 *byte
    _p0, err = syscall.BytePtrFromString(path)
    if err != nil {
        return
    }
    var _p1 *byte
    _p1, err = syscall.BytePtrFromString(attr)
    if err != nil {
        return
    }
    var _p2 unsafe.Pointer
    if len(data) > 0 {
        _p2 = unsafe.Pointer(&data[0])
    } else {
        _p2 = unsafe.Pointer(&_zero)
    }
    _, _, e1 := syscall.Syscall6(syscall.SYS_EXTATTR_SET_FILE, uintptr(unsafe.Pointer(_p0)), uintptr(EXTATTR_NAMESPACE_USER), uintptr(unsafe.Pointer(_p1)), uintptr(_p2), uintptr(len(data)), uintptr(flags))
    if e1 != 0 {
        err = errnoErr(e1)
    }
    return
}

func Removexattr(path string, attr string) (err error) {
    var _p0 *byte
    _p0, err = syscall.BytePtrFromString(path)
    if err != nil {
        return
    }
    var _p1 *byte
    _p1, err = syscall.BytePtrFromString(attr)
    if err != nil {
        return
    }
    _, _, e1 := syscall.Syscall(syscall.SYS_EXTATTR_DELETE_FILE, uintptr(unsafe.Pointer(_p0)), uintptr(EXTATTR_NAMESPACE_USER), uintptr(unsafe.Pointer(_p1)))
    if e1 != 0 {
        err = errnoErr(e1)
    }
    return
}
