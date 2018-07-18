package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg.Add(10)

	hostURL := "https://studygolang.com/projects"
	for i := 0; i < 10; i++ {
		url := hostURL + "?p=" + string(i)
		go crawl(url)
	}

	wg.Wait()
}

func crawl(url string) {
	fmt.Println("scrapy url")
	defer wg.Done()
	data, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	s, err := ioutil.ReadAll(data.Body)
	if err != nil {
		log.Fatalln(err)
	}
	src := string(s)
	articleRegExp := regexp.MustCompile(`(?Ums)<article class=\"article box_white\">.*</article>`)
	r := articleRegExp.FindAllString(src, -1)
	if r != nil {
		fmt.Println(r[0])
		for i, article := range r {
			titleRegExp := regexp.MustCompile(`(?Ums)<a href="(.*)" target="_blank" title="(?P<title>.*)">.*<span class="view" title="阅读数">.*<span>(?P<vote>.*)</span>`)
			t := titleRegExp.FindStringSubmatch(article)
			// <span class="view" title="阅读数">
			// <i class="glyphicon glyphicon-eye-open"></i>
			// 	阅读:<span>206</span>次
			// </span>
			for j, p := range t {
				fmt.Println(i, j, p)
			}
			// fmt.Println(i, t)
			// fmt.Printf("$title", s)
		}
	}

}
