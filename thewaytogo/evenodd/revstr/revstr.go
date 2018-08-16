package revstr

import (
	"fmt"
)

func Reverse(s string) (str string) {
	b := []rune(s)
	fmt.Println(b)
	l := len(b)
	for i := 0; i < l/2; i++ {
		b[i], b[l-i-1] = b[l-i-1], b[i]
	}
	fmt.Println(b)
	str = string(b)
	return
}
