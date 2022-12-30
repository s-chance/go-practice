package main

import (
	"fmt"
)

// 运算符
func main() {
	var a int = 21
	var b int = 10
	var c int

	// 算术运算符
	fmt.Printf("a = %d, b = %d\n", a, b)
	c = a + b
	fmt.Printf("a + b = %d\n", c)
	c = a - b
	fmt.Printf("a - b = %d\n", c)
	c = a * b
	fmt.Printf("a * b = %d\n", c)
	c = a / b
	fmt.Printf("a / b = %d\n", c)
	c = a % b
	fmt.Printf("a %% b = %d\n", c)
	a++
	fmt.Printf("a++ = %d\n", a)
	a = 21 //重新赋值为21
	a--
	fmt.Printf("a-- = %d\n", a)

	// 关系运算符
	if a == b {
		fmt.Printf("a == b\n")
	}
	if a < b {
		fmt.Printf("a < b\n")
	}
	if a > b {
		fmt.Printf("a > b\n")
	}
	a = 5
	b = 20
	fmt.Printf("a = %d, b = %d\n", a, b)
	if a <= b {
		fmt.Printf("a <= b\n")
	}
	if a >= b {
		fmt.Printf("a >= b\n")
	}

	// 逻辑运算符
	var aa bool = true
	var bb bool = false
	fmt.Printf("aa = %t, bb = %t\n", aa, bb)
	if aa && bb {
		fmt.Printf("aa && bb\n")
	}
	if aa || bb {
		fmt.Printf("aa || bb\n")
	}

	aa = false
	bb = true
	fmt.Printf("aa = %t, bb = %t\n", aa, bb)
	if aa && bb {
		fmt.Printf("aa && bb\n")
	}
	if aa || bb {
		fmt.Printf("aa || bb\n")
	}
	if !(aa && bb) {
		fmt.Printf("!(aa && bb)\n")
	}

	// 位运算符
	var i uint = 60 //60 = 0011 1100
	var j uint = 13 //13 = 0000 1101
	var k uint = 0

	k = i & j //12 = 0000 1100
	fmt.Printf("i & j = %d\n", k)

	k = i | j //61 = 0011 1101
	fmt.Printf("i | j = %d\n", k)

	k = i ^ j //49 = 0011 0001
	fmt.Printf("i ^ j = %d\n", k)

	k = i << 2 //240 = 1111 0000
	fmt.Printf("i << 2 = %d\n", k)

	k = i >> 2 //15 = 0000 1111
	fmt.Printf("i >> 2 = %d\n", k)

	// 赋值运算符
	var a1 int = 21
	var c1 int

	c1 = a1
	fmt.Printf("%d\n", c1)

	c1 += a1
	fmt.Printf("%d\n", c1)

	c1 -= a1
	fmt.Printf("%d\n", c1)

	c1 *= a1
	fmt.Printf("%d\n", c1)

	c1 /= a1
	fmt.Printf("%d\n", c1)

	c1 = 200
	fmt.Printf("c1 = %d\n", c1)
	two(200)
	println()
	c1 <<= 2
	fmt.Printf("%d\n", c1)

	c1 >>= 2
	fmt.Printf("%d\n", c1)

	c1 &= 2
	fmt.Printf("%d\n", c1)

	c1 ^= 2
	fmt.Printf("%d\n", c1)

	c1 |= 2
	fmt.Printf("%d\n", c1)

	// 其它运算符
	var v1 int = 4
	var v2 int32
	var v3 float32
	var ptr *int

	fmt.Printf("%T\n", v1)
	fmt.Printf("%T\n", v2)
	fmt.Printf("%T\n", v3)

	ptr = &v1 //ptr指向a变量的地址
	fmt.Printf("v1 = %d\n", v1)
	fmt.Printf("*ptr = %d\n", *ptr)

}

// 二进制转换
func two(num int) {
	if num == 0 {
		return
	}
	b := num % 2
	num /= 2
	two(num)
	fmt.Print(b)
}
