package main

import (
	"fmt"
	"math"
	"reflect"
	"time"
)

type Address struct {
	string
}

type Vcard struct {
	Name     string
	Birthday time.Time
	Image    string
	Addr     *Address
}

type Point struct {
	X float64
	Y float64
}

type Rectangle struct {
	x int
	y int
}

func main() {
	vcard()
	point()
	rect()
	tag()
}

func vcard() {
	addr := new(Address)
	addr.string = "12"
	fmt.Println(*addr)
	card := Vcard{"wo", time.Now(), "png", addr}
	fmt.Println(card)
}

func point() {
	p := Point{3, 4}
	l := math.Sqrt(math.Pow(p.X, 2) + math.Pow(p.Y, 2))
	fmt.Println(l)
}

func rect() {
	r := Rectangle{3, 4}
	r.Area()
	r.Perimeter()
}

func (r Rectangle) Area() {
	fmt.Println(r.x * r.y)
}

func (r Rectangle) Perimeter() {
	fmt.Println((r.x + r.y) * 2)
}

type TagType struct { // tags
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

func tag() {
	tt := TagType{true, "Barak Obama", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Println(ixField)
	fmt.Printf("%v %v %v\n", ixField.Tag, ixField.Type, ixField.Index)
}
