package main

import (
	"fmt"

	"./even"
)

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Printf("Is the inteer %d even? %v\n", i, even.Even(i))
	}
}
