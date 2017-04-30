package xattr

import "github.com/kormoc/xattr/xattrsyscall"

func GetBytes(filePath string, xattrName string) ([]byte, error) {
    size, err := xattrsyscall.Getxattr(filePath, xattrName, nil)
    err = syscallErrorToXAttrError(err)
    if err != nil {
        return nil, err
    }
    buf := make([]byte, size)
    read, err := xattrsyscall.Getxattr(filePath, xattrName, buf)
    return buf[:read], syscallErrorToXAttrError(err)
}

func ListBytes(filePath string) ([]byte, error) {
    size, err := xattrsyscall.Listxattr(filePath, nil)
    err = syscallErrorToXAttrError(err)
    if err != nil {
        return nil, err
    }
    buf := make([]byte, size)
    read, err := xattrsyscall.Listxattr(filePath, buf)
    return buf[:read], syscallErrorToXAttrError(err)
}

func SetBytes(filePath string, xattrName string, value []byte) error {
    return syscallErrorToXAttrError(xattrsyscall.Setxattr(filePath, xattrName, value))
}
