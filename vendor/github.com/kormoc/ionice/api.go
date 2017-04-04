package ionice

import "os"

func IONiceSelf(class uint32, classdata uint32) (err error) {
    err = SetIOPriority(IOPRIO_WHO_PROCESS, os.Getpid(), ClassAndClassdataToPrio(class, classdata))
    return
}
