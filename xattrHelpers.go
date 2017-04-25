package main

import "github.com/davecheney/xattr"
import "runtime"
import "strconv"
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

func GetxattrInt64(path string, name string) (int64, error) {
	dataStr, err := GetxattrString(path, name)
	if err != nil {
		return 0, err
	}
	data, err := strconv.ParseInt(dataStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return data, nil
}

func SetxattrInt64(path string, name string, data int64) error {
	return SetxattrString(path, name, strconv.FormatInt(data, 10))
}

func GetxattrString(path string, name string) (string, error) {
	dataBytes, err := xattr.Getxattr(path, fixNameForLinux(name))
	if err != nil {
		return "", err
	}
	data := string(dataBytes)
	Trace.Printf("%v: Get: %v: %v\n", path, name, data)
	return data, nil
}

func SetxattrString(path string, name string, data string) error {
	Trace.Printf("%v: Set: %v: %v\n", path, name, data)

	return xattr.Setxattr(path, fixNameForLinux(name), []byte(data))
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


// bitrot-scanner specific helpers

func GetMTimeXattr(path string) int64 {
    mtime, err := GetxattrInt64(path, xattrRoot+"mtime")
    if err != nil {
        if !strings.HasSuffix(err.Error(), "attribute not found") && !strings.HasSuffix(err.Error(), "errno 0"){
            Error.Fatalf("%v: MTime Error: %v\n", path, err)
        }
    }
    return mtime
}

func SetMTimeXattr(path string, value int64) {
    err := SetxattrInt64(path, xattrRoot+"mtime", value)
    if err != nil {
        if !strings.HasSuffix(err.Error(), "errno 0"){
            Error.Fatalf("%v: MTime Error: %v\n", path, err)
        }
    }
}

func GetChecksumXattr(path string, checksum string) string {
    data, err := GetxattrString(path, xattrRoot+checksum)
    if err != nil {
        if strings.HasSuffix(err.Error(), "attribute not found"){
            return ""
        }
        if !strings.HasSuffix(err.Error(), "errno 0") {
            Error.Fatalf("%v: %v Error: %v\n", path, checksum, err)
        }
    }
    return data
}

func SetChecksumXattr(path string, checksum string, value string) {
    err := SetxattrString(path, xattrRoot+checksum, value)
    if err != nil {
        if !strings.HasSuffix(err.Error(), "errno 0"){
            Error.Fatalf("%v: %v Error: %v\n", path, checksum, err)
        }
    }
}

func RemoveChecksumXattr(path string, checksum string) {
    err := Removexattr(path, xattrRoot+checksum)
    if err != nil {
        Error.Fatalf("%v: %v Error: %v\n", path, checksum, err)
    }
}
