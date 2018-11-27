package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	s, sep := "", " "
	start := time.Now()
	for i := 1; i < 100; i++ {
		for _, arg := range os.Args[:] {
			s += arg + sep
		}
		fmt.Println(s)
	}
	fmt.Println(time.Since(start))

	start1 := time.Now()
	for i := 1; i < 100; i++ {
		fmt.Println(strings.Join(os.Args[:], " "))
	}
	fmt.Println(time.Since(start1))
}
