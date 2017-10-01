package main

import "github.com/kormoc/ionice"
import "os"
import "syscall"

func setNice() {
	if err := syscall.Setpriority(syscall.PRIO_PROCESS, os.Getpid(), nice); err != nil {
		Warn.Println("Setting nice failed.")
	}

	if err := ionice.IONiceSelf(uint32(ioniceClass), uint32(ioniceClassdata)); err != nil {
		Warn.Println("Setting ionice failed.")
	}
}
