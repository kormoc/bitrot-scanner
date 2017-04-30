package main

import "encoding/hex"
import "sync"
import "time"

var workerEndJobs chan job
var workerEndJobswg sync.WaitGroup

func initWorkerEnd() {
    workerEndJobs = make(chan job, workerCount*2)
    workerEndJobswg.Add(workerCount)
    for i := 0; i < workerCount; i++ {
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

                checksumValue := GetChecksumXattr(currentJob.path, checksumAlgo)

                // If the checksum is missing, just store it
                if len(checksumValue) == 0 {
                    SetMTimeXattr(currentJob.path, currentJob.mtime)
                    SetChecksumXattr(currentJob.path, checksumAlgo, currentHash)
                    continue
                }

                // Do we match? Yay!
                if currentHash == checksumValue {
                    continue
                }

                // No match, but the mtime was updated and the user requested that we update
                // if this happens
                if currentJob.mtime > currentJob.checksumMTime && updateOnNewMTime {
                    Warn.Printf("%v: Updating checksum due to updated mtime\n", currentJob.path)
                    SetMTimeXattr(currentJob.path, currentJob.mtime)
                    SetChecksumXattr(currentJob.path, checksumAlgo, currentHash)
                    continue
                }

                // If this goes backwards, we're kinda confused
                if currentJob.mtime < currentJob.checksumMTime && updateOnNewMTime {
                    Error.Printf("%v: Failed to update checksum due to mtime reversing\n", currentJob.path)
                }

                // Sadness abounds!
                Error.Printf("%v: CHECKSUM MISMATCH!\n\tComputed: %v\n\tExpected: %v\n", currentJob.path, currentHash, checksumValue)
            }

            duration := time.Since(time_start)
            currentJob.duration += duration
            Trace.Printf("%v: End Processing took %v\n", currentJob.path, duration)
            Info.Printf("%v: Processing took %v @ %v.\n", currentJob.path, currentJob.duration, currentJob.dataRate)
            return nil
        }()

        if err != nil {
            Error.Printf("%v: End Processing: %v\n", currentJob.path, err)
        }
    }
}
