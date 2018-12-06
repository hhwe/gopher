package main

// In Bitcoint specification Timestamp, PrevBlockHash, and Hash are block headers, which form a separate data structure, and transactions (Data in our case) is a separate data structure. So we’re mixing them here for simplicity
type Block struct {
	// Timestamp is the current timestamp when the block created
	Timestamp int64
	// Actual valuable information containing in the block
	Data []byte
	// Stores the hash of the previous block
	PrevBlockHash []byte
	// The hash of current block
	Hash []byte
}
