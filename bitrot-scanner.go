package main

import (
	"time"

	"github.com/nightlyone/lockfile"
)

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

	if config.SortedJobs {
		sortedJobs(startTime, endTime)
	} else {
		unsortedJobs(startTime, endTime)
	}

	shutdownWorkers()

	Info.Printf("Ran in %v", time.Since(time_start))
}
