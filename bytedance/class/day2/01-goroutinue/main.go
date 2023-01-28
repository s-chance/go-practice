package main

import (
	"fmt"
	"time"
)

func main() {
	HelloGoRoutinue()
}

func hello(i int) {
	println("hello goroutinue: " + fmt.Sprint(i))
}

func HelloGoRoutinue() {
	for i := 0; i < 5; i++ {
		go func(j int) {
			hello(j)
		}(i)
	}
	time.Sleep(time.Second)
}
