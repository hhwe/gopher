/*
4.8
time
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(start)
	fmt.Printf("%02d.%02d.%04d\n", start.Day(), start.Month(), start.Year())

	t := time.Now().UTC()
	fmt.Println(t)
	p := t.Add(time.Hour)
	fmt.Println(p)
	fmt.Println("t == p", t.Equal(p))
	fmt.Printf("%02d.%02d.%04d\n", t.Day(), t.Month(), t.Year())

	i := time.Now().Unix()
	fmt.Println(i)

	// 格式化的时间不是随意的, 而是特定的
	fmt.Println(t.Format("02 Jan 2006 15:00"))
	fmt.Println(t.Format(time.ANSIC), t.Format(time.RFC822), t.Format("20060101"))

	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	fmt.Print(then)

	end := time.Now()
	elapse := end.Sub(start)
	deadline := time.Time{}
	fmt.Println(elapse)
	fmt.Println(time.Since(start)) // shortof time.Now().Sub(start)
	fmt.Println(time.Until(start)) // short of t.Sub(time.Now())
	fmt.Println(time.Now().Before(deadline), time.Now().After(deadline))

	fmt.Println("\n====== Timer ==========")
	start = time.Now()
	tc := time.After(time.Second)
	<-tc
	fmt.Println(time.Since(start))
	f := func() {
		fmt.Println("time out")
	}
	timer := time.AfterFunc(time.Second, f) // 会起一个goroutine运行f
	timer.Stop()                            // 会stop挂起的goroutine
	time.Sleep(2 * time.Second)

	fmt.Println("\n====== Ticker ==========")
	ticker := time.Tick(time.Second)
	for i := 1; i < 10; i++ {
		<-ticker
		fmt.Println(i)
	}

	// var pt *int = nil
	// *pt = 0 // 错误的使用空指针

	// if true {
	// 	return 1	// 多个return
	// }
	// return 2
}
