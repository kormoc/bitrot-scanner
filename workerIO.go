package main

import (
	"io"
	"os"
	"sync"
	"time"

	"github.com/kormoc/unit/datarate"
	"github.com/ncw/directio"
)

var workerIOJobs chan job
var workerIOJobswg sync.WaitGroup

func initWorkerIO() {
	workerIOJobs = make(chan job, config.WorkerCountIO)
	workerIOJobswg.Add(config.WorkerCountIO)
	for i := 0; i < config.WorkerCountIO; i++ {
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
			fp, err := directio.OpenFile(currentJob.path.String(), os.O_RDONLY, 0000)
			if err != nil {
				fp, err = os.OpenFile(currentJob.path.String(), os.O_RDONLY, 0000)
			}
			if err != nil {
				return err
			}

			defer fp.Close()

			writers := make([]io.Writer, 0, len(currentJob.hashers))
			for checksumAlgo := range currentJob.hashers {
				writers = append(writers, currentJob.hashers[checksumAlgo])
			}

			totalRead, err := io.CopyBuffer(
				io.MultiWriter(writers...),
				fp,
				directio.AlignedBlock(directio.BlockSize*config.BufferSize),
			)

			if err != nil {
				return err
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
