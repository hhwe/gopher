// 这个示例程序展示 Go 语言里如何使用接口
package main
import (
	"fmt"
	"reflect"
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

	// 创建一个 users 类型的值,并发送通知
	u := users{"Bill", "bill@email.com"}
	u.notify()
	(&u).notify()
	sendNotification(&u)

	var n notifier
	n = &users{"Jack", "jack@email.com"}
	fmt.Println(n)


	var i int = 64
	t := reflect.TypeOf(i)    //得到类型的元数据,通过t我们能获取类型定义里面的所有元素
	v := reflect.ValueOf(i)   //得到实际的值，通过v我们获取存储在里面的值，还可以去改变值
	tag := t.Elem().Field(0).Tag  //获取定义在struct里面的标签
	name := v.Elem().Field(0).String()  //获取存储在第一个字段里面的值
	fmt.Println(tag, name)


	//var x float64 = 3.4
	//v := reflect.ValueOf(x)
	//
	//fmt.Println("type:", v.Type())
	//fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	//fmt.Println("value:", v.Float())
	//
	//v = reflect.ValueOf(x)
	//v.SetFloat(7.1)
	//fmt.Println("type:", v.Type())
	//fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	//fmt.Println("value:", v.Float())
}

