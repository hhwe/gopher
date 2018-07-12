// 这个示例程序展示如何声明
// 并使用方法
package main

import (
	"fmt"
	"time"
)

// user 在程序里定义一个用户类型
type user struct {
	name  string
	email string
}

// notify 使用值接收者实现了一个方法
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

// changeEmail 使用指针接收者实现了一个方法
func (u *user) changeEmail(email string) {
	u.email = email
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func concat(s, t *string) {
	*s = *s + *t
}

func isArr(arr [5]int) {
	fmt.Printf("after: %d  %p \n", arr, &arr)
	arr[0] = 500
	fmt.Printf("change: %d  %p \n", arr, &arr)
}

func isSlice(arr []int) {
	// 传递的时候还是通过复制标头数据的形式, 只不过标头本身带有指向底层数据的指针
	fmt.Printf("after: %d  %p \n", arr, &arr)
	arr[0] = 500
	fmt.Printf("change: %d  %p \n", arr, &arr)
}

type IP []int

func (ip IP) MarshalText() {
	ip[0] = 10
}

// main 是应用程序的入口
func main() {

	// 内置类型: 变量本质上是其存储的数据的内存引用, 都是通过值传递, 可以通过&获取变量地址
	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)

	s, t := "a", "b"
	concat(&s, &t)
	fmt.Println(s, t)

	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("before: %d  %p \n", arr, &arr)
	isArr(arr)
	fmt.Println(arr)

	// 引用类型: 当声明上述类型的变量时,创建的变量被称作标头(header)值。从技术细节上说,字符串也是一种引用类型。
	// 每个引用类型创建的标头值是包含一个指向底层数据结构的指针。每个引用类型还包含一组独特
	// 的字段,用于管理底层数据结构。因为标头值是为复制而设计的,所以永远不需要共享一个引用
	// 类型的值。标头值里包含一个指针,因此通过复制来传递一个引用类型的值的副本,本质上就是
	// 在共享底层数据结构。
	// 在传递的时候也是本质上也是传递值, 只不过不是数据本身的底层结构,
	// 而是标头(底层结构的引用, 具有指向底层结构的指针)
	var sli = []int{1, 2, 3, 4, 5}
	fmt.Printf("before: %d  %p \n", sli, &sli)
	isSlice(sli)
	fmt.Println(sli)

	var ip IP = []int{1, 2, 3, 4, 5}
	ip.MarshalText()
	fmt.Println(ip)

	// user 类型的值可以用来调用
	// 使用值接收者声明的方法
	bill := user{"Bill", "bill@email.com"}
	bill.notify()
	// 指向 user 类型值的指针也可以用来调用
	// 使用值接收者声明的方法
	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()
	// user 类型的值可以用来调用
	// 使用指针接收者声明的方法
	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	// 指向 user 类型值的指针可以用来调用
	// 使用指针接收者声明的方法
	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()

	fmt.Println(time.Time{})
	fmt.Println(time.Now())
}
