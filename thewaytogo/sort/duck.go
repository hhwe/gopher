package sort

import (
	"errors"
	"fmt"
	"strconv"
)

func p() {
	b := new(Bird)
	DuckDance(b)

	type any interface{}
	var li = []any{1, "a", 2, "b"}
	fmt.Println(li)

	mapFunction()
}

type Stringer interface {
	String() string
}

type Celsius float64

func (c Celsius) String() string {
	return strconv.FormatFloat(float64(c), 'f', 1, 64) + " *C"
}

type Day int

var dayName = []string{""}

type IDuck interface {
	Quack()
	Walk()
}

func DuckDance(duck IDuck) {
	for i := 1; i <= 3; i++ {
		duck.Quack()
		duck.Walk()
	}
}

type Bird struct {
}

func (b *Bird) Quack() {
	fmt.Println("I am quacking!")
}

func (b *Bird) Walk() {
	fmt.Println("I am walking!")
}

type obj interface{}

func mapFunction() {
	// define a generic lambda function mf:
	mf := func(i obj) obj {
		switch i.(type) {
		case int:
			return i.(int) * 2
		case string:
			return i.(string) + i.(string)
		}
		return i
	}

	isl := []obj{0, 1, 2, 3, 4, 5}
	res1 := mapFunc(mf, isl)
	for _, v := range res1 {
		fmt.Println(v)
	}
	println()

	ssl := []obj{"0", "1", "2", "3", "4", "5"}
	res2 := mapFunc(mf, ssl)
	for _, v := range res2 {
		fmt.Println(v)
	}
}

func mapFunc(mf func(obj) obj, list []obj) []obj {
	result := make([]obj, len(list))

	for ix, v := range list {
		result[ix] = mf(v)
	}

	// Equivalent:
	/*
		for ix := 0; ix<len(list); ix++ {
			result[ix] = mf(list[ix])
		}
	*/
	return result
}

// type Stack interface {
// 	Len() int
// 	IsEmpty() bool
// 	Push(x interface{})
// 	Pop() (x interface{}, error)
// }

type Stack []interface{}

func (stack Stack) Len() int {
	return len(stack)
}

func (stack Stack) Cap() int {
	return cap(stack)
}

func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}

func (stack *Stack) Push(e interface{}) {
	*stack = append(*stack, e)
}

func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("stack is empty")
	}
	return stack[len(stack)-1], nil
}

func (stack *Stack) Pop() (interface{}, error) {
	stk := *stack // dereference to a local variable stk
	if len(stk) == 0 {
		return nil, errors.New("stack is empty")
	}
	top := stk[len(stk)-1]
	*stack = stk[:len(stk)-1] // shrink the stack
	return top, nil
}
