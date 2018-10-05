package main

import (
	"fmt"
)

type MyWriter struct{}

func (w *MyWriter) Write(p []byte) (n int, err error) {
	//b.lastRead = opInvalid
	//m, ok := b.tryGrowByReslice(len(p))
	//if !ok {
	//	m = b.grow(len(p))
	//}
	return 12, nil
}


func main() {
	var b MyWriter
	b.Write([]byte("Hello "))

	fmt.Println(b)
}
