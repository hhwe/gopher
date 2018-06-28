package main

import (
	"sync"
	"fmt"
)

type threadSafeSet struct {
	sync.RWMutex
	s []int
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		set.RLock()

		for elem := range set.s {
			ch <- elem
		}

		close(ch)
		set.RUnlock()

	}()
	return ch
}

func main() {
	t := threadSafeSet{s:[]int{1,2,3}}
	z := t.Iter()
	for i := 0; i < 3; i++ {
		if c, ok := <-z; ok {
			fmt.Println(c)
		}
	}
}
