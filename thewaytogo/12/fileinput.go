package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	file, inputError := os.Open("readinput.go")
	if inputError != nil {
		panic("error")
	}
	defer file.Close()

	inputReader := bufio.NewReader(file)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		if readerError == io.EOF {
			break
		}
	}
	for {
		buf := make([]byte, 1024)
		n, err := inputReader.Read(buf)
		if n == 0 || err == nil {
			break
		}
		fmt.Printf("%d, %s", n, string(buf))
	}

	inputFile := "readinput.go"
	outPutFile := "tmp.go"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Eroor: %s\n", err)
	}

	err = ioutil.WriteFile(outPutFile, buf, 0644)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(filepath.Base("/"))
	os.Stdout.WriteString("hello, world\n")
}
