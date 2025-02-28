package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

const (
	a = iota
	b
	c
	d = a | b
	e
	f
	g
	h = iota
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	file, err := os.OpenFile("error.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	Trace = log.New(ioutil.Discard, "TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(os.Stdout, "INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(os.Stderr, "WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(file, os.Stderr), "TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

func main() {
	Trace.Println("I have something standard to say")
	Info.Println("Special Information")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")
	//log.Panicln("abc")
	//log.Fatalln("cdf")
	log.Println(a, b, c, d, e, f, g, h)

	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	log.Panicln("panic")

}
