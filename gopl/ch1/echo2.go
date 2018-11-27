package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s, sep := "", " "
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
		s += arg + sep
	}
	fmt.Println(s)

	fmt.Println(strings.Join(os.Args[:], " "))
}
