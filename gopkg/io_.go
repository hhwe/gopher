package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	p, err := io.Copy(os.Stdout, r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)
}
