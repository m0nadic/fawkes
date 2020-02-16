package main

import "sync"

var mutex = &sync.Mutex{}
var isServerHealthy = true

func isHealthy() bool {
	return isServerHealthy
}

func makeHealthy() {
	mutex.Lock()
	defer mutex.Unlock()
	isServerHealthy = true
}

func makeSick() {
	mutex.Lock()
	defer mutex.Unlock()
	isServerHealthy = false
}
