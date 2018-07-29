/*
5.4
for
*/
package main

import (
	"fmt"
	"strings"
)

// const Season (
// 	Spring = iota
// 	Summon
// 	Autumn
// 	Winter
// )

// var SeasonNumber [int]string

func main() {
	i := 0
	for ; i < 15; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println(i)

add:
	if i < 30 {
		fmt.Println(i)
		i++
		goto add
	}
	fmt.Println(i)

	fmt.Println()

	for i := 1; i <= 25; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("G")
		}
		fmt.Println()
	}

	str := "G"
	for i := 0; i < 25; i++ {
		fmt.Println(str)
		str += "G"
	}

	for i := 1; i <= 25; i++ {
		fmt.Println(strings.Repeat("G", i))
	}

	fmt.Println(i) // 块作用域定义的变量和全局变量没影响

	for i := 0; i <= 10; i++ {
		fmt.Printf("%10b >>> %10b\n", i, ^i)
	}

	k := 0
	for k < 100 {
		// if i%15 == 0 {
		// 	fmt.Println("FizzBuzz")
		// } else if i%3 == 0 {
		// 	fmt.Println("Buzz")
		// } else if i%5 == 0 {
		// 	fmt.Println("Buzz")
		// } else {
		// 	fmt.Println(i)
		// }
		switch {
		case k%15 == 0:
			fmt.Println("FizzBuzz")
		case k%3 == 0:
			fmt.Println("Buzz")
		case k%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(k)
		}
		k++
	}

	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j, s = i+1, j+1, s+"a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}
}
