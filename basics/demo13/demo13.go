package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

// 范围range
func main() {
	// range通常用于数组、切片、通道、集合或字符串等

	for i, v := range pow {
		fmt.Printf("2^%d = %d\n", i, v)
	}

	// 可参考demo7.go range在循环中的用法

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum: ", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// 这里c输出的是unicode值
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
