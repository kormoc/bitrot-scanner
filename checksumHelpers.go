package main

func checksumCount(path string) (count int) {
	for checksumAlgo := range checksumLookupTable {
		checksumPath := config.XattrRoot + checksumAlgo.String()
		checksumValue, _ := GetxattrString(path, checksumPath)
		if len(checksumValue) > 0 {
			count += 1
		}
	}
	return count
}
