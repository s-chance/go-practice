package main

import (
	"fmt"
	"math"
)

// 函数
func main() {
	var a int = 100
	var b int = 200
	var res int

	res = max(a, b)

	fmt.Printf("the max value is %d\n", res)

	c, d := swap("Google", "Runoob")
	fmt.Println(c, d)

	fmt.Printf("before swap1, a = %d, b = %d\n", a, b)
	swap1(a, b)
	fmt.Printf("after swap1, a = %d, b = %d\n", a, b)

	fmt.Printf("before swap2, a = %d, b = %d\n", a, b)
	swap2(&a, &b)
	fmt.Printf("after swap2, a = %d, b = %d\n", a, b)

	// 函数作为实参
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(9))

	nextNumber := getSequence()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea())
}

// 返回一个最大值
func max(num1, num2 int) int {
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

// 返回多个值
func swap(x, y string) (string, string) {
	return y, x
}

// 值传递
func swap1(x, y int) {
	var temp int
	temp = x
	x = y
	y = temp
}

// 引用传递
func swap2(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

// 闭包匿名函数
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// 结构体
type Circle struct {
	radius float64
}

// 方法method，区别于普通函数function，能直接用指针类型的变量调用方法
// 该方法与结构体绑定，一般情况下只能由结构体变量来调用
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}
