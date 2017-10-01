package main

import "github.com/kormoc/unit/datarate"
import "github.com/ncw/directio"
import "io"
import "os"
import "sync"
import "time"

var workerIOJobs chan job
var workerIOJobswg sync.WaitGroup

func initWorkerIO() {
	workerIOJobs = make(chan job, workerCountIO*2)
	workerIOJobswg.Add(workerCountIO)
	for i := 0; i < workerCountIO; i++ {
		go workerIO()
	}
	go func() {
		workerIOJobswg.Wait()
		close(workerEndJobs)

	}()
}

func workerIO() {
	defer workerIOJobswg.Done()
	for currentJob := range workerIOJobs {
		err := func() error {
			time_start := time.Now()
			Trace.Printf("%v: IO Processing...\n", currentJob.path)
			// Try direct io. Fail back to normal IO if needed
			fp, err := directio.OpenFile(currentJob.path, os.O_RDONLY, 0000)
			if err != nil {
				fp, err = os.OpenFile(currentJob.path, os.O_RDONLY, 0000)
				if err != nil {
					return err
				}
			}

			defer fp.Close()

			buffer := directio.AlignedBlock(directio.BlockSize * bufferSize)
			totalRead := 0

			for {
				amountRead, err := io.ReadAtLeast(fp, buffer, 1)
				totalRead += amountRead
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
			currentJob.dataRate = datarate.NewDatarateSIBytes(datarate.Datarate(totalRead)*datarate.Byte, duration)
			Trace.Printf("%v: IO Processing took %v at %v\n", currentJob.path, duration, currentJob.dataRate)
			workerEndJobs <- currentJob
			return nil
		}()

		if err != nil {
			Error.Printf("%v: IO Processing: %v\n", currentJob.path, err)
		}
	}
}
