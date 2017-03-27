package main

import "flag"
import "github.com/kormoc/ionice"

var bufferSize int
var checksums string
var debug bool
var ioniceClass int
var ioniceClassdata int
var lockfilePath string
var mtimeSettle int64
var nice int
var enableProgressBar bool
var resetXattrs bool
var skipCreate bool
var skipValidation bool
var updateOnNewMTime bool
var verbose bool
var workerCount int
var xattrRoot string

func processFlags() {
    flag.BoolVar(   &debug,             "debug",            false,                         "Output debug and verbose log messages")
    flag.BoolVar(   &enableProgressBar, "progressBar",      false,                         "Display a progress bar as checksum")
    flag.BoolVar(   &resetXattrs,       "resetXattrs",      false,                         "Don't checksum, just reset any potential checksums")
    flag.BoolVar(   &skipCreate,        "skipCreate",       false,                         "Skip creating new hashes. Useful for just validating existing hashes")
    flag.BoolVar(   &skipValidation,    "skipValidation",   false,                         "Skip validating existing hashes. Useful to just generate for new files")
    flag.BoolVar(   &updateOnNewMTime,  "updateOnNewMTime", false,                         "Update hashes if mtime is newer then last check time")
    flag.BoolVar(   &verbose,           "verbose",          false,                         "Output verbose log messages")
    flag.Int64Var(  &mtimeSettle,       "mtimeSettle",      1800,                          "Don't create a hash until the mtime is at least this many seconds old")
    flag.IntVar(    &bufferSize,        "bufferSize",       1024*1024*8,                   "Read buffer size")
    flag.IntVar(    &ioniceClass,       "ioniceClass",      int(ionice.IOPRIO_CLASS_IDLE), "ionice class. 0: none, 1: realtime, 2: best-effort, 3: idle")
    flag.IntVar(    &ioniceClassdata,   "ioniceClassdata",  0,                             "ionice classdata. Only useful for realtime/best-effort")
    flag.IntVar(    &nice,              "nice",             20,                            "Nice value to use")
    flag.IntVar(    &workerCount,       "workerCount",      1,                             "Maximum number of workers to use")
    flag.StringVar( &checksums,         "checksums",        "sha512",                      "Which checksum(s) algorithm to use. Comma delimited")
    flag.StringVar( &lockfilePath,      "lockfile",         "",                            "Path to use for a lockfile")
    flag.StringVar( &xattrRoot,         "xattrRoot",        "user.checksum.",              "base xattr path for checksums")

    flag.Parse()
}
