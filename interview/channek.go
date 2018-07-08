package main

import "fmt"

func main() {
	n := make(chan int)
	n <-1  // 将会阻塞当代前goroutine， 直到其他goroutine去读取
	fmt.Println(<-n)  // 上个通道操作以及阻塞，这句不会执行


	// 必须设置为有缓冲通道，否则写入通道会阻塞当前goroutine
	c := make(chan int, 1)
	c <-1
	fmt.Println(<-c)

	//go func() {
	//	c <- 1
	//	//fmt.Println(t)
	//}()
	//go func() {
	//	fmt.Println(<-c)
	//}()
	//
	//t := <- c



}
