/*
5.3
switch
*/
package main

import (
	"fmt"
	"os"
)

// const Season (
// 	Spring = iota
// 	Summon
// 	Autumn
// 	Winter
// )

// var SeasonNumber [int]string

func main() {
	for {
		var ia int
		switch fmt.Scanf("%d", &ia); ia {
		case 1:
			fallthrough
		case 2:
			fallthrough
		case 3:
			fmt.Println("Spring")
		case 4:
			fallthrough
		case 5:
			fallthrough
		case 6:
			fmt.Println("Summon")
		case 7:
			fallthrough
		case 8:
			fallthrough
		case 9:
			fmt.Println("Autuom")
		case 10:
			fallthrough
		case 11:
			fallthrough
		case 12:
			fmt.Println("Winter")
		default:
			os.Exit(0)
		}
	}
}
