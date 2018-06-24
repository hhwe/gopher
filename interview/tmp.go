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


	//var m map[string]int
	//m = nil
	//print(m, "\n")
	//n := make(map[string]int)
	//print(n, "\n")

}
