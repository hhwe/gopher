package main

import (
	"fmt"

	"./revstr"
)

func main() {
	i := "12345"
	fmt.Printf("%s => %s\n", i, revstr.Reverse(i))

}
