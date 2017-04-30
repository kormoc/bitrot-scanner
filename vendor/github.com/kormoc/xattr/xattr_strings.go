package xattr

import "strings"

func GetString(filePath string, xattrName string) (string, error) {
    value, err := GetBytes(filePath, xattrName)
    return string(value), syscallErrorToXAttrError(err)
}

func ListStrings(filePath string) ([]string, error) {
    value, err :=  ListBytes(filePath)
    list := strings.Split(string(value), "\x00")
    // Slice off the last element if it's empty
    if list[len(list)-1] == "" {
        list = list[:len(list)-1]
    }
    return list, syscallErrorToXAttrError(err)
}

func SetString(filePath string, xattrName string, value string) error {
    return syscallErrorToXAttrError(SetBytes(filePath, xattrName, []byte(value)))
}
