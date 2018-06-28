//package main
//
//import (
//	"sync"
//	"fmt"
//)
//
//type threadSafeSet struct {
//	sync.RWMutex
//	s []int
//}
//
//func (set *threadSafeSet) Iter() <-chan interface{} {
//	ch := make(chan interface{})
//	go func() {
//		set.RLock()
//
//		for elem := range set.s {
//			ch <- elem
//		}
//
//		close(ch)
//		set.RUnlock()
//
//	}()
//	return ch
//}
//
//func main() {
//	t := threadSafeSet{s:[]int{1,2,3}}
//	z := t.Iter()
//	for i := 0; i < 100; i++ {
//		fmt.Println(<-z)
//	}
//
//
//}



//author: ysqi ,https://yushuangqi.com

package main

import "fmt"
import "sync"
import "time"

type ThreadSafeSet struct {
	sync.RWMutex
	s []int
}

func (set *ThreadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		set.RLock()

		for elem := range set.s {
			ch <- elem
			fmt.Print("get:", elem, ",")
		}

		close(ch)
		set.RUnlock()

	}()
	return ch
}

func main() {
	read()
	unRead()
}
func read() {
	set := ThreadSafeSet{}
	set.s = make([]int, 100)
	ch := set.Iter()
	closed := false
	for {
		select {
		case v, ok := <-ch:
			if ok {
				fmt.Print("read:", v, ",")
			} else {
				closed = true
			}
		}
		if closed {
			fmt.Print("closed")
			break
		}
	}
	fmt.Print("Done")
}

func unRead() {
	set := ThreadSafeSet{}
	set.s = make([]int, 100)
	ch := set.Iter()
	_ = ch
	time.Sleep(5 * time.Second)
	fmt.Print("Done")
}
