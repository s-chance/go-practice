package main

import (
	"fmt"
	"sync"
)

func hello(i int) {
	println("hello goroutinue: " + fmt.Sprint(i))
}

func GoWait() {
	var wg sync.WaitGroup
	wg.Add(5) //先初始化计数器再启动协程
	for i := 0; i < 5; i++ {
		go func(j int) {
			defer wg.Done()
			hello(j)
		}(i)
	}
	wg.Wait()
}
func main() {
	GoWait()
}
