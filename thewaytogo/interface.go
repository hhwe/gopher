package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	var s Simpler = &Simple{12}
	simple(s)

	poly()

	methodset2()

	// mainSort()

	reflection()
}

func simple(s Simpler) {
	fmt.Println(s.Get())
	s.Set(26)
	fmt.Println(s.Get())
}

type Simpler interface {
	Get() int
	Set(int)
}

type Simple struct {
	p int
}

func (s *Simple) Get() int {
	return s.p
}

func (s *Simple) Set(i int) {
	s.p = i
}

type Shaper interface {
	Area() float32
}

// type PeriInterface interface {
// 	Perimeter
// }

type Square struct {
	side float32
}

func (sq *Square) Perimeter() float32 {
	return sq.side * sq.side
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

type Circle struct {
	radio float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.radio * c.radio
}

type Triangle struct {
	h, l float32
}

func (t Triangle) Area() float32 {
	return t.h * t.l / 2
}

type Shape struct {
}

func (s Shape) Area() float32 {
	return -1
}

func poly() {
	r := Rectangle{5, 3} // Area() of Rectangle needs a value
	q := &Square{5}      // Area() of Square needs a pointer
	c := Circle{5}
	// shapes := []Shaper{Shaper(r), Shaper(q)}
	// or shorter
	shapes := []Shaper{r, q, c}
	fmt.Println("Looping through shapes for area ...")
	for _, v := range shapes {
		// fmt.Println("Shape details: ", v)
		// if t, ok := v.(Rectangle); ok {
		// 	fmt.Println("Area of this Rectangle is: ", v.Area(), t)
		// }
		// if t, ok := v.(*Square); ok {
		// 	fmt.Println("Area of this Square is: ", v.Area(), t)
		// }
		// if t, ok := v.(Circle); ok {
		// 	fmt.Println("Area of this Circle is: ", v.Area(), t)
		// }
		switch t := v.(type) {
		case Rectangle:
			fmt.Println("Area of this Rectangle is: ", v.Area(), t)
		case *Square:
			fmt.Println("Area of this Square is: ", v.Area(), t)
		case Circle:
			fmt.Println("Area of this Circle is: ", v.Area(), t)
		default:
			fmt.Printf("Unexpected type %T\n", t)
		}
		if _, ok := v.(Shaper); ok {
			fmt.Print("ok")
		}
	}
}

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

func methodset2() {
	// A bare value
	var lst List
	// compiler error:
	// cannot use lst (type List) as type Appender in argument to CountInto:
	//       List does not implement Appender (Append method has pointer receiver)
	// CountInto(lst, 1, 10)
	if LongEnough(lst) { // VALID:Identical receiver type
		fmt.Printf("- lst is long enough\n")
	}

	// A pointer value
	plst := new(List)
	CountInto(plst, 1, 10) //VALID:Identical receiver type
	if LongEnough(plst) {
		// VALID: a *List can be dereferenced for the receiver
		fmt.Printf("- plst is long enough\n")
	}
}

type MyInt int

func reflection() {
	var m MyInt = 5
	v := reflect.ValueOf(m)
	fmt.Println(v, v.Int(), v.Type(), v.Kind(), v.CanSet(), v.Interface())
	var n = 1
	p := reflect.ValueOf(&n)
	fmt.Println(p, p.Type(), p.Kind(), p.CanSet(), p.Interface())
	P := p.Elem()
	P.SetInt(100)
	fmt.Println(P.Interface(), n)
}
