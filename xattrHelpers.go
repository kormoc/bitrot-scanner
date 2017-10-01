package main

import "github.com/kormoc/xattr"

func GetxattrInt64(path string, name string) (int64, error) {
	return xattr.GetStringInt64(path, name)
}

func SetxattrInt64(path string, name string, data int64) error {
	return xattr.SetStringInt64(path, name, data)
}

func GetxattrString(path string, name string) (string, error) {
	return xattr.GetString(path, name)
}

func SetxattrString(path string, name string, data string) error {
	return xattr.SetString(path, name, data)
}

func Removexattr(path string, name string) error {
	has, err := xattr.Has(path, name)
	if !has {
		return nil
	}
	if err != nil {
		return err
	}
	Info.Printf("%v: Removing: %v", path, name)
	return xattr.Remove(path, name)
}

// bitrot-scanner specific helpers

func GetMTimeXattr(path string) int64 {
	mtime, err := GetxattrInt64(path, xattrRoot+"mtime")
	if xattr.XAttrErrorIsFatal(err) {
		Error.Fatalf("%v: MTime Error: %v\n", path, err)
	}
	return mtime
}

func SetMTimeXattr(path string, value int64) {
	err := SetxattrInt64(path, xattrRoot+"mtime", value)
	if xattr.XAttrErrorIsFatal(err) {
		Error.Fatalf("%v: MTime Error: %v\n", path, err)
	}
}

func GetChecksumXattr(path string, checksum string) string {
	data, err := GetxattrString(path, xattrRoot+checksum)
	if xattr.XAttrErrorIsFatal(err) {
		Error.Fatalf("%v: %v Error: %v\n", path, checksum, err)
	}
	return data
}

func SetChecksumXattr(path string, checksum string, value string) {
	err := SetxattrString(path, xattrRoot+checksum, value)
	if xattr.XAttrErrorIsFatal(err) {
		Error.Fatalf("%v: %v Error: %v\n", path, checksum, err)
	}
}

func RemoveChecksumXattr(path string, checksum string) {
	err := Removexattr(path, xattrRoot+checksum)
	if err != nil {
		Error.Fatalf("%v: %v Error: %v\n", path, checksum, err)
	}
}
