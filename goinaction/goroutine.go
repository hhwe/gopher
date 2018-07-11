// 这个示例程序展示 goroutine 调度器是如何在单个线程上
// 切分时间片的
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// wg 用来等待程序完成
var wg sync.WaitGroup
var counter int64

// main 是所有 Go 程序的入口
func main() {
	//// 分配一个逻辑处理器给调度器使用
	//runtime.GOMAXPROCS(runtime.NumCPU())
	//// 计数加 2,表示要等待两个 goroutine
	//wg.Add(2)
	//// 创建两个 goroutine
	//fmt.Println("Create Goroutines")
	//go printPrime("A")
	//go printPrime("B")
	//// 等待 goroutine 结束
	//fmt.Println("Waiting To Finish")
	//wg.Wait()
	//fmt.Println("Terminating Program")
	// 计数加 2,表示要等待两个 goroutine
	wg.Add(2)
	// 创建两个 goroutine
	go incCounter(1)
	go incCounter(2)
	// 等待 goroutine 结束
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

// incCounter 增加包里 counter 变量的值
func incCounter(id int) {
	// 在函数退出时调用 Done 来通知 main 函数工作已经完成
	defer wg.Done()
	for count := 0; count < 2; count++ {
		//// 捕获 counter 的值
		//value := counter
		//// 当前 goroutine 从线程退出,并放回到队列
		//runtime.Gosched()
		//// 增加本地 value 变量的值
		//value++
		//// 将该值保存回 counter
		//counter = value

		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}

// printPrime 显示 5000 以内的素数值
func printPrime(prefix string) {
	// 在函数退出时调用 Done 来通知 main 函数工作已经完成
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)

	//go func() {
	//	for {
	//		time.Sleep(1)
	//	}
	//}()
	//
	//for ; ; {
	//	fmt.Println("123")
	//	time.Sleep(100)
	//
	//}

}
