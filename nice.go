package main

import (
	"os"
	"syscall"

	"github.com/kormoc/ionice"
)

func setNice() {
	if err := syscall.Setpriority(syscall.PRIO_PROCESS, os.Getpid(), config.Nice); err != nil {
		Warn.Println("Setting nice failed.")
	}

	if err := ionice.IONiceSelf(uint32(config.IoniceClass), uint32(config.IoniceClassdata)); err != nil {
		Warn.Println("Setting ionice failed.")
	}
}
