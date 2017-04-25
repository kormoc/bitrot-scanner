package main

import "github.com/nightlyone/lockfile"
import "os"
import "path/filepath"
import "time"
import flag "github.com/ogier/pflag"

func main() {
    time_start := time.Now()

	processFlags()
	setupLogs()
    versionFlag()
    setNice()

	if lockfilePath != "" {
        lock, err := lockfile.New(lockfilePath)
		if err != nil {
			Error.Fatalf("Lockfile failed. reason: %v\n", err)
		}
		if err := lock.TryLock(); err != nil {
			Error.Fatalf("Lockfile failed. reason: %v\n", err)
		}
		defer lock.Unlock()
	}

    initWorkers()

    if !resetXattrs {
        filterChecksumAlgos()
    }

	// Loop over the passed in directories and hash and/or validate

	for _, path := range flag.Args() {
		Info.Printf("Processing %v...\n", path)
		if err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
            if err != nil {
                return err
            }

            j := newJob(path, info)

            if resetXattrs {
                workerResetJobs <- j
            } else {
                workerStartJobs <- j
            }

            return nil

        }); err != nil {
			Error.Println(err)
		}
	}

    shutdownWorkers()

    Info.Printf("Ran in %v", time.Since(time_start))
}
