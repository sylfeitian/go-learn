package main

import "fmt"

var x, y int
var (
	a int
	b bool
)

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var socketcode = 123
	var enddate = "2020-12-31"
	fmt.Println(enddate)
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, socketcode, enddate)
	fmt.Println(target_url)

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	intVal := 10
	fmt.Println(intVal)

	a1, a2, a3 := 1, "12", true
	fmt.Println(a1, a2, a3)
	var ret int

	ret = max(12, 14)

	fmt.Printf("最大值：%d\n", ret)

	a4, a5 := swap("Google", "Runoob")
	fmt.Println(a4, a5)

	var a6 int = 100
	var a7 int = 200

	fmt.Printf("交换前，a6 的值：%d\n", a6)
	fmt.Printf("交换后，a7 的值：%d\n", a7)
	swap2(&a6, &a7)

	fmt.Printf("交换前，a6 的值：%d\n", a6)
	fmt.Printf("交换后，a7 的值：%d\n", a7)

	var Book1 Books /* 声明 Book1 为 Books 类型 */
	var Book2 Books /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	printBook(&Book1)

	/* 打印 Book2 信息 */
	printBook(&Book2)

	day02()

	day03()
}

func printBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func max(num1, num2 int) int {
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}

	return result
}

func swap(x, y string) (string, string) {
	return y, x
}

func swap2(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
