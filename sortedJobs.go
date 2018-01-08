package main

import (
	"os"
	"path/filepath"
	"sort"
	"time"

	flag "github.com/ogier/pflag"
)

func sortedJobs(startTime, endTime int64) {

	var allJobs []job

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
			allJobs = append(allJobs, j)
			return nil
		}); err != nil {
			Error.Println(err)
		}
	}

	// Sort the slice so we check the oldest first
	Info.Printf("Sorting paths...\n")
	sort.Slice(allJobs,
		func(i, j int) bool {
			if allJobs[i].checkedTime == allJobs[j].checkedTime {
				return allJobs[i].path.String() < allJobs[j].path.String()
			}
			return allJobs[i].checkedTime < allJobs[j].checkedTime
		})

	// Loop over the passed in directories and hash and/or validate
	Info.Printf("Starting jobs...\n")
	for _, j := range allJobs {
		// Did we run out of time?
		if time.Now().Unix() >= endTime && config.MaxRunTime != 0 {
			Info.Printf("Max Runtime Reached. Stopping queues...\n")
			break
		}

		if config.ResetXattrs {
			workerResetJobs <- j
		} else {
			workerStartJobs <- j
		}
	}

}
