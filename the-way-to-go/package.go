package main

import (
	"container/list"
	"fmt"
)

func main() {
	doubleLinkList()
}

func doubleLinkList() {
	l := list.New()
	fmt.Println(l)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
