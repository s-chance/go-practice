package main

import (
	"sync"
	"time"
)

func main() {
	Add()
}

var (
	x    int64
	lock sync.Mutex
)

func addWithoutLock() {
	for i := 0; i < 2000; i++ {
		x += 1
	}
}

func addWithLock() {
	for i := 0; i < 2000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
}

func Add() {
	x = 0
	for i := 0; i < 5; i++ {
		go addWithoutLock()
	}
	time.Sleep(time.Second)
	println("without lock:", x)
	x = 0
	for i := 0; i < 5; i++ {
		go addWithLock()
	}
	time.Sleep(time.Second)
	println("with lock:", x)
}
