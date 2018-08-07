package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"../sort"
)

func main() {
	p1()

	// // 练习 12.1: word_letter_count.go
	// // 编写一个程序，从键盘读取输入。当用户输入 'S' 的时候表示输入结束，这时程序输出 3 个数字：
	// // i) 输入的字符的个数，包括空格，但不包括 '\r' 和 '\n'
	// // ii) 输入的单词的个数
	// // iii) 输入的行数
	// p2()

	// 练习 12.2: calculator.go
	// 编写一个简单的逆波兰式计算器，它接受用户输入的整型数（最大值 999999）和运算符 +、-、*、/。
	// 输入的格式为：number1 ENTER number2 ENTER operator ENTER --> 显示结果
	// 当用户输入字符 'q' 时，程序结束。请使用您在练习11.3中开发的 stack 包。
	p3()
}

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

func p1() {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	// fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
	// 输出结果: From the string we read: 56.12 5212 Go

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	fmt.Printf("Your name is %s", input)
	// For Unix: test with delimiter "\n", for Windows: test with "\r\n"
	switch input {
	case "Philip\n":
		fmt.Println("Welcome Philip!")
	case "Chris\r\n":
		fmt.Println("Welcome Chris!")
	case "Ivo\r\n":
		fmt.Println("Welcome Ivo!")
	default:
		fmt.Printf("You are not welcome here! Goodbye!")
	}

	// version 2:
	switch input {
	case "Philip\r\n":
		fallthrough
	case "Ivo\r\n":
		fallthrough
	case "Chris\r\n":
		fmt.Printf("Welcome %s\n", input)
	default:
		fmt.Printf("You are not welcome here! Goodbye!\n")
	}

	// version 3:
	switch input {
	case "Philip\r\n", "Ivo\r\n":
		fmt.Printf("Welcome %s\n", input)
	default:
		fmt.Printf("You are not welcome here! Goodbye!\n")
	}
}

func p2() {
	input := bufio.NewReader(os.Stdin)
	in, err := input.ReadString('S')
	if err != nil {
		panic("input error")
	}
	fmt.Println(in)

	p := 0
	q := 1
	l := 1
	for _, i := range in {
		fmt.Println(i)
		if i != '\r' && i != '\n' {
			p++
		}
		if i == '\n' {
			l++
		}
	}
	q = len(strings.Fields(in))
	fmt.Println(p, q, l)
}

func p3() {
	input := bufio.NewReader(os.Stdin)
	in, err := input.ReadString('q')
	if err != nil {
		panic("input error")
	}
	s := Stack{}
	format = "%f\n%f\n%c\n"
	var a, b float64
	var c char
	fmt.Sscanf(in, "%f\n%f\n%c\n", a, b, c)
}
