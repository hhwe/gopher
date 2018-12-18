package main

import (
	"github.com/boltdb/bolt"
)

const (
	dbFile       = "blockchain.db"
	blocksBucket = "blocks"
)

// In its essence Blockchain is just a database with certain structure:
// it’s an ordered, back-linked list.
// Which means that blocks are stored in the insertion order
// and that each block is linked to the previous one.
// This structure allows to quickly get the latest block in a chain
// and to (efficiently) get a block by its hash.
type Blockchain struct {
	// blocks []*Block
	tip []byte
	db  *bolt.DB
}

// AddBlock add a new block to Blockchain
// func (bc *Blockchain) AddBlock(data string) {
// 	prevBlock := bc.blocks[len(bc.blocks)-1]
// 	newBlock := NewBlock(data, prevBlock.Hash)
// 	bc.blocks = append(bc.blocks, newBlock)
// }
func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte

	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})
	if err != nil {
		panic(err)
	}

	newBlock := NewBlock(data, lastHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			panic(err)
		}
		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			panic(err)
		}
		bc.tip = newBlock.Hash

		return nil
	})
	if err != nil {
		panic(err)
	}
}

// NewGenesisBlock is To add a new block we need an existing block,
// but there’re not blocks in our Blockchain!
// So, in any Blockchain, there must be at least one block,
// and such block, the first in the chain, is called genesis block.
// Let’s implement a method that creates such a block:
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// NewBlockchain creates a Blockchain with the genesis block:
func NewBlockchain() *Blockchain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				panic(err)
			}
			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				panic(err)
			}
			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil {
				panic(err)
			}
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	bc := Blockchain{tip, db}

	return &bc
}

// Iterator create a BlockchainIterator structor.
func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.db}

	return bci
}

// BlockchainIterator iterate over blocks in a blockchain.
type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

// Next will return the next block from a blockchain.
func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})
	if err != nil {
		panic(err)
	}

	i.currentHash = block.PrevBlockHash

	return block
}
