package main

import "fmt"

// 指针
func main() {
	var a int = 20
	var ip *int
	ip = &a
	fmt.Printf("a的地址: %x\n", &a)
	fmt.Printf("ip存储的地址: %x\n", ip)
	fmt.Printf("ip自身的地址: %x\n", &ip)
	fmt.Printf("*ip的值: %d\n", *ip)

	// 空指针
	var ptr *int

	if ptr == nil {
		fmt.Println("ptr is nil")
	}
	fmt.Printf("ptr的值: %x\n", ptr)

	// 指针数组
	const MAX int = 3
	b := []int{10, 100, 200}
	var i int
	var pb [MAX]*int
	for i = 0; i < MAX; i++ {
		pb[i] = &b[i]
	}
	for i = 0; i < MAX; i++ {
		fmt.Printf("b[%d] = %d\n", i, *pb[i])
	}
}

// 向函数传递指针参数，参考demo8.go的引用传递
