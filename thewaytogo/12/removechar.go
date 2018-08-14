package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, err := os.Open("abc.txt")
	if err != nil {
		panic(err)
	}
	outputFile, err := os.OpenFile("abcd.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()
	defer outputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
	for {
		inputString, isprefix, readerError := inputReader.ReadLine()
		fmt.Println(isprefix)
		if readerError == io.EOF {
			return
		}

		outputString := string(inputString[2:5]) + "\r\n"
		n, err := outputWriter.WriteString(outputString)
		fmt.Println(n)
		if err != nil {
			fmt.Println(err)
			return
		}
		outputWriter.Flush()
	}
	fmt.Println("Conversion Done!")
}
