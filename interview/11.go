package main

import (
	"fmt"
)

type People1 interface {
	Show()
}

type Student1 struct{}

func (stu *Student1) Show() {

}

func live() People1 {
	var stu *Student1
	return stu
}

func main() {
	var a int
	fmt.Println(a)
	var p []int
	fmt.Println(p)
	var l map[string]int
	fmt.Println(l)

	fmt.Println(live())
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}
