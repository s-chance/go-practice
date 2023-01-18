package main

import "fmt"

// 接口
func main() {

	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

}

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("This is Nokia")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("This is iPhone")
}
