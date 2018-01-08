package main

import (
	"os"
	"path/filepath"

	flag "github.com/ogier/pflag"
)

func unsortedJobs(startTime, endTime int64) {
	// Loop over the passed in directories and get all files
	for _, path := range flag.Args() {
		Info.Printf("Processing %v...\n", path)
		if err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.Mode()&os.ModeSymlink != 0 {
				return nil
			}

			// Only process regular files
			if !info.Mode().IsRegular() {
				return nil
			}

			j := newJob(path, info)

			if config.ResetXattrs {
				workerResetJobs <- j
			} else {
				workerStartJobs <- j
			}

			return nil
		}); err != nil {
			Error.Println(err)
		}
	}

}
