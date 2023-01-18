package main

import "fmt"

// 结构体
func main() {
	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})

	fmt.Println(Books{title: "Go语言", author: "www.runoob.com"})

	var Book1 Books
	var Book2 Books

	// 访问结构体成员
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	// fmt.Printf("Book1 title: %s\n", Book1.title)
	// fmt.Printf("Book1 author: %s\n", Book1.author)
	// fmt.Printf("Book1 subject: %s\n", Book1.subject)
	// fmt.Printf("Book1 book_id: %d\n", Book1.book_id)

	// fmt.Printf("Book2 title: %s\n", Book2.title)
	// fmt.Printf("Book2 author: %s\n", Book2.author)
	// fmt.Printf("Book2 subject: %s\n", Book2.subject)
	// fmt.Printf("Book2 book_id: %d\n", Book2.book_id)

	// 结构体作为函数参数
	printBook(Book1)
	printBook(Book2)

	// Book1.print()
	// Book2.print()

	// 结构体指针
	var ptr1 *Books = &Book1
	var ptr2 *Books = &Book2
	fmt.Printf("ptr1: %v\n", ptr1)
	fmt.Printf("ptr2: %v\n", ptr2)
}

// 定义结构体
type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

// 函数
func printBook(book Books) {
	fmt.Printf("Book title: %s\n", book.title)
	fmt.Printf("Book author: %s\n", book.author)
	fmt.Printf("Book subject: %s\n", book.subject)
	fmt.Printf("Book book_id: %d\n", book.book_id)
}

// 方法
func (book Books) print() {
	fmt.Printf("Book title: %s\n", book.title)
	fmt.Printf("Book author: %s\n", book.author)
	fmt.Printf("Book subject: %s\n", book.subject)
	fmt.Printf("Book book_id: %d\n", book.book_id)
}
