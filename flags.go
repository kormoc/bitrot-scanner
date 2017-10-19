package main

import "github.com/kormoc/ionice"
import "runtime"
import flag "github.com/ogier/pflag"

var config struct {
	bufferSize       int
	checksums        string
	consoleLevel     string
	ioniceClass      int
	ioniceClassdata  int
	lockfilePath     string
	logfileLevel     string
	logfilePath      string
	maxRunTime       int64
	mtimeSettle      int64
	nice             int
	resetXattrs      bool
	skipCreate       bool
	skipValidation   bool
	updateOnNewMTime bool
	version          bool
	workerCount      int
	workerCountIO    int
	xattrRoot        string
}

func processFlags() {
	flag.BoolVar(&config.resetXattrs, "resetXattrs", false, "Don't checksum, just reset any potential checksums")
	flag.BoolVar(&config.skipCreate, "skipCreate", false, "Skip creating new hashes. Useful for just validating existing hashes")
	flag.BoolVar(&config.skipValidation, "skipValidation", false, "Skip validating existing hashes. Useful to just generate for new files")
	flag.BoolVar(&config.updateOnNewMTime, "updateOnNewMTime", false, "Update hashes if mtime is newer then last check time")
	flag.BoolVar(&config.version, "version", false, "Display the version")
	flag.Int64Var(&config.maxRunTime, "maxRunTime", 0, "Stop queueing new jobs after maxRunTime seconds. 0 to disable")
	flag.Int64Var(&config.mtimeSettle, "mtimeSettle", 1800, "Don't create a hash until the mtime is at least this many seconds old")
	flag.IntVar(&config.bufferSize, "bufferSize", 2048, "Read buffer size in blocks")
	flag.IntVar(&config.ioniceClass, "ioniceClass", int(ionice.IOPRIO_CLASS_IDLE), "ionice class. 0: none, 1: realtime, 2: best-effort, 3: idle")
	flag.IntVar(&config.ioniceClassdata, "ioniceClassdata", 0, "ionice classdata. Only useful for realtime/best-effort")
	flag.IntVar(&config.nice, "nice", 20, "Nice value to use")
	flag.IntVar(&config.workerCount, "workerCount", 0, "Maximum number of workers per stage to use for scanning for all other stages. 0 to detect and use the number of cpu cores")
	flag.IntVar(&config.workerCountIO, "workerCountIO", 1, "Maximum number of workers to use for reading file data. 0 to detect and use the number of cpu cores")
	flag.StringVar(&config.checksums, "checksums", "sha512", "Which checksum(s) algorithm to use. Comma delimited")
	flag.StringVar(&config.consoleLevel, "consoleLevel", "warn", "Log level for console output. error, warn, verbose, debug")
	flag.StringVar(&config.lockfilePath, "lockfile", "", "Path to use for a lockfile")
	flag.StringVar(&config.logfileLevel, "logfileLevel", "warn", "Log level for log file. error, warn, verbose, debug")
	flag.StringVar(&config.logfilePath, "logfilePath", "", "Path to logfile")
	flag.StringVar(&config.xattrRoot, "xattrRoot", "user.checksum.", "base xattr path for checksums")

	flag.Parse()

	if config.workerCount == 0 {
		config.workerCount = runtime.NumCPU()
	}
	if config.workerCount < 1 {
		config.workerCount = 1
	}

	if config.workerCountIO == 0 {
		config.workerCountIO = runtime.NumCPU()
	}
	if config.workerCountIO < 1 {
		config.workerCountIO = 1
	}
}
