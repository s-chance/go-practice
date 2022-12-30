package main

import (
	"fmt"
	"time"
)

// 条件语句
func main() {
	var a int = 10

	// if
	if a < 20 {
		fmt.Printf("a is smaller than 20\n")
	}
	fmt.Printf("a = %d\n", a)

	var b int = 100

	// if...else
	if b < 20 {
		fmt.Printf("b is smaller than 20\n")
	} else {
		fmt.Printf("b is not smaller than 20\n")
	}
	fmt.Printf("b = %d\n", b)

	var c int = 100
	var d int = 200

	// 嵌套if
	if c == 100 {
		if d == 200 {
			fmt.Printf("a = 100, b = 200\n")
		}
	}
	fmt.Printf("a = %d\n", a)
	fmt.Printf("b = %d\n", b)

	var grade string = "B"
	var marks int = 90

	// switch
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀\n")
	case grade == "B":
		fmt.Printf("良好\n")
	case grade == "C":
		fmt.Printf("及格\n")
	case grade == "D":
		fmt.Printf("下次努力\n")
	default:
		fmt.Printf("加把劲\n")
	}
	fmt.Printf("你的grade是 %s\n", grade)

	var x interface{}

	// type switch
	switch i := x.(type) {
	case nil:
		fmt.Printf("%T\n", i)
	case int:
		fmt.Printf("int\n")
	case float64:
		fmt.Printf("float64\n")
	case func(int) float64:
		fmt.Printf("func(int)\n")
	case bool, string:
		fmt.Printf("bool or string\n")
	default:
		fmt.Printf("unknown\n")
	}

	// fallthrough无视条件强制执行下一条case
	switch {
	case false:
		fmt.Printf("1")
		fallthrough
	case true:
		fmt.Printf("2")
		fallthrough
	case false:
		fmt.Printf("3")
		fallthrough
	case true:
		fmt.Printf("4")
		fallthrough
	case false:
		fmt.Printf("5")
		fallthrough
	default:
		fmt.Printf("default(6)")
	}
	println()

	// select
	// 通道
	c1 := make(chan string)
	c2 := make(chan string)

	// 协程
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("receivd", msg2)
		}
	}

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for {
			ch1 <- "from 1"
		}
	}()
	go func() {
		for {
			ch2 <- "from 2"
		}
	}()

	// 使用select非阻塞地从两个通道中获取数据
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		default:
			fmt.Println("no message received")
		}
	}
}
