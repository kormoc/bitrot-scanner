package main

import "os"

func workerReset(path string, info os.FileInfo, err error) error {
    if err != nil {
        return err
    }

    for checksumAlgo := range checksumLookupTable {
        checksumPath := xattrRoot + checksumAlgo

        if err := Removexattr(path, checksumPath); err != nil {
            return err
        }
    }

    // Also clean up mtimes
    if err := Removexattr(path, xattrRoot+"mtime"); err != nil {
        return err
    }

    return nil
}
