package main

import "log"
import "os"
import "io/ioutil"

var Error *log.Logger
var Warn *log.Logger
var Info *log.Logger
var Trace *log.Logger

func setupLogs() {
    Error = log.New(os.Stderr, "ERROR: ", 0)
    Warn = log.New(os.Stderr, "WARN: ", 0)

    if verbose || debug {
        Info = log.New(os.Stdout, "INFO: ", 0)
    } else {
        Info = log.New(ioutil.Discard, "INFO: ", 0)
    }

    if debug {
        Trace = log.New(os.Stdout, "DEBUG: ", 0)
    } else {
        Trace = log.New(ioutil.Discard, "DEBUG: ", 0)
    }
}
