package main

import "fmt"

// 数组
func main() {
	var n [10]int
	var i, j int

	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}

	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}

	var k int
	balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	for i = 0; i < 5; i++ {
		fmt.Printf("balance[%d] = %f\n", i, balance[i])
	}

	// ...自行推断数组长度
	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	for j = 0; j < 5; j++ {
		fmt.Printf("balance2[%d] = %f\n", j, balance2[j])
	}

	// 初始化索引1和3的值
	balance3 := [5]float32{1: 2.0, 3: 7.0}
	for k = 0; k < 5; k++ {
		fmt.Printf("balance3[%d] = %f\n", k, balance3[k])
	}

	// 二维数组
	values := [][]int{}

	rows1 := []int{1, 2, 3}
	rows2 := []int{4, 5, 6}
	values = append(values, rows1)
	values = append(values, rows2)

	fmt.Println("Row 1")
	fmt.Println(values[0])
	fmt.Println("Row 2")
	fmt.Println(values[1])

	fmt.Println("the first element: ")
	fmt.Println(values[0][0])

	// 初始化二维数组
	sites := [2][2]string{}
	sites[0][0] = "Google"
	sites[0][1] = "Runoob"
	sites[1][0] = "Taobao"
	sites[1][1] = "Weibo"
	fmt.Println(sites)

	// 遍历二维数组
	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	// var i, j int

	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}

	// 创建维度元素数量不一致的多维数组
	animals := [][]string{}

	r1 := []string{"fish", "shark", "eel"}
	r2 := []string{"bird"}
	r3 := []string{"lizard", "salamander"}

	animals = append(animals, r1)
	animals = append(animals, r2)
	animals = append(animals, r3)

	for i := range animals {
		fmt.Printf("Row: %v\n", i)
		fmt.Println(animals[i])
	}

	var balance4 = [5]int{1000, 2, 3, 17, 50}
	var avg float32

	avg = getAverage(balance4, 5)

	fmt.Printf("avg = %f\n", avg)

	// 提高精度
	avg *= 1000
	ans := float64(avg) / 1000
	fmt.Printf("ans = %f\n", ans)
}

// 向函数传递数组
func getAverage(arr [5]int, size int) float32 {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum) / float32(size)

	return avg
}
