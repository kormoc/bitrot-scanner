package main

import (
	"github.com/kormoc/xattr"
)

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

func GetCheckedTimeXattr(path string) int64 {
	mtime, err := GetxattrInt64(path, config.XattrRoot+"checkedtime")
	if xattr.XAttrErrorIsFatal(err) {
		Error.Fatalf("%v: CheckedTime Error: %v\n", path, err)
	}
	return mtime
}

func SetCheckedTimeXattr(path string, value int64) {
	err := SetxattrInt64(path, config.XattrRoot+"checkedtime", value)
	if xattr.XAttrErrorIsFatal(err) {
		Error.Fatalf("%v: CheckedTime Error: %v\n", path, err)
	}
}

func GetMTimeXattr(path string) int64 {
	mtime, err := GetxattrInt64(path, config.XattrRoot+"mtime")
	if xattr.XAttrErrorIsFatal(err) {
		Error.Fatalf("%v: MTime Error: %v\n", path, err)
	}
	return mtime
}

func SetMTimeXattr(path string, value int64) {
	err := SetxattrInt64(path, config.XattrRoot+"mtime", value)
	if xattr.XAttrErrorIsFatal(err) {
		Error.Fatalf("%v: MTime Error: %v\n", path, err)
	}
}

func GetChecksumXattr(path string, checksum ChecksumType) string {
	data, err := GetxattrString(path, config.XattrRoot+checksum.String())
	if xattr.XAttrErrorIsFatal(err) {
		Error.Fatalf("%v: %v Error: %v\n", path, checksum.String(), err)
	}
	return data
}

func SetChecksumXattr(path string, checksum ChecksumType, value string) {
	err := SetxattrString(path, config.XattrRoot+checksum.String(), value)
	if xattr.XAttrErrorIsFatal(err) {
		Error.Fatalf("%v: %v Error: %v\n", path, checksum.String(), err)
	}
}

func RemoveChecksumXattr(path string, checksum ChecksumType) {
	err := Removexattr(path, config.XattrRoot+checksum.String())
	if err != nil {
		Error.Fatalf("%v: %v Error: %v\n", path, checksum.String(), err)
	}
}

func RemoveTimeXattr(path string, timeName string) {
	err := Removexattr(path, config.XattrRoot+timeName)
	if err != nil {
		Error.Fatalf("%v: %v Error: %v\n", path, timeName, err)
	}
}
