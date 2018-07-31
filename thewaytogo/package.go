package main

import (
	"container/list"
	"fmt"
	"math"
	"math/big"
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
	bigMath()
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
	sync.RWMutex
	mp map[string]int
}

func async() {
	var wg sync.WaitGroup
	var g = 20
	wg.Add(g)
	m := map[string]int{
		"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98,
	}
	for i := 0; i < g/2; i++ {
		go changeMap(m, &wg)
	}
	mu := muxMap{mp: m}
	for i := 0; i < g/2; i++ {
		go changeMutexMap(&mu, &wg)
	}
	wg.Wait()
}

func changeMap(m map[string]int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("======")
	for k, v := range m {
		v += 10
		fmt.Printf("%s, %d;", k, v)
	}
	fmt.Println()
	fmt.Println("======")
}

func changeMutexMap(m *muxMap, wg *sync.WaitGroup) {
	defer wg.Done()
	m.Lock()
	defer m.Unlock()
	fmt.Println("======")
	for k, v := range m.mp {
		v += 10
		fmt.Printf("%s, %d;", k, v)
	}
	fmt.Println()
	fmt.Println("======")
}

func bigMath() {
	// Here are some calculations with bigInts:
	fmt.Println(math.MaxInt8)
	im := big.NewInt(math.MaxInt64)
	in := im
	io := big.NewInt(1956)
	ip := big.NewInt(1)
	ip.Mul(im, in).Add(ip, im).Div(ip, io)
	fmt.Printf("Big Int: %v\n", ip)
	// Here are some calculations with bigInts:
	rm := big.NewRat(math.MaxInt64, 1956)
	rn := big.NewRat(-1956, math.MaxInt64)
	ro := big.NewRat(19, 56)
	rp := big.NewRat(1111, 2222)
	rq := big.NewRat(1, 1)
	rq.Mul(rm, rn).Add(rq, ro).Mul(rq, rp)
	fmt.Printf("Big Rat: %v\n", rq)
}
