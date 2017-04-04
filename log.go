package main

import "io"
import "io/ioutil"
import "log"
import "os"

var Error *log.Logger
var Warn *log.Logger
var Info *log.Logger
var Trace *log.Logger

func setupLogs() {
    // Log to file if defined
    var stderr io.Writer
    var stdout io.Writer
    if logfilePath != "" {
        fp, err := os.OpenFile(logfilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
        if err != nil {
            log.Fatal(err)
        }

        stderr = io.MultiWriter(os.Stderr, fp)
        stdout = io.MultiWriter(os.Stdout, fp)
    } else {
        stderr = os.Stderr
        stdout = os.Stdout
    }

    Error = log.New(stderr, "ERROR: ", 0)
    Warn = log.New(stderr, "WARN: ", 0)

    if verbose || debug {
        Info = log.New(stdout, "INFO: ", 0)
    } else {
        Info = log.New(ioutil.Discard, "INFO: ", 0)
    }

    if debug {
        Trace = log.New(stdout, "DEBUG: ", 0)
    } else {
        Trace = log.New(ioutil.Discard, "DEBUG: ", 0)
    }
}
