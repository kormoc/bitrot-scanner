package xattr

import "strconv"

func GetStringInt64(filePath string, xattrName string) (int64, error) {
    value, err := GetString(filePath, xattrName)
    if err != nil {
        return 0, err
    }
    data, err := strconv.ParseInt(value, 10, 64)
    if err != nil {
        return 0, err
    }
    return data, nil
}

func SetStringInt64(filePath string, xattrName string, value int64) error {
    return SetString(filePath, xattrName, strconv.FormatInt(value, 10))
}
