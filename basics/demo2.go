package main

import "fmt"

// 变量
func main() {
	var a string = "Runoob"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d int
	fmt.Println(d)

	var e bool
	fmt.Println(e)

	// 以下默认值均为nil
	// var a *int
	// var a []int
	// var a map[string] int
	// var a chan int
	// var a func(string) int
	// var a error // error 是接口

	// 短声明
	f := "Runoob"
	fmt.Println(f)
}
