package main

import "github.com/nightlyone/lockfile"
import "os"
import "path/filepath"
import "sort"
import "time"
import flag "github.com/ogier/pflag"

func main() {
	time_start := time.Now()

	processFlags()
	setupLogs()
	versionFlag()
	setNice()

	if config.LockfilePath != "" {
		lock, err := lockfile.New(config.LockfilePath)
		if err != nil {
			Error.Fatalf("Lockfile failed. reason: %v\n", err)
		}
		if err := lock.TryLock(); err != nil {
			Error.Fatalf("Lockfile failed. reason: %v\n", err)
		}
		defer lock.Unlock()
	}

	startTime := time.Now().Unix()
	endTime := startTime + config.MaxRunTime

	initWorkers()

	if !config.ResetXattrs {
		filterChecksumAlgos()
	}

	var allJobs []job

	// Loop over the passed in directories and get all files
	for _, path := range flag.Args() {
		Info.Printf("Processing %v...\n", path)
		if err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
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
	sort.Slice(allJobs, func(i, j int) bool { return allJobs[i].checkedTime < allJobs[j].checkedTime })

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

	shutdownWorkers()

	Info.Printf("Ran in %v", time.Since(time_start))
}
