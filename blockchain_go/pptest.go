package main

import (
	"math/big"
	"fmt"
	"math"
	"bytes"
	"crypto/sha256"
)

var (
	maxNonce = math.MaxInt64
)


func main() {
	args := big.NewInt(1)
	fmt.Println(args)
	args.Lsh(args, uint(256 - 24))
	fmt.Printf("%x\n", args)
	p := "10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	fmt.Println(len(p))
	fmt.Printf("%x", maxNonce)

	fmt.Println()

	var hInt big.Int
	var hash [32]byte
	data := bytes.Join(
		[][]byte{},
		[]byte{},
	)
	fmt.Println(data)
	hash = sha256.Sum256(data)
	fmt.Println(hash)
	hInt.SetBytes(hash[:])
	fmt.Println(hInt)
	hInt.Cmp(args)
}
