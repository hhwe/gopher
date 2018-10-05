package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a interface{}
	a = 6
	fmt.Println(reflect.TypeOf(a), reflect.ValueOf(a))

	// todo: 1.Reflection goes from interface value to reflection object.
	// At the basic level, reflection is just a mechanism to examine the type
	// and value pair stored inside an interface variable. To get started,
	// there are two types we need to know about in package reflect:
	// Type and Value. Those two types give access to the contents of an
	// interface variable, and two simple functions, called reflect.TypeOf
	// and reflect.ValueOf, retrieve reflect.Type and reflect.Value pieces
	// out of an interface value. (Also, from the reflect.Value it's easy to
	// get to the reflect.Type, but let's keep the Value and Type concepts
	// separate for now.)
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

	var y uint8 = 'x'
	v = reflect.ValueOf(y)
	fmt.Println("type:", v.Type())                            // uint8.
	fmt.Println("kind is uint8: ", v.Kind() == reflect.Uint8) // true.
	//y = v.Uint() 反射返回的是对应类型的极限类型，如：int返回的是int64
	y = uint8(v.Uint()) // v.Uint returns a uint64.

	type MyInt int
	var z MyInt = 7
	v = reflect.ValueOf(z)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is Int:", v.Kind() == reflect.Int)
	fmt.Println("value:", v.Int())

	type man struct {
		name string
		age  int
	}
	m := new(man)
	v = reflect.ValueOf(m)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is Struct:", v.Kind() == reflect.Interface)
	//fmt.Println("value:", v.Int())

	// todo: 2. Reflection goes from reflection object to interface value.
	// Like physical reflection, reflection in Go generates its own inverse.
	// Given a reflect.Value we can recover an interface value using the
	// Interface method; in effect the method packs the type and value
	// information back into an interface representation and returns the result
	var n = 6
	v = reflect.ValueOf(n)
	r := v.Interface().(int)
	fmt.Println(r)
	//fmt.Println(v) it's ok but not we want the concrete value it holds
	fmt.Println(v.Interface())

	// todo: 3. To modify a reflection object, the value must be settable.
	// The third law is the most subtle and confusing, but it's easy enough to
	// understand if we start from first principles
	// Just keep in mind that reflection Values need the address of something
	// in order to modify what they represent.

	// we pass a copy of x to reflect.ValueOf, so the interface value created as
	// the argument to reflect.ValueOf is a copy of x, not x itself.
	var f float64 = 3.4
	v = reflect.ValueOf(f)
	// were allowed to succeed, it would not update x, even though v looks like
	// it was created from x. Instead, it would update the copy of x stored
	// inside the reflection value and x itself would be unaffected. That would
	// be confusing and useless, so it is illegal, and settability is the
	// property used to avoid this issue. v.SetFloat(7.1) // Error: will panic
	fmt.Println("settability of v:", v.CanSet())

	var g float64 = 3.4
	p := reflect.ValueOf(&g) // Note: take the address of x.
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())
	// To get to what p points to, we call the Elem method of Value
	v = p.Elem()
	fmt.Println("settability of v:", v.CanSet(), v.Kind(), v.Float())
	v.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(g)

	// fixme: Structs
	// In our previous webexample v wasn't a pointer itself, it was just derived from one. A common way for this situation to arise is when using reflection to modify the fields of a structure. As long as we have the address of the structure, we can modify its fields.
	//	Here's a simple webexample that analyzes a struct value, t. We create the reflection object with the address of the struct because we'll want to modify it later. Then we set typeOfT to its type and iterate over the fields using straightforward method calls (see package reflect for details). Note that we extract the names of the fields from the struct type, but the fields themselves are regular reflect.Value objects.

	type T struct {
		A int
		B string
	}

	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)

}
