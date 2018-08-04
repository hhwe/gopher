/*
问题 10.1

我们在某个类型的变量上使用点号调用一个方法：variable.method()，在使用 Go 以前，在哪儿碰到过面向对象的点号？

问题 10.2

a）假设定义： type Integer int，完成 get() 方法的方法体: func (p Integer) get() int { ... }。
b）定义： func f(i int) {}; var v Integer ，如何就 v 作为参数调用f？
c）假设 Integer 定义为 type Integer struct {n int}，完成 get() 方法的方法体：func (p Integer) get() int { ... }。
d）对于新定义的 Integer，和 b）中同样的问题。
*/

package main

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
)

func main() {
	integer()
	printString()

	gc()
}

type Integer int

func (p Integer) get() Integer {
	fmt.Println(p)
	return p
}

func f(i int) {
	fmt.Println(i)
}

func integer() {
	var p Integer = 23
	fmt.Println(p)
	f(123)
	// f(p)
}

type TwoInts struct {
	a int
	b int
}

func (tn *TwoInts) String() string {
	return "(" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")"
}

type Celsius float64

func (c Celsius) string() string {
	return strconv.FormatFloat(float64(c), 'E', -1, 64) + "*C"
}

type Day int

var days = []string{
	"sunday",
	"monday",
	"thuesday",
	"wednesday",
	"thursday",
	"friday",
	"saturday",
}

const (
	sunday Day = iota
	monday
	thuesday
	wednesday
	thursday
	friday
	saturday
)

func (d Day) String() string {
	return days[d]
}

type TZ int
type p struct {
	TZ
}

var utc TZ = 1

var m = map[TZ]string{
	1: "Universal Greenwich Time",
}

func (tz TZ) String() string {
	return m[tz]
}

type stack struct {
	l []int
}

func NewStack() *stack {
	return &stack{make([]int, 0, 4)}
}

func (s *stack) String() string {
	var str string
	for i, v := range s.l {
		str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(v) + "]"
	}
	return str
}

func (s *stack) Push(i int) {
	if len(s.l) > 3 {
		log.Fatalln("too long")
	}
	s.l = append(s.l, i)
}

func (s *stack) Pop() int {
	if len(s.l) < 1 {
		log.Fatalln("too short")
	}
	out := s.l[len(s.l)-1]
	p := make([]int, len(s.l)-1)
	copy(p, s.l[:len(s.l)-1])
	s.l = p
	return out
}

func printString() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10
	fmt.Printf("two1 is: %v\n", two1)
	fmt.Println("two1 is:", two1)
	fmt.Printf("two1 is: %T\n", two1)
	fmt.Printf("two1 is: %#v\n", two1)
	c := new(Celsius)
	*c = 123.123
	fmt.Println(c)
	var d Day = 0
	fmt.Println(d)
	var tz TZ = 1
	fmt.Println(tz)

	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s)
	s.Push(5)
	// s.Push(6)
	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s)
}

func gc() {
	// fmt.Printf("%d\n", runtime.MemStats.Alloc/1024)
	// 此处代码在 Go 1.5.1下不再有效，更正为
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc/1024)
}
