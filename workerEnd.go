package main

import (
	"encoding/hex"
	"sync"
	"time"
)

var workerEndJobs chan job
var workerEndJobswg sync.WaitGroup

func initWorkerEnd() {
	workerCountEnd := config.WorkerCount
	workerEndJobs = make(chan job, workerCountEnd)
	workerEndJobswg.Add(workerCountEnd)
	for i := 0; i < workerCountEnd; i++ {
		go workerEnd()
	}
}

func workerEnd() {
	defer workerEndJobswg.Done()
	for currentJob := range workerEndJobs {
		err := func() error {
			time_start := time.Now()
			Trace.Printf("%v: End Processing...\n", currentJob.path)
			for checksumAlgo := range currentJob.hashers {
				currentHash := hex.EncodeToString(currentJob.hashers[checksumAlgo].Sum(nil))

				checksumValue := GetChecksumXattr(currentJob.path.String(), checksumAlgo)

				// If the checksum is missing, just store it
				if len(checksumValue) == 0 {
					SetMTimeXattr(currentJob.path.String(), currentJob.mtime)
					SetChecksumXattr(currentJob.path.String(), checksumAlgo, currentHash)
					continue
				}

				// Do we match? Yay!
				if currentHash == checksumValue {
					continue
				}

				// No match, but the mtime was updated and the user requested that we update
				// if this happens
				if currentJob.mtime > currentJob.checksumMTime && config.UpdateOnNewMTime {
					Warn.Printf("%v: Updating checksum due to updated mtime\n", currentJob.path)
					SetMTimeXattr(currentJob.path.String(), currentJob.mtime)
					SetChecksumXattr(currentJob.path.String(), checksumAlgo, currentHash)
					continue
				}

				// If this goes backwards, we're kinda confused
				if currentJob.mtime < currentJob.checksumMTime && config.UpdateOnNewMTime {
					Error.Printf("%v: Failed to update checksum due to mtime reversing\n", currentJob.path)
				}

				// Sadness abounds!
				Error.Printf("%v: CHECKSUM MISMATCH!\n\tComputed: %v\n\tExpected: %v\n", currentJob.path, currentHash, checksumValue)
			}

			SetCheckedTimeXattr(currentJob.path.String(), time.Now().Unix())

			duration := time.Since(time_start)
			currentJob.duration += duration
			Trace.Printf("%v: Finished in %v @ %v.\n", currentJob.path, currentJob.duration, currentJob.dataRate)
			return nil
		}()

		if err != nil {
			Error.Printf("%v: End Processing: %v\n", currentJob.path, err)
		}
	}
}
