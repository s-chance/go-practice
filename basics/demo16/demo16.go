package main

import (
	"fmt"
	"strconv"
)

// 类型转换
func main() {
	// int转float32
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Printf("the value of mean is %f\n", mean)

	// go不支持隐式转换

	// string转int
	var str string = "123"
	num, err := strconv.Atoi(str)
	if err == nil {
		fmt.Printf("str: %T %s, num: %T %d\n", str, str, num, num)
	}

	// int转string
	var aNum int = 555
	s := strconv.Itoa(aNum)
	fmt.Printf("aNum: %T %d, s: %T %s\n", aNum, aNum, s, s)

	// string <-> int64
	var target string = "567"
	i, err2 := strconv.ParseInt(target, 10, 64)
	if err2 == nil {
		fmt.Printf("target: %T %d\n", i, i)
	}
	s2 := strconv.FormatInt(i, 10)
	fmt.Printf("target: %T %s\n", s2, s2)
}
