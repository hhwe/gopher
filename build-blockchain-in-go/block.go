package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// In Bitcoint specification Timestamp, PrevBlockHash, and Hash are block headers,
// which form a separate data structure,
// and transactions (Data in our case) is a separate data structure.
// So we’re mixing them here for simplicity
type Block struct {
	// Timestamp is the current timestamp when the block created
	Timestamp int64
	// Stores the hash of the previous block
	PrevBlockHash []byte
	// The hash of current block
	Hash []byte
	// Actual valuable information containing in the block
	Data []byte
}

// For now, we’ll just take block fields, concatenate them,
// and calculate a SHA-256 hash on the concatenated combination.
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp},
		[]byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// Creation of a block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		Timestamp:     time.Now().Unix(),
		PrevBlockHash: prevBlockHash,
		Data:          []byte(data),
	}
	block.SetHash()
	return &block
}
