package main

import (
	"fmt"
)

const (
	true  = 0 == 0
	false = 0 != 0
	a     = iota
	b
	c
	d
)

func main() {
	s := "123"
	fmt.Println(s)
	fmt.Println(true, false, nil, a, b, c, d)

	p := make([]int, 5, 8)
	l := append(p, 1, 2, 3)
	fmt.Println(p, len(p), cap(p))
	fmt.Println(l, len(l), cap(l))

	ch := make(chan int, 1)
	ch <- 12
	fmt.Println(<-ch)
	close(ch)

	cmp := complex(1, 2)
	fmt.Println(cmp, imag(cmp), real(cmp))

	mp := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(mp)
	delete(mp, "a")
	fmt.Println(mp)

	type i int
	z := new(i)
	fmt.Println(z, *z)

	print("print ")
	println("println")

	pr()
}

func pr() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("123")
}
