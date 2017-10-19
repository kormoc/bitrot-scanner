package main

import "github.com/kormoc/ionice"
import "gopkg.in/yaml.v2"
import "io/ioutil"
import "runtime"
import flag "github.com/ogier/pflag"

var config struct {
	BufferSize       int    `yaml:"bufferSize"`
	Checksums        string `yaml:"checksums"`
	ConsoleLevel     string `yaml:"consoleLevel"`
	IoniceClass      int    `yaml:"ioniceClass"`
	IoniceClassdata  int    `yaml:"ioniceClassdata"`
	LockfilePath     string `yaml:"lockfile"`
	LogfileLevel     string `yaml:"logfileLevel"`
	LogfilePath      string `yaml:"logfilePath"`
	MaxRunTime       int64  `yaml:"maxRunTime"`
	MtimeSettle      int64  `yaml:"mtimeSettle"`
	Nice             int    `yaml:"nice"`
	ResetXattrs      bool   `yaml:"resetXattrs"`
	SkipCreate       bool   `yaml:"skipCreate"`
	SkipValidation   bool   `yaml:"skipValidation"`
	UpdateOnNewMTime bool   `yaml:"updateOnNewMTime"`
	Version          bool   `yaml:"version"`
	WorkerCount      int    `yaml:"workerCount"`
	WorkerCountIO    int    `yaml:"workerCountIO"`
	XattrRoot        string `yaml:"xattrRoot"`
}

func processFlags() (err error) {
	var configData []byte
	configData, err = ioutil.ReadFile("/etc/bitrot-scanner.yaml")
	if err == nil {
		err = yaml.Unmarshal(configData, &config)
		if err != nil {
			return
		}
	}

	flag.BoolVar(&config.ResetXattrs, "resetXattrs", false, "Don't checksum, just reset any potential checksums")
	flag.BoolVar(&config.SkipCreate, "skipCreate", false, "Skip creating new hashes. Useful for just validating existing hashes")
	flag.BoolVar(&config.SkipValidation, "skipValidation", false, "Skip validating existing hashes. Useful to just generate for new files")
	flag.BoolVar(&config.UpdateOnNewMTime, "updateOnNewMTime", false, "Update hashes if mtime is newer then last check time")
	flag.BoolVar(&config.Version, "version", false, "Display the version")
	flag.Int64Var(&config.MaxRunTime, "maxRunTime", 0, "Stop queueing new jobs after maxRunTime seconds. 0 to disable")
	flag.Int64Var(&config.MtimeSettle, "mtimeSettle", 1800, "Don't create a hash until the mtime is at least this many seconds old")
	flag.IntVar(&config.BufferSize, "bufferSize", 2048, "Read buffer size in blocks")
	flag.IntVar(&config.IoniceClass, "ioniceClass", int(ionice.IOPRIO_CLASS_IDLE), "ionice class. 0: none, 1: realtime, 2: best-effort, 3: idle")
	flag.IntVar(&config.IoniceClassdata, "ioniceClassdata", 0, "ionice classdata. Only useful for realtime/best-effort")
	flag.IntVar(&config.Nice, "nice", 20, "Nice value to use")
	flag.IntVar(&config.WorkerCount, "workerCount", 0, "Maximum number of workers per stage to use for scanning for all other stages. 0 to detect and use the number of cpu cores")
	flag.IntVar(&config.WorkerCountIO, "workerCountIO", 1, "Maximum number of workers to use for reading file data. 0 to detect and use the number of cpu cores")
	flag.StringVar(&config.Checksums, "checksums", "sha512", "Which checksum(s) algorithm to use. Comma delimited")
	flag.StringVar(&config.ConsoleLevel, "consoleLevel", "warn", "Log level for console output. error, warn, verbose, debug")
	flag.StringVar(&config.LockfilePath, "lockfile", "", "Path to use for a lockfile")
	flag.StringVar(&config.LogfileLevel, "logfileLevel", "warn", "Log level for log file. error, warn, verbose, debug")
	flag.StringVar(&config.LogfilePath, "logfilePath", "", "Path to logfile")
	flag.StringVar(&config.XattrRoot, "xattrRoot", "user.checksum.", "base xattr path for checksums")

	flag.Parse()

	if config.WorkerCount == 0 {
		config.WorkerCount = runtime.NumCPU()
	}
	if config.WorkerCount < 1 {
		config.WorkerCount = 1
	}

	if config.WorkerCountIO == 0 {
		config.WorkerCountIO = runtime.NumCPU()
	}
	if config.WorkerCountIO < 1 {
		config.WorkerCountIO = 1
	}

	return
}
