package main

import (
	"net/http"
	"os"
	"log"
	"io"
)

func main() {
	r, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()

	file, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	//p := make([]byte, 228)
	//n, err := r.Body.Read(p)
	//fmt.Println(n, string(p))

	dest := io.MultiWriter(os.Stdout, file)
	io.Copy(dest, r.Body)

	//body, err := ioutil.ReadAll(r.Body)
	//fmt.Println(string(body))
	//io.Copy(dest, r.Body)
	//file.Write([]byte(r.Body))
}
