package notify

import (
    "unsafe"
)

type ObserverFunc func(...interface{})

var hooks map[string]map[uint64]ObserverFunc = make(map[string]map[uint64]ObserverFunc)

func AddObserver(notify string, observer ObserverFunc) {
    pt := unsafe.Pointer(&observer)
    key := uint64(*(*uintptr)(pt))

    if _, ok := hooks[notify]; !ok {
        hooks[notify] = make(map[uint64]ObserverFunc)
    }

    hooks[notify][key] = observer
}

func Publish(notify string, args ...interface{}) {
    if observers, ok := hooks[notify]; ok {
        for _, observerFunc := range observers {
            go observerFunc(args...)
        }
    }
}
