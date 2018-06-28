package main

import (
	"fmt"
	"runtime"
	"sync"
)

// sync.Mutex 互斥锁，同时只能有一个goroutine获得锁
// sync.RWMutex 读写锁, 可以加多个读锁或者一个写锁
type UserAges struct {
	ages map[string]int
	sync.RWMutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	ua.RLock()
	defer ua.RUnlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

type counter struct {
	sync.RWMutex
	m map[string]int
}

func (c *counter) set() {
	c.Lock()
	c.m["s"]++
	c.Unlock()
}

func (c *counter) get() int {
	c.RLock()
	n := c.m["s"]
	c.RUnlock()

	return n
}

func main() {
	var counter = counter{m: make(map[string]int)}

	runtime.GOMAXPROCS(4)
	wg := new(sync.WaitGroup)
	wg.Add(100000)

	ua := UserAges{ages: make(map[string]int)}

	//m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6}
	for i := 0; i < 100000; i++ {
		go func(i int) {
			ua.Add("key", 2)
			ua.Get("key")
			//m["a"] = i
			wg.Done()
		}(i)

	}

	fmt.Println(counter)

	//p := make(map[string]int)
	//p["key"] = 23
	//
	//fmt.Println(p)

	wg.Wait()
	fmt.Println(ua.ages)
}

//package main
//
//import (
//	"fmt"
//	"sync"
//	"runtime"
//)
//
//func main() {
//
//	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6}
//
//	fmt.Println(m)
//	fmt.Println("-------")
//
//	runtime.GOMAXPROCS(4)
//	wg := &sync.WaitGroup{}
//	wg.Add(4)
//
//	// 因为go语言中的map为引用类型， 虽然go语言函数以传值方式调用，
//	// 即函数持有的为参数副本，但因是引用类型， 故依然指向映射m ， 类似c的指针.
//	for i := 0; i < 4; i++ {
//		go func(p map[string]int, i int) {
//
//			defer wg.Done()
//			p["a"] = i
//
//		}(m, i)
//	}
//
//
//
//	wg.Wait()
//
//	fmt.Println(m)
//}
