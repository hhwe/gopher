package main

import (
	"fmt"
)

func main() {
	var a = 123
	var b = true
	// var c = "a"
	var f = 12.1234567879
	var s = "hello world!"
	var i interface{} = 23

	fmt.Printf("%v, %#v, %T, %% \n", a, a, a)
	fmt.Printf("%10d, %10.2d, %2d \n", a, a, a)
	fmt.Printf("%t \n", b)
	fmt.Printf("%b, %d, %o, %q, %x, %X, %U \n", a, a, a, a, a, a, a)
	fmt.Printf("%#o, %#q, %#x, %#X, %#U \n", a, a, a, a, a)
	// %c
	fmt.Printf("%b, %e, %E, %f, %F, %g, %G \n", f, f, f, f, f, f, f)
	fmt.Printf("%10.2f, %2.10f, %0.10f, %10.0f \n", f, f, f, f)
	fmt.Printf("%s, %q, %x, %X \n", s, s, s, s)
	fmt.Printf("%p, %p, %p, %p \n", &a, &b, &f, &s)
	fmt.Printf("%+v, %-v, %#v, %v, %0v \n", a, f, a, a, a)
	fmt.Printf("%v %d\n", i, i)

	fmt.Printf("%[2]d %[1]d\n", 11, 22)
	fmt.Printf("%[3]*.[2]*[1]f\n", 12.0, 2, 6) // equivalent to Sprintf("%6.2f", 12.0)
	fmt.Printf("%d %d %#[1]x %#x\n", 16, 17)

	var sa string
	var ia int64
	fmt.Scan(&sa, &ia)
	fmt.Printf("%s %d", sa, ia)

	fmt.Sscan(sa, ia)

	fmt.Errorf("error", sa, ia)
	// w := ""
	// fmt.Fprint(w, "123")
}
