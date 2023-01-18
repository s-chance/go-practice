package main

import (
	"fmt"
	"time"
)

// 并发
func main() {
	// 使用go开启两个goroutine
	// 一般情况下默认只有一个goroutine
	go say("world")
	say("hello")

	// 通道channel
	// 默认无缓冲
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c //从通道c中接收数据

	fmt.Println(x, y, x+y)

	// 通道缓冲区
	// 创建带缓冲的通道
	ch := make(chan int, 2)

	// 缓冲允许异步获取数据，即不需要立刻接收数据
	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// 遍历通道与关闭通道
	chnnal := make(chan int, 10)
	go fibonacci(cap(chnnal), chnnal)

	for i := range chnnal {
		fmt.Println(i)
	}
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // sum的数据发送给通道c
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // 关闭通道否则会进入阻塞
}
