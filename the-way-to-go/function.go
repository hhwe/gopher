package main

import (
	"errors"
	"fmt"
	"log"
	"math"
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
