package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strings"
	"time"
)

var num = 10
var numx2, numx3 int

func main() {
	numx2, numx3 = getX2AndX3(num)
	PrintValues()
	numx2, numx3 = getX2AndX3_2(num)
	PrintValues()
	x1, x2, x3 := get(num)
	fmt.Println(x1, x2, x3)

	fmt.Println(multiRetern1, "  ", multiRetern2)

	fmt.Println(MySqrt(10))
	fmt.Println(MySqrt(-2))

	varargs([]int{1, 2, 3, 4}...)

	fmt.Println(deferReturn())
	logArgs()

	fmt.Printf("%d is even: is %t\n", 16, even(16)) // 16 is even: is true
	fmt.Printf("%d is odd: is %t\n", 17, odd(17))
	// 17 is odd: is true
	fmt.Printf("%d is odd: is %t\n", 18, odd(18))
	// 18 is odd: is false
	fmt.Println(fibonacci(5))
	reverseInt(10)
	fmt.Println(factorial(18))
	fmt.Println(strings.Map(mapping, "sdfj艰苦奋斗副k-0-23家房客的"))
	fv := func() { fmt.Println("hello world") }
	fmt.Printf("%T", fv)
	fc := fibonacci1()
	for i := 0; i < 10; i++ {
		fmt.Println(fc())
	}

	timeIt()

	logWhere()
}

func PrintValues() {
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
}

func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	// return x2, x3
	return
}

func get(input int) (x1, x2, x3 int) {
	x1 = input * 1
	x2 = input * 2
	x3 = input * 3
	return
}

func multiRetern1(x, y int) (int, int, int) {
	return x + y, x * y, x - y
}

func multiRetern2(x, y int) (gentle, accumulate, errand int) {
	gentle = x + y
	accumulate = x * y
	errand = x - y
	return
}

func MySqrt(x float64) (sqrt float64, err error) {
	if x < 0 {
		err = errors.New("negative numbers")
		return
	}
	sqrt = math.Sqrt(x)
	return
}

func varargs(args ...int) {
	for _, i := range args {
		fmt.Println(i)
	}
}

func deferReturn() (i int) {
	defer func() { i++ }()
	return 10 // i=10, i++； defer在return之后调用，最后返回
}

func logArgs() {
	defer log.Println("12")
}

func even(nr int) bool {
	if nr == 0 {
		return true
	}
	return odd(RevSign(nr) - 1)
}

func odd(nr int) bool {
	if nr == 0 {
		return false
	}
	return even(RevSign(nr) - 1)
}

func RevSign(nr int) int {
	if nr < 0 {
		return -nr
	}
	return nr
}

func fibonacci(n int) (pos, value int) {
	if n < 2 {
		return n, 1
	}
	_, value1 := fibonacci(n - 1)
	_, value2 := fibonacci(n - 2)
	pos = n
	value = value1 + value2
	return
}

func fibonacci1() func() int {
	var value1, value2 = 0, 1
	return func() int {
		tmp := value2
		value1, value2 = value2, value1+value2
		return tmp
	}
}

func reverseInt(n int) int {
	if n < 1 {
		return 0
	}
	fmt.Println(n)
	return reverseInt(n - 1)
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func mapping(c rune) rune {
	if c < 255 {
		return c
	}
	return 32
}

func timeIt() {
	var result uint64
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci2(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

func fibonacci2(n int) (res uint64) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci2(n-1) + fibonacci2(n-2)
	}
	return
}

const LIM = 41

var fibs [LIM]uint64

func fibonacci3(n int) (res uint64) {
	// memoization: check if fibonacci(n) is already known in array:
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci3(n-1) + fibonacci3(n-2)
	}
	fibs[n] = res
	return
}

func logWhere() {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()

	log.SetFlags(log.Llongfile)
	log.Print("213")
}
