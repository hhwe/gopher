package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var NewLine = flag.Bool("n", false, "print newline")

const (
	Space   = " "
	Newline = "\n"
)

func main() {
	flag.PrintDefaults()
	flag.Parse()
	fmt.Println()
	s := ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "
			if *NewLine {
				s += Newline
			}
		}
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)

	who := "Alice"
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Good Morning: ", who)

	f := Page{"tmp.go", []byte("hello, world")}
	f.save()
	f.load("tmp.go")

	CopyFile("tmp.go", "fileinput.go")
	fmt.Println("Copy Done!")
}

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() (err error) {
	return ioutil.WriteFile(p.Title, p.Body, 0666)
}

func (p *Page) load(title string) (err error) {
	p.Title = title
	p.Body, err = ioutil.ReadFile(p.Title)
	return err
}

func CopyFile(dstName, srcName string) (wirtten int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
