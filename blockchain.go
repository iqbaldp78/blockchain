package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"time"
)

// Blockchain represents the entire chain
type Blockchain struct {
	blocks []*Block
}

// AddBlock saves the block into the blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// NewBlockchain creates a new Blockchain with a genesis block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

// Validate checks the integrity of the blockchain
func (bc *Blockchain) Validate() bool {
	for i := 1; i < len(bc.blocks); i++ {
		currentBlock := bc.blocks[i]
		prevBlock := bc.blocks[i-1]

		// Recalculate the hash of the current block
		timestamp := []byte(time.Unix(currentBlock.Timestamp, 0).String())
		headers := bytes.Join([][]byte{currentBlock.PrevBlockHash, currentBlock.Data, timestamp}, []byte{})
		hash := sha256.Sum256(headers)

		// Check if the stored hash is correct
		if !bytes.Equal(currentBlock.Hash, hash[:]) {
			fmt.Printf("Invalid hash at block %d\n", i)
			return false
		}

		// Check if the current block points to the correct previous block
		if !bytes.Equal(currentBlock.PrevBlockHash, prevBlock.Hash) {
			fmt.Printf("Invalid previous hash at block %d\n", i)
			return false
		}
	}

	return true
}
