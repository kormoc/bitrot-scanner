package xattrsyscall

import (
    "syscall"
)

func Getxattr(path string, attr string, dest []byte) (sz int, err error) {
    return syscall.Getxattr(path, attr, dest)
}

func Listxattr(path string, dest []byte) (sz int, err error) {
    return syscall.Listxattr(path, dest)
}

func Removexattr(path string, attr string) (err error) {
    return syscall.Removexattr(path, attr)
}

func Setxattr(path string, attr string, data []byte) (err error) {
    return syscall.Setxattr(path, attr, data, 0)
}
