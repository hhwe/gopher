package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println(time.Now())

	var a uint8 = 127
	var b uint8 = 200
	fmt.Println(a - b)

	http.ListenAndServe(":8000", nil)
}
