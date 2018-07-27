package main

import (
	"fmt"
)

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday = 1
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println("123ß", Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}
