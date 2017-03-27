package main

import "encoding/binary"
import "encoding/hex"
import "errors"
import "github.com/davecheney/xattr"
import "runtime"
import "strings"

func fixNameForLinux(name string) string {
    // Due to our xattr lib, all attributes on linux are prepended with "user." already
    // so strip this off for our ends if we're on linux

    var userPrefix = "user."

    if runtime.GOOS == "linux" && strings.HasPrefix(name, userPrefix) {
        return name[len(userPrefix):]
    }
    return name
}

func GetxattrHex(path string, name string) (string, error) {
    dataBytes, err := xattr.Getxattr(path, fixNameForLinux(name))
    if err != nil {
        return "", err
    }
    data := hex.EncodeToString(dataBytes)
    Trace.Printf("%v: Get: %v: %v\n", path, name, data)
    return data, nil
}

func SetxattrHex(path string, name string, data string) error {
    dataBytes, err := hex.DecodeString(data)
    if err != nil {
        return err
    }
    Trace.Printf("%v: Set: %v: %v\n", path, name, dataBytes)
    return xattr.Setxattr(path, fixNameForLinux(name), dataBytes)
}

func GetxattrInt64(path string, name string) (int64, error) {
    dataBytes, err := xattr.Getxattr(path, fixNameForLinux(name))
    if err != nil {
        return 0, err
    }
    data, n := binary.Varint(dataBytes)
    if n <= 0 {
        return 0, errors.New("Binary conversion failed")
    }
    Trace.Printf("%v: Get: %v: %v\n", path, name, data)
    return data, nil
}

func SetxattrInt64(path string, name string, data int64) error {
    dataBytes := make([]byte, 8)
    binary.PutVarint(dataBytes, data)
    Trace.Printf("%v: Set: %v: %v\n", path, name, dataBytes)

    return xattr.Setxattr(path, fixNameForLinux(name), dataBytes)
}

func Removexattr(path string, name string) error {
    if !Hasxattr(path, name) {
        return nil
    }
    Info.Printf("%v: Removing: %v", path, name)
    return xattr.Removexattr(path, fixNameForLinux(name))
}

func Hasxattr(path string, name string) bool {
    data, _ := xattr.Getxattr(path, fixNameForLinux(name))
    if len(data) != 0 {
        return true
    }
    return false
}