package notify

import (
    "unsafe"
)

type ObserverFunc func(...interface{})

var hooks map[string]map[uint64]ObserverFunc = make(map[string]map[uint64]ObserverFunc)

func observerKey(observer ObserverFunc) uint64 {
    pt := unsafe.Pointer(&observer)
    return uint64(*(*uintptr)(pt))
}

func AddObserver(notify string, observer ObserverFunc) {
    key := observerKey(observer)

    if _, ok := hooks[notify]; !ok {
        hooks[notify] = make(map[uint64]ObserverFunc)
    }

    hooks[notify][key] = observer
}

func RemoveObserver(notify string, observer ObserverFunc) bool {
    key := observerKey(observer)
    if hk, ok := hooks[notify]; ok {
        if _, ok := hk[key]; ok {
            delete(hk, key)
            return true
        }
    }
    return false
}

func Publish(notify string, args ...interface{}) {
    if observers, ok := hooks[notify]; ok {
        for _, observerFunc := range observers {
            go observerFunc(args...)
        }
    }
}
