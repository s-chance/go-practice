package main

import (
	"fmt"
	"unsafe"
)

const (
	i = "abc"
	j = len(i)
	k = unsafe.Sizeof(i)
)

// 常量
func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" // 多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d", area)
	println()
	println(a, b, c)

	println(i, j, k)

	// 特殊常量iota
	const (
		l = iota //0
		m        //1
		n        //2
		o = "ha" //独立值，iota += 1
		p        //"ha"   iota += 1
		q = 100  //iota += 1
		r        //100    iota += 1
		s = iota //7,恢复计数
		t        //8
	)
	fmt.Println(l, m, n, o, p, q, r, s, t)
}
