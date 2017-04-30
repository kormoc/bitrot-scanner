package xattr

import "errors"
import "os"

var XAttrErrorAttributeNotFound = errors.New("attribute not found")
var XAttrErrorNoDataAvailable = errors.New("no data available")
var XAttrErrorResultTooLarge = errors.New("result too large")

func syscallErrorToXAttrError(err error) error {
    if err == nil {
        return nil
    }
    switch err.Error() {
        case "attribute not found":
            return XAttrErrorAttributeNotFound
        case "errno 0":
            return nil
        case "no data available":
            return XAttrErrorNoDataAvailable
        case "result too large":
            return XAttrErrorResultTooLarge
        case "permission denied":
            return os.ErrPermission
        default:
            return err
    }
}

func XAttrErrorIsFatal(err error) bool {
    switch err {
        case nil:
            return false
        case XAttrErrorAttributeNotFound:
            return false
        case XAttrErrorNoDataAvailable:
            return false
        case XAttrErrorResultTooLarge:
            return true
        case os.ErrPermission:
            return true
        default:
            return true
    }
}
