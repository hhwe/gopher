// 这个示例程序展示 Go 语言里如何使用接口
package main
import (
	"fmt"
)

// notifier 是一个定义了
// 通知类行为的接口
type notifier interface {
	notify()
}

// users 在程序里定义一个用户类型
type users struct {
	name string
	email string
}

// sendNotification 接受一个实现了 notifier 接口的值
// 并发送通知
func sendNotification(n notifier) {
	n.notify()
}

// notify 是使用指针接收者实现的方法
func (u *users) notify() {
	fmt.Printf("Sending users email to %s<%s>\n",
		u.name,
		u.email)
}

// duration 是一个基于 int 类型的类型
type duration1 int
// 使用更可读的方式格式化 duration 值
func (d *duration1) pretty() string {
	return fmt.Sprintf("Duration: %d", *d)
}


func printAll(vals interface{}) { //1
	fmt.Println(vals)
}

func main(){
	names := []string{"stanley", "david", "oscar"}
	printAll(names)

	var d = duration1(42)
	fmt.Println(d.pretty())
	// ./listing46.go:17: 不能通过指针调用 duration(42)的方法
	// ./listing46.go:17: 不能获取 duration(42)的地址

	// 创建一个 users 类型的值,并发送通知
	u := users{"Bill", "bill@email.com"}
	u.notify()
	(&u).notify()
	sendNotification(&u)

	var n notifier
	n = &users{"Jack", "jack@email.com"}
	fmt.Println(n)
	// ./listing36.go:32: 不能将 u(类型是 users)作为
	// sendNotification 的参数类型 notifier:
	// users 类型并没有实现 notifier
	// (notify 方法使用指针接收者声明)
}

