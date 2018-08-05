package main

import (
	"container/list"
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
	em()
	person()
	inheritanceCar()

	magic()
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

type employee struct {
	salary float64
}

func (self *employee) giveRaise(p float64) {
	self.salary = self.salary * (1 + p)
}

func em() {
	var p = employee{salary: 1000}
	fmt.Println(p)
	p.giveRaise(0.2)
	fmt.Println(p)
	ll()
}

type li list.List

func (p *li) Iter() {
	fmt.Println(p)
}

func ll() {
	lst := new(li)
	lst.Iter()
}

type Person struct {
	firstName string
	lastName  string
}

func (p *Person) FirstName() string { // getter
	return p.firstName
}

func (p *Person) SetFirstName(newName string) { // setter
	p.firstName = newName
}

func person() {
	p := Person{"韩", "伟"}
	fmt.Println(p.firstName, p.lastName)
}

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	Engine
	wheelCount int
}

func (c *Car) numberOfWheels() {
	fmt.Println(c.wheelCount)
}

type Mercedes struct {
	Car
}

func (m *Mercedes) sayHiToMerkel() {
	m.numberOfWheels()
}

func inheritanceCar() {
	me := Mercedes{Car{wheelCount: 5}}
	me.sayHiToMerkel()
}

type Base struct {
	id int
}

func (b *Base) GetId() int {
	return b.id
}

func (b *Base) SetId(id int) {
	b.id = id
}

type Human struct {
	Base
	firstName string
	lastName  string
}

type Employee struct {
	Human
	salary float64
}

type Base1 struct{}

func (Base1) Magic() {
	fmt.Println("base magic")
}

func (self Base1) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base1
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}

func magic() {
	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()
}
