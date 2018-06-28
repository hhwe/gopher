package main

import (
	"sync"
	"runtime"
	"fmt"
)

type UserAges struct {
	ages map[string]int
	sync.RWMutex
}

func (ua *UserAges) Add(name string, age int)  {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}

	return -1
}

func main() {
	runtime.GOMAXPROCS(4)
	fmt.Println(runtime.NumCPU())
	wg := new(sync.WaitGroup)
	wg.Add(4)

	ua := new(UserAges)

	for i := 0; i<4; i++ {
		go func() {
			fmt.Println(i)
			ua.Add("key", 2)
			fmt.Println(ua.Get("key"))
			wg.Done()
		}()

	}

	//p := make(map[string]int)
	//p["key"] = 23
	//
	//fmt.Println(p)

	wg.Wait()
	fmt.Println(ua.ages)
}
