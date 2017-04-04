package ionice

func PrioToClass(prio uint32) uint32 {
    return prio >> IOPRIO_CLASS_SHIFT
}

func PrioToClassdata(prio uint32) uint32 {
    return prio & IOPRIO_PRIO_MASK
}

func ClassAndClassdataToPrio(class uint32, classdata uint32) uint32 {
    return ((class) << IOPRIO_CLASS_SHIFT) | classdata
}

func PrioToClassAndClassdata(prio uint32) (uint32, uint32) {
    return PrioToClass(prio), PrioToClassdata(prio)
}
