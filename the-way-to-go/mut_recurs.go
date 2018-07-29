package main

import (
	"fmt"
	"strings"
)

func main() {
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
