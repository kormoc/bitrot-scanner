package main

import (
	"sync"
	"time"
)

var workerResetJobs chan job
var workerResetJobswg sync.WaitGroup

func initWorkerReset() {
	workerCountReset := config.WorkerCount
	workerResetJobs = make(chan job, workerCountReset)
	workerResetJobswg.Add(workerCountReset)
	for i := 0; i < workerCountReset; i++ {
		go workerReset()
	}
}

func workerReset() {
	defer workerResetJobswg.Done()
	for currentJob := range workerResetJobs {
		time_start := time.Now()
		Trace.Printf("%v: Reset Processing...\n", currentJob.path)

		err := func() error {
			for _, checksumAlgo := range allChecksumAlgos {
				RemoveChecksumXattr(currentJob.path, checksumAlgo)
			}

			// Also clean up mtimes
			RemoveChecksumXattr(currentJob.path, "mtime")
			// Also clean up checkedtime
			RemoveChecksumXattr(currentJob.path, "checkedtime")

			return nil
		}()

		if err != nil {
			Error.Printf("%v: Reset Processing: %v\n", currentJob.path, err)
		}

		duration := time.Since(time_start)
		currentJob.duration += duration
		Trace.Printf("%v: Reset Processing took %v\n", currentJob.path, duration)
		Info.Printf("%v: Processing took %v\n", currentJob.path, currentJob.duration)
	}
}
