package main

import "io"
import "io/ioutil"
import "log"
import "os"
import "strings"

var Error *log.Logger
var Warn *log.Logger
var Info *log.Logger
var Trace *log.Logger

var logLevels = map[string]int{
	"err":     1,
	"error":   1,
	"warn":    2,
	"warning": 2,
	"info":    3,
	"verbose": 3,
	"debug":   4,
	"trace":   4,
}

func getLogLevelOutput(level string, currentLevel string, output io.Writer) io.Writer {
	if logLevels[strings.ToLower(level)] >= logLevels[strings.ToLower(currentLevel)] {
		return output
	}
	return ioutil.Discard
}

func getLogLevelOutputs(currentLevel string, console io.Writer, logfile io.Writer) io.Writer {
	return io.MultiWriter(
		getLogLevelOutput(consoleLevel, currentLevel, console),
		getLogLevelOutput(logfileLevel, currentLevel, logfile),
	)
}

func setupLogs() {
	// Log to file if defined
	var fp io.Writer
	var err error
	if logfilePath != "" {
		fp, err = os.OpenFile(logfilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fp = ioutil.Discard
	}

	Error = log.New(getLogLevelOutputs("error", os.Stderr, fp), "ERROR: ", 0)
	Warn = log.New(getLogLevelOutputs("warn", os.Stderr, fp), "WARN : ", 0)
	Info = log.New(getLogLevelOutputs("verbose", os.Stdout, fp), "INFO : ", 0)
	Trace = log.New(getLogLevelOutputs("debug", os.Stdout, fp), "DEBUG: ", 0)
}
