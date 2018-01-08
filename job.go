package main

import (
	"hash"
	"os"
	"time"

	"github.com/kormoc/bitrot-scanner/pathing"
	"github.com/kormoc/unit/datarate"
)

type job struct {
	checkedTime   int64
	checksumCount int
	checksumMTime int64
	dataRate      datarate.DatarateSIByte
	duration      time.Duration
	hashers       map[ChecksumType]hash.Hash
	mtime         int64
	path          *pathing.Path
}

func newJob(path string, info os.FileInfo) job {
	p, _ := pathing.New(path)
	j := job{
		checkedTime:   GetCheckedTimeXattr(path),
		checksumCount: checksumCount(path),
		mtime:         info.ModTime().Unix(),
		path:          p,
	}

	Trace.Printf("%v: Created job - mtime: %v\n", j.path, j.mtime)

	j.initalizeChecksums()
	return j
}

func (j *job) initalizeChecksums() {
	j.hashers = make(map[ChecksumType]hash.Hash)

	for checksumAlgo := range checksumLookupTable {
		j.hashers[checksumAlgo] = checksumLookupTable[checksumAlgo].New()
	}
}

func (j job) missingChecksums() bool {
	return j.checksumCount != len(checksumLookupTable)
}
