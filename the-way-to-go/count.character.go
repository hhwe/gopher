/*

 */

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

var (
	r = "asSASA ddd dsjkdsjs dk"
	s = "asSASA ddd dsjkdsjsこん dk"
)

func main() {

	fmt.Println(len(r), len(s))
	fmt.Println(utf8.RuneCountInString(r), utf8.RuneCountInString(s))

	r := strings.NewReader("some io.Reader stream to be read\n")
	p, err := io.Copy(os.Stdout, r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)
}
