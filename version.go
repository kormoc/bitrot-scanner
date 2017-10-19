package main

import "fmt"
import "os"

var Version = "__BITROT_SCANNER_VERSION__"

func versionFlag() {
	if config.version {
		fmt.Printf("Version: %v\n", Version)
		os.Exit(0)
	}
}
