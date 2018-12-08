package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

var (
	maxNonce = math.MaxInt64
)

// targetBits is the block header storing the difficulty
// at which the block was mined.
const targetBits = 24

// ProofOfWork holds a pointer to a block and a pointer to a target.
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// NewProofOfWork build and return a new proof of work.
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	return &ProofOfWork{block: b, target: target}
}

// prepareData is merge block fields with the target and nonce.
// nonce here is the counter from the Hashcash description above,
// this is a cryptographic term.
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	return bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
}

// Run implement the core of the PoW algorithm.
// First, we initialize variables: hashInt is the integer representation of hash; nonce is the counter.
// Next, we run an “infinite” loop: it’s limited by maxNonce, which equals to math.MaxInt64; this is done to avoid a possible overflow of nonce.
// Although the difficulty of our PoW implementation is too low for the counter to overflow, it’s still better to have this check, just in case.
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}
