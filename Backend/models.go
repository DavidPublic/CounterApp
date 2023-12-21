package main

import (
	"sync"
)

type Counter struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type CountersMap struct {
	sync.RWMutex
	m      map[string]*Counter
	nextID int
}

var counters = NewCountersMap()

func NewCountersMap() *CountersMap {
	return &CountersMap{
		m:      make(map[string]*Counter),
		nextID: 1,
	}
}

func (cm *CountersMap) DeleteCounter(name string) {
    cm.Lock()
    defer cm.Unlock()
    delete(cm.m, name)
}

