package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	// input := bufio.NewScanner(os.Stdin)
	// for input.Scan() {
	// 	counts[input.Text()]++
	// }

	// // NOTE: ignoring potential errors from input.Err()
	// for line, n := range counts {
	// 	if n > 1 {
	// 		fmt.Printf("%d\t%s\n", n, line)
	// 	}
	// }

	// files := os.Args[1:]
	// if len(files) == 0 {
	// 	countLines(os.Stdin, counts)
	// } else {
	// 	for _, arg := range files {
	// 		f, err := os.Open(arg)
	// 		if err != nil {
	// 			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
	// 			continue
	// 		}
	// 		countLines(f, counts)
	// 	}
	// }

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
