package main

import (
	"math/big"
	"fmt"
)

const bits = 24

func main() {
	target := big.NewInt(1)
	fmt.Printf("%T, %d\n", target, target)
	target.Lsh(target, uint(256-bits))
	fmt.Println(target)
}
