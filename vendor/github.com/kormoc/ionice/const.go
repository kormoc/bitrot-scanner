package ionice

const IOPRIO_CLASS_NONE=uint32(0)
const IOPRIO_CLASS_RT=uint32(1)
const IOPRIO_CLASS_BE=uint32(2)
const IOPRIO_CLASS_IDLE=uint32(3)

const IOPRIO_WHO_PROCESS=1
const IOPRIO_WHO_PGRP=2
const IOPRIO_WHO_USER=3

const IOPRIO_CLASS_SHIFT=uint32(13)
const IOPRIO_PRIO_MASK=((uint32(1) << IOPRIO_CLASS_SHIFT) - 1)

var ClassToString = map[uint32]string{
    IOPRIO_CLASS_NONE: "none",
    IOPRIO_CLASS_RT:   "realtime",
    IOPRIO_CLASS_BE:   "best-effort",
    IOPRIO_CLASS_IDLE: "idle",
}

var StringToClass = map[string]uint32 {
    "none":        IOPRIO_CLASS_NONE,
    "0":           IOPRIO_CLASS_NONE,
    "realtime":    IOPRIO_CLASS_RT,
    "1":           IOPRIO_CLASS_RT,
    "best-effort": IOPRIO_CLASS_BE,
    "2":           IOPRIO_CLASS_BE,
    "idle":        IOPRIO_CLASS_IDLE,
    "3":           IOPRIO_CLASS_IDLE,
}
