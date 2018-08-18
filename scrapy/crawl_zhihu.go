package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://www.zhihu.com/"
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	request.Header.Set("cookie", `@@@`)

	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Status)

	html, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("zhihu.html", html, 0666)
	if err != nil {
		panic(err)
	}
}
