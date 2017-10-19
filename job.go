package main

import "github.com/kormoc/unit/datarate"
import "hash"
import "os"
import "time"

type job struct {
	checkedTime   int64
	checksumCount int
	checksumMTime int64
	dataRate      datarate.DatarateSIByte
	duration      time.Duration
	hashers       map[string]hash.Hash
	info          os.FileInfo
	mtime         int64
	path          string
}

func newJob(path string, info os.FileInfo) job {
	j := job{
		checkedTime:   GetCheckedTimeXattr(path),
		checksumCount: checksumCount(path),
		info:          info,
		mtime:         info.ModTime().Unix(),
		path:          path,
	}

	Trace.Printf("%v: Created job - mtime: %v\n", j.path, j.mtime)

	j.initalizeChecksums()
	return j
}

func (j *job) initalizeChecksums() {
	j.hashers = make(map[string]hash.Hash)

	for checksumAlgo := range checksumLookupTable {
		j.hashers[checksumAlgo] = checksumLookupTable[checksumAlgo].New()
	}
}

func (j job) missingChecksums() bool {
	return j.checksumCount != len(checksumLookupTable)
}
