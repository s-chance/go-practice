package main

import "fmt"

// 基础语法
func main() {
	fmt.Println("Hello, World!")
	var a = 12
	var b = "2020-12-21"
	var url = "Code=%d&Date=%s"
	var target = fmt.Sprintf(url, a, b)
	fmt.Println(target)
	fmt.Printf(url, a, b)
}
