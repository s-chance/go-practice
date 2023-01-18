package main

import "fmt"

// 切片slice
func main() {
	// make(type, len, cap)
	var numbers = make([]int, 3, 5)
	printSlice(numbers)

	// 空(nil)切片
	var num []int
	printSlice(num)

	if num == nil {
		fmt.Printf("空切片\n")
	}

	// 切片截取
	n := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(n)

	fmt.Println("n == ", n)
	// 切片索引1(包含)到4(不包含)
	fmt.Println("n[1:4] == ", n[1:4]) //1 2 3
	// 默认下限为0
	fmt.Println("n[:3] == ", n[:3]) //0 1 2
	// 默认上限为len(s)
	fmt.Println("n[4:] == ", n[4:]) //4 5 6 7 8

	n1 := make([]int, 0, 5)
	printSlice(n1)

	n2 := n[:2]
	// len = 2 - 0 = 2
	// cap = 9 - 0 = 9
	printSlice(n2)

	n3 := n[2:5]
	// len = 5 - 2 = 3
	// cap = 9 - 2 = 7
	printSlice(n3)

	// append追加元素
	var nn []int
	printSlice(nn)

	nn = append(nn, 0)
	printSlice(nn)

	nn = append(nn, 1)
	printSlice(nn)
	// 在元素数量超过2时, cap会自动扩容2个容量
	// 之后每超过最大容量都会进行扩容2个容量
	// len % 2 + len = cap
	nn = append(nn, 2, 3, 4)
	printSlice(nn)

	// 2倍扩容
	// 创建新的切片
	nnPro := make([]int, len(nn), (cap(nn))*2)
	// 拷贝旧切片的元素到新切片
	// copy(new, old)
	copy(nnPro, nn)
	printSlice(nnPro)
}

func printSlice(x []int) {
	// len当前元素数量, cap当前最大容量
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(x), cap(x), x)
}
