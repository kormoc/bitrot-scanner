package main

import "os"
import "time"

func workerChecksum(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	// Only process regular files
	if !info.Mode().IsRegular() {
		return nil
	}

	// Skip files if they've been modified too recently
	if info.ModTime().Unix() > (time.Now().Unix() - mtimeSettle) {
		Info.Println(path + ": Has been modified too recently. Skipping due to -mtimeSettle")
		return nil
	}

	if skipValidation && !missingChecksums(path) {
		Info.Println(path + ": Already has all checksums. Skipping due to -skipValidation")
		return nil
	}

	if skipCreate && missingChecksums(path) && checksumCount(path) == 0 {
		Info.Println(path + ": Missing all checksums. Skipping due to -skipCreate")
		return nil
	}

	checksumMTimePath := xattrRoot + "mtime"

	Info.Println(path + ": Processing...")

	modTime := info.ModTime().Unix()
	Trace.Printf("%v: mtime: %v\n", path, modTime)

	checksumMTime, _ := GetxattrInt64(path, checksumMTimePath)

	if checksumMTime == 0 {
		SetxattrInt64(path, checksumMTimePath, modTime)
	}

	// Validate that the file hasn't received a new mtime
	if checksumMTime != 0 && modTime > checksumMTime {
		Warn.Printf("%v has a mtime after checksums were generated!\n", path)
	}

	// Get hashes
	hashes, err := checksumPath(path)
	if err != nil {
		return err
	}

	for checksumAlgo := range hashes {
		checksumPath := xattrRoot + checksumAlgo
		checksumValue, _ := GetxattrString(path, checksumPath)

		// If the checksum is missing, just store it
		if len(checksumValue) == 0 {
			SetxattrString(path, checksumPath, hashes[checksumAlgo])
		} else {
			if hashes[checksumAlgo] != checksumValue {
				if modTime > checksumMTime && updateOnNewMTime {
					Warn.Printf("%v: Updating checksum due to updated mtime\n", path)
					SetxattrString(path, checksumPath, hashes[checksumAlgo])
					SetxattrInt64(path, checksumMTimePath, modTime)
				} else {
					if modTime < checksumMTime && updateOnNewMTime {
						Error.Printf("%v: Failed to update checksum due to mtime reversing\n", path)
					}
					Error.Printf("%v: CHECKSUM MISMATCH!\n\tComputed: %v\n\tExpected: %v\n", path, hashes[checksumAlgo], checksumValue)
				}
			}
		}
	}

	return nil
}
