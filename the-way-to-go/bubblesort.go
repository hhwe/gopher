package main

import (
	"fmt"
)

func main() {
	p := []int{9, 5, 3, 4, 8, 3, 7, 2, 1, 6}
	fmt.Println(bubllesort(p))
	fmt.Println(mapFunc(multi, p))
}

func bubllesort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func mapFunc(fn func(int, int) int, a []int) []int {
	for i := range a {
		a[i] = fn(a[i], 10)
	}
	return a
}

func multi(a, b int) int {
	a *= b
	return a
}
