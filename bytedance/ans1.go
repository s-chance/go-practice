package main

import "fmt"

func main() {
	var num1 string
	fmt.Scanln(&num1)

	var num2 string
	fmt.Scanln(&num2)

	cal(num1, num2)

	fmt.Println(num1)
}

func cal(num1 string, num2 string) {

	for i := 0; i < len(num1); i++ {
		fmt.Printf("%c", num1[i])
	}
	num1 = num2
}
