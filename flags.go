package main

import "github.com/kormoc/ionice"
import "runtime"
import flag "github.com/ogier/pflag"

var bufferSize int
var checksums string
var consoleLevel string
var ioniceClass int
var ioniceClassdata int
var lockfilePath string
var logfileLevel string
var logfilePath string
var mtimeSettle int64
var nice int
var resetXattrs bool
var skipCreate bool
var skipValidation bool
var updateOnNewMTime bool
var version bool
var workerCount int
var workerCountIO int
var xattrRoot string

func processFlags() {
	flag.BoolVar(&resetXattrs, "resetXattrs", false, "Don't checksum, just reset any potential checksums")
	flag.BoolVar(&skipCreate, "skipCreate", false, "Skip creating new hashes. Useful for just validating existing hashes")
	flag.BoolVar(&skipValidation, "skipValidation", false, "Skip validating existing hashes. Useful to just generate for new files")
	flag.BoolVar(&updateOnNewMTime, "updateOnNewMTime", false, "Update hashes if mtime is newer then last check time")
	flag.BoolVar(&version, "version", false, "Display the version")
	flag.Int64Var(&mtimeSettle, "mtimeSettle", 1800, "Don't create a hash until the mtime is at least this many seconds old")
	flag.IntVar(&bufferSize, "bufferSize", 2048, "Read buffer size in blocks")
	flag.IntVar(&ioniceClass, "ioniceClass", int(ionice.IOPRIO_CLASS_IDLE), "ionice class. 0: none, 1: realtime, 2: best-effort, 3: idle")
	flag.IntVar(&ioniceClassdata, "ioniceClassdata", 0, "ionice classdata. Only useful for realtime/best-effort")
	flag.IntVar(&nice, "nice", 20, "Nice value to use")
	flag.IntVar(&workerCount, "workerCount", 0, "Maximum number of workers per stage to use for scanning for all other stages. 0 to detect and use the number of cpu cores")
	flag.IntVar(&workerCountIO, "workerCountIO", 1, "Maximum number of workers to use for reading file data. 0 to detect and use the number of cpu cores")
	flag.StringVar(&checksums, "checksums", "sha512", "Which checksum(s) algorithm to use. Comma delimited")
	flag.StringVar(&consoleLevel, "consoleLevel", "warn", "Log level for console output. error, warn, verbose, debug")
	flag.StringVar(&lockfilePath, "lockfile", "", "Path to use for a lockfile")
	flag.StringVar(&logfileLevel, "logfileLevel", "warn", "Log level for log file. error, warn, verbose, debug")
	flag.StringVar(&logfilePath, "logfilePath", "", "Path to logfile")
	flag.StringVar(&xattrRoot, "xattrRoot", "user.checksum.", "base xattr path for checksums")

	flag.Parse()

	if workerCount == 0 {
		workerCount = runtime.NumCPU()
	}
	if workerCount < 1 {
		workerCount = 1
	}

	if workerCountIO == 0 {
		workerCountIO = runtime.NumCPU()
	}
	if workerCountIO < 1 {
		workerCountIO = 1
	}
}
