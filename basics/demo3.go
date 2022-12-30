package main

// 全局变量
var x, y int
var (
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

// 全局变量无法使用短声明
func main() {
	// 短声明
	g, h := 123, "hello"
	// 注意短声明和赋值的区别
	println(x, y, a, b, c, d, e, f, g, h)
}
