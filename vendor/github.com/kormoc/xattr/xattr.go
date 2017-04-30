package xattr

import "github.com/kormoc/xattr/xattrsyscall"

func Remove(filePath string, xattrName string) error {
    return syscallErrorToXAttrError(xattrsyscall.Removexattr(filePath, xattrName))
}

func Has(filePath string, xattrName string) (bool, error) {
    _, err := GetBytes(filePath, xattrName)
    switch syscallErrorToXAttrError(err) {
        case nil:
            return true, nil
        case XAttrErrorAttributeNotFound:
            return false, nil
        case XAttrErrorNoDataAvailable:
            return false, nil
        case XAttrErrorResultTooLarge:
            return true, nil
        default:
            return false, err
    }
}
