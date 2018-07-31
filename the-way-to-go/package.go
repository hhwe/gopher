package main

import (
	"container/list"
	"fmt"
	"regexp"
	"strconv"
	"sync"
	"unsafe"
)

func main() {
	doubleLinkList()
	sizeOfVar()
	regexExp()

	async()
}

func doubleLinkList() {
	l := list.New()
	z := l.PushFront(101)
	l.PushFront(102)
	l.PushFront(103)
	l.PushBack(201)
	l.PushBack(202)
	a := l.InsertAfter(500, z)
	b := l.InsertBefore(400, z)
	l.MoveAfter(b, a)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func sizeOfVar() {
	fmt.Println(unsafe.Sizeof(1))
	fmt.Println(unsafe.Sizeof("1"))
	fmt.Println(unsafe.Sizeof(true))
	fmt.Println(unsafe.Sizeof(1.2))
	fmt.Println(unsafe.Sizeof([]int{1, 2, 3, 4}))
}

func regexExp() {
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+" //正则

	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}

	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("Match Found!")
	}

	re, _ := regexp.Compile(pat)
	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)
	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)

	rp := regexp.MustCompile(pat)
	stp := rp.FindAllString(searchIn, -1)
	fmt.Println(stp)
	stp2 := rp.FindAllStringSubmatch(searchIn, -1)
	fmt.Println(stp2)
}

type muxMap struct {
	mu sync.RWMutex
	mp map[string]int
}

func async() {
	var wg sync.WaitGroup
	wg.Add(1)
	m := map[string]int{
		"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98,
	}
	go changeMap(m, wg)
	wg.Wait()
}

func changeMap(m map[string]int, wg sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("======")
	for k, v := range m {
		v += 10
		fmt.Println(k, v)
	}
	fmt.Println("======")
}
