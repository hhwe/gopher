/*
4.7
strings
strconv
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var str string = "Hello, how is it going, Hugo? Hello"
	var manyG = "gggggggggg"

	fmt.Println(strings.HasPrefix(str, "Hello"))
	fmt.Println(strings.HasSuffix(str, "Hugo?"))
	fmt.Println(strings.Contains(str, "hello"))
	fmt.Println(strings.Index(str, "h"))
	fmt.Println(strings.LastIndex(str, "H"))
	fmt.Println(strings.IndexRune(str, 72))

	fmt.Println(strings.Replace(str, "Hugo", "Hanhw", -1))
	fmt.Println(str)

	fmt.Println(strings.Count(manyG, "gg"), strings.Count(str, "Hello"))

	fmt.Println(strings.Repeat("ni", 5))

	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.ToUpper(manyG))
	fmt.Println(strings.ToTitle(str))

	fmt.Println(strings.Trim(" fk  j ", " "), "\n",
		strings.TrimLeft(" fk  j ", " "), "\n",
		strings.TrimRight(" fk  j ", " "))

	fmt.Println(strings.Fields("f j k l jk dfs fd "))
	fmt.Println(strings.Join(strings.Split("f,j,k,l,jk,dfs,fd,", ","), manyG))

	fmt.Println(strings.NewReader("some io.Reader stream to be read\n").ReadByte())

	fmt.Println(strconv.Itoa(10))
	fmt.Println(strconv.Atoi("10"))
	fmt.Println(strconv.ParseFloat("123.456", 32))
	fmt.Println(strconv.FormatFloat(123.456, 'b', 10, 64))

	fmt.Println(strconv.IntSize)
}
