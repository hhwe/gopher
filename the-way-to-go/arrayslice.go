package main

import (
	"bytes"
	"fmt"
)

func main() {
	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}

	arrayValue()

	for i := 0; i < LIM; i++ {
		fmt.Println(fibonacci(i))
	}

	bufferByte()

	mk := make([]int, 3)
	pk := append(mk, 1, 2, 3)
	fmt.Println(pk)

	// sl := make([]byte, 3)
	var sl = []byte{1, 2, 3}
	dt := make([]byte, 3)
	sd := sAppend(sl, dt)
	fmt.Println(sd)

	var buf bytes.Buffer
	fmt.Println(buf.Bytes())

	sliceFunc()
}

func arrayValue() {
	var arr [15]int
	for i := 0; i < 15; i++ {
		arr[i] = i
	}
	fv := func(arr [15]int) {
		arr[0] = 100
	}
	fv(arr)
	fmt.Println(arr)
	fc := func(arr *[15]int) {
		arr[0] = 100
	}
	fc(&arr)
	fmt.Println(arr)
}

const LIM = 41

var fibs [LIM]uint64

func fibonacci(n int) (res uint64) {
	// memoization: check if fibonacci(n) is already known in array:
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibs[n] = res
	return
}

func bufferByte() {
	var buffer bytes.Buffer
	buffer.WriteString("hello world")
	fmt.Print(buffer.String(), "\n")
}

func sAppend(slice, data []byte) []byte {
	for _, i := range data {
		slice = append(slice, i)
	}
	return slice
}

func sliceFunc() {
	items := [...]int{10, 20, 30, 40, 50}
	for _, item := range items {
		item *= 2
	}
	fmt.Println(items)
	for i := range items {
		items[i] *= 2
	}
	fmt.Println(items)

	sli := []int{10, 20, 30, 40, 50}
	sh(sli...)
	sumArray()
	sumSlice(sli)

	sumAndAverage(sli)
	minSlice(sli)
	maxSlice(sli)

	out()

	fmt.Println(ext(sli, 5))
}

func sh(a ...int) {
	fmt.Println(a)
}

func sumArray() {
	arr := [10]float32{1, 2, 3}
	var sum float32
	for _, i := range arr {
		sum += i
	}
	fmt.Println(sum)
}

func sumSlice(sli []int) {
	sum := 0
	for _, i := range sli {
		sum += i
	}
	fmt.Println(sum)
}

func sumAndAverage(sli []int) (int, float32) {
	sum := 0
	for _, i := range sli {
		sum += i
	}
	fmt.Println(sum)

	var average float32
	average = float32(sum / len(sli))
	fmt.Println(sum, average)
	return sum, average
}

func minSlice(sli []int) int {
	min := 0
	// min = math.MaxInt3 // why??
	for _, i := range sli {
		if i < min {
			min = i
		}
	}
	fmt.Println(min)
	return min
}

func maxSlice(sli []int) int {
	max := 0
	for _, i := range sli {
		if i > max {
			max = i
		}
	}
	fmt.Println(max)
	return max
}

func out() {
	arr := [5]int{10, 20, 30, 40, 50}
	sli := arr[:]
	p := append(sli, 2, 3, 4, 5)
	sli[3] = 100
	p[1] = 200
	fmt.Println(arr, sli, p)
}

func ext(sli []int, fac int) []int {
	l := len(sli) * fac
	if cap(sli) > l {
		o := make([]int, l-len(sli))
		return append(sli, o...)
	}
	newSli := make([]int, l)
	copy(newSli, sli)
	return newSli
}
