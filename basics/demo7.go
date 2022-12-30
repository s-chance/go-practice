package main

import (
	"fmt"
)

// 循环
func main() {

	// 普通for循环
	sum := 0
	// 写法1
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	// 写法2
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)

	sum = 1
	// 写法3
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)

	// 无限循环
	// for {
	// 	sum++
	// }

	// for-each range循环
	strings := []string{"google", "runoob"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
		fmt.Printf("%d : x = %d\n", i, x)
	}

	// 3种range格式的写法
	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0
	// key和value
	for key, value := range map1 {
		fmt.Printf("%d : %f\n", key, value)
	}
	// 仅key
	for key := range map1 {
		fmt.Printf("key : %d\n", key)
	}
	// 仅value
	for _, value := range map1 {
		fmt.Printf("value : %f\n", value)
	}

	// 嵌套for循环
	// 求2-100内素数
	var i, j int
	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break
			}
		}
		if j > (i / j) {
			fmt.Printf("%d\n", i)
		}
	}

	// 循环控制
	// break
	var a int = 10
	for a < 20 {
		fmt.Printf("a = %d\n", a)
		a++
		if a > 15 {
			break
		}
	}
	// 使用标记能够直接跳出嵌套循环
	// 无标记
	fmt.Println("break without label")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for j := 11; j <= 13; j++ {
			fmt.Printf("j: %d\n", j)
			break // 只跳出内层循环
		}
	}
	// 有标记
	fmt.Println("break with label")
re:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for j := 11; j <= 13; j++ {
			fmt.Printf("j: %d\n", j)
			break re
		}
	}

	// continue
	var b int = 10

	for b < 20 {
		if b == 15 {
			b = b + 1
			continue
		}
		fmt.Printf("b = %d\n", b)
		b++
	}
	// 有无标记的区别
	fmt.Println("continue without label")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for j := 11; j <= 13; j++ {
			fmt.Printf("j: %d\n", j)
			continue
		}
	}
	fmt.Println("continue with label")
ret:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for j := 11; j <= 13; j++ {
			fmt.Printf("j: %d\n", j)
			continue ret
		}
	}

	// goto
	var c int = 10

LOOP:
	for c < 20 {
		if c == 15 {
			c = c + 1
			goto LOOP
		}
		fmt.Printf("c = %d\n", c)
		c++
	}
}
