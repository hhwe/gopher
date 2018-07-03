// 这个示例程序展示如何声明
// 并使用方法
package main
import (
	"fmt"
	"time"
)

// user 在程序里定义一个用户类型
type user struct {
	name string
	email string
}

type IP []byte

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
	fmt.Printf("after: %d  %p \n",arr, &arr)
	arr[0] = 500
	fmt.Printf("change: %d  %p \n",arr, &arr)
}

func isSlice(arr []int) {
	fmt.Printf("after: %d  %p \n",arr, &arr)
	arr[0] = 500
	fmt.Printf("change: %d  %p \n",arr, &arr)
}

// main 是应用程序的入口
func main() {

	// 内置类型: 变量本质上是内存引用, 可以通过&获取变量地质
	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)

	s, t := "a", "b"
	concat(&s, &t)
	fmt.Println(s, t)

	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("before: %d  %p \n",arr, &arr)
	isArr(arr)
	fmt.Println(arr)

	//
	var sli = []int{1, 2, 3, 4, 5}
	fmt.Printf("before: %d  %p \n",sli, &sli)
	isSlice(sli)
	fmt.Println(sli)

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