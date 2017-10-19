package main

import "github.com/kormoc/ionice"
import "os"
import "syscall"

func setNice() {
	if err := syscall.Setpriority(syscall.PRIO_PROCESS, os.Getpid(), config.nice); err != nil {
		Warn.Println("Setting nice failed.")
	}

	if err := ionice.IONiceSelf(uint32(config.ioniceClass), uint32(config.ioniceClassdata)); err != nil {
		Warn.Println("Setting ionice failed.")
	}
}
