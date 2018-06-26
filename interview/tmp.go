package main

import (
	"time"
	"fmt"
)

//var p map[string]int

type clockFun func() time.Time
var cc clockFun = time.Now
var dd = clockFun(time.Now)

type notint int
var i notint = 9
var j = notint(2)



func main() {

	fmt.Println(cc())
	fmt.Println(dd())

	fmt.Println(i)
	fmt.Println(j)

	var k = new(notint)
	k = &j
	fmt.Println("k", *k)

	var arr [1e6]int
	fmt.Printf("arr %p\n", &arr)

	//source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	//slice := source[2:4:5]
	//fmt.Println(slice)
	//
	//newslice := append(slice, "Kiwi")
	//fmt.Println(newslice)
	//newslice[0] = "wode"
	//fmt.Println(slice)
	//fmt.Println(newslice)
	//
	//s := append(slice, newslice...)
	//fmt.Println(s, len(s), cap(s))
	//
	//var m map[string]int
	//m = nil
	//print(m, "\n")
	//n := make(map[string]int)
	//print(n, "\n")

	// 创建一个整型切片
	// 其长度和容量都是 4 个元素
	slice := []int{10, 20, 30, 40}
	// 迭代每个元素,并显示值和地址
	for index, value := range slice {
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n",
			value, &value, &slice[index])

		fmt.Println()
	}

	// 创建一个映射,存储颜色以及颜色对应的十六进制代码
	colors := map[string]string{
		"AliceBlue": "#f0f8ff",
		"Coral": "#ff7F50",
		"DarkGray": "#a9a9a9",
		"ForestGreen": "#228b22",
	}
	// 显示映射里的所有颜色
	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s Value-Addr: %X\n", key, value, &value)
	}

}
