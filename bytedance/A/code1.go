// 题目

// 给定一个十进制整数字符串，判断它是否是 4 的幂。

// 示例 1

// 输入："16"，输出：true

// 示例 2

// 输入："101"，输出：false

// 示例 3

// 输入："70368744177664"，输出：true

// 限定语言：C、 C++、Java、Python、JavaScript V8

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num string
	fmt.Scanf("%s", &num)

	b := powOf4(num)

	fmt.Println(b)
}

func powOf4(s string) bool {
	num, _ := strconv.ParseInt(s, 10, 64)
	n := 0 //记录1的个数
	for ; num != 0; num /= 2 {
		bit := num % 2
		n = n + int(bit)
		if n > 1 {
			return false
		}
	}
	return true
}
