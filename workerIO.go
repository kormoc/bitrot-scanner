package main

import "io"
import "os"
import "sync"
import "time"

var workerIOJobs chan job
var workerIOJobswg sync.WaitGroup

func initWorkerIO() {
    workerIOJobs = make(chan job, workerCountIO)
    for i := 0; i < workerCountIO; i++ {
        go workerIO()
    }
    go func() {
        workerIOJobswg.Wait()
        close(workerEndJobs)
    }()
}

func workerIO() {
    workerIOJobswg.Add(1)
    defer workerIOJobswg.Done()
    for currentJob := range workerIOJobs {
        err := func() error {
            time_start := time.Now()
            Trace.Printf("%v: IO Processing...\n", currentJob.path)
            fp, err := os.Open(currentJob.path)
            if err != nil {
                return err
            }

            defer fp.Close()

            buffer := make([]byte, bufferSize)

            for {
                amountRead, err := fp.Read(buffer)
                if err == io.EOF {
                    break
                }
                if err != nil {
                    return err
                }

                for checksumAlgo := range currentJob.hashers {
                    currentJob.hashers[checksumAlgo].Write(buffer[:amountRead])
                }
            }

            duration := time.Since(time_start)
            currentJob.duration += duration
            Trace.Printf("%v: IO Processing took %v\n", currentJob.path, duration)
            workerEndJobs <- currentJob
            return nil
        }()

        if err != nil {
            Error.Printf("%v: IO Processing: %v\n", currentJob.path, err)
        }
    }
}
