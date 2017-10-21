package main

import (
	"io/ioutil"
	"runtime"

	"github.com/kormoc/ionice"
	flag "github.com/ogier/pflag"
	"gopkg.in/yaml.v2"
)

var config = struct {
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
}{
	BufferSize:       2048,
	Checksums:        "sha512",
	ConsoleLevel:     "warn",
	IoniceClass:      int(ionice.IOPRIO_CLASS_IDLE),
	IoniceClassdata:  0,
	LockfilePath:     "",
	LogfileLevel:     "warn",
	LogfilePath:      "",
	MaxRunTime:       0,
	MtimeSettle:      1800,
	Nice:             20,
	ResetXattrs:      false,
	SkipCreate:       false,
	SkipValidation:   false,
	UpdateOnNewMTime: false,
	Version:          false,
	WorkerCount:      0,
	WorkerCountIO:    1,
	XattrRoot:        "user.checksum.",
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

	flag.BoolVar(&config.ResetXattrs, "resetXattrs", config.ResetXattrs, "Don't checksum, just reset any potential checksums")
	flag.BoolVar(&config.SkipCreate, "skipCreate", config.SkipCreate, "Skip creating new hashes. Useful for just validating existing hashes")
	flag.BoolVar(&config.SkipValidation, "skipValidation", config.SkipValidation, "Skip validating existing hashes. Useful to just generate for new files")
	flag.BoolVar(&config.UpdateOnNewMTime, "updateOnNewMTime", config.UpdateOnNewMTime, "Update hashes if mtime is newer then last check time")
	flag.BoolVar(&config.Version, "version", config.Version, "Display the version")
	flag.Int64Var(&config.MaxRunTime, "maxRunTime", config.MaxRunTime, "Stop queueing new jobs after maxRunTime seconds. 0 to disable")
	flag.Int64Var(&config.MtimeSettle, "mtimeSettle", config.MtimeSettle, "Don't create a hash until the mtime is at least this many seconds old")
	flag.IntVar(&config.BufferSize, "bufferSize", config.BufferSize, "Read buffer size in blocks")
	flag.IntVar(&config.IoniceClass, "ioniceClass", config.IoniceClass, "ionice class. 0: none, 1: realtime, 2: best-effort, 3: idle")
	flag.IntVar(&config.IoniceClassdata, "ioniceClassdata", config.IoniceClassdata, "ionice classdata. Only useful for realtime/best-effort")
	flag.IntVar(&config.Nice, "nice", config.Nice, "Nice value to use")
	flag.IntVar(&config.WorkerCount, "workerCount", config.WorkerCount, "Maximum number of workers per stage to use for scanning for all other stages. 0 to detect and use the number of cpu cores")
	flag.IntVar(&config.WorkerCountIO, "workerCountIO", config.WorkerCountIO, "Maximum number of workers to use for reading file data. 0 to detect and use the number of cpu cores")
	flag.StringVar(&config.Checksums, "checksums", config.Checksums, "Which checksum(s) algorithm to use. Comma delimited")
	flag.StringVar(&config.ConsoleLevel, "consoleLevel", config.ConsoleLevel, "Log level for console output. error, warn, verbose, debug")
	flag.StringVar(&config.LockfilePath, "lockfile", config.LockfilePath, "Path to use for a lockfile")
	flag.StringVar(&config.LogfileLevel, "logfileLevel", config.LogfileLevel, "Log level for log file. error, warn, verbose, debug")
	flag.StringVar(&config.LogfilePath, "logfilePath", config.LogfilePath, "Path to logfile")
	flag.StringVar(&config.XattrRoot, "xattrRoot", config.XattrRoot, "base xattr path for checksums")

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
