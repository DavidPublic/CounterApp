package main

import (
    "sort"
)


func (cm *CountersMap) CreateCounter(name string) {
    cm.Lock()
    defer cm.Unlock()
    cm.m[name] = &Counter{ID: cm.nextID, Name: name, Value: 0}
    cm.nextID++
}

func (cm *CountersMap) IncrementCounter(name string) {
    cm.Lock()
    defer cm.Unlock()
    if counter, ok := cm.m[name]; ok {
        counter.Value++
    }
}

func (cm *CountersMap) GetCounter(name string) (*Counter, bool) {
    cm.RLock()
    defer cm.RUnlock()
    counter, ok := cm.m[name]
    return counter, ok
}

func (cm *CountersMap) GetAllCounters() []*Counter {
    cm.RLock()
    defer cm.RUnlock()
    counters := make([]*Counter, 0, len(cm.m))
    for _, counter := range cm.m {
        counters = append(counters, counter)
    }

    // Sorting the counters slice by ID
    sort.Slice(counters, func(i, j int) bool {
        return counters[i].ID < counters[j].ID
    })

    return counters
}