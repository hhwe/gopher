package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "https://www.baidu.com/"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)

	p := make([]byte, 100)
	n, err := response.Body.Read(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(n, string(p))
}
