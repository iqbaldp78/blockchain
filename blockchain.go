package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"time"
)

type Blockchain struct {
	Blocks   []*Block
	Balances map[string]float64
}

func (bc *Blockchain) AddBlock(transactions []*Transaction) {
	for _, tx := range transactions {
		bc.Balances[tx.Sender] -= tx.Amount
		bc.Balances[tx.Receiver] += tx.Amount
	}
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(transactions, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) AddBlockFromBlock(block *Block) bool {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	if !bytes.Equal(block.PrevBlockHash, prevBlock.Hash) {
		return false
	}
	for _, tx := range block.Transactions {
		bc.Balances[tx.Sender] -= tx.Amount
		bc.Balances[tx.Receiver] += tx.Amount
	}
	bc.Blocks = append(bc.Blocks, block)
	return true
}

func (bc *Blockchain) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(bc)
	if err != nil {
		panic(err)
	}
	return result.Bytes()
}

func DeserializeBlockchain(d []byte) *Blockchain {
	var blockchain Blockchain
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&blockchain)
	if err != nil {
		panic(err)
	}
	return &blockchain
}

func (bc *Blockchain) Update(newBlockchain *Blockchain) bool {
	if len(newBlockchain.Blocks) <= len(bc.Blocks) {
		return false
	}

	for i := 1; i < len(newBlockchain.Blocks); i++ {
		currentBlock := newBlockchain.Blocks[i]
		prevBlock := newBlockchain.Blocks[i-1]

		timestamp := []byte(time.Unix(currentBlock.Timestamp, 0).String())
		var txs []byte
		for _, tx := range currentBlock.Transactions {
			txs = append(txs, []byte(tx.Sender+tx.Receiver+fmt.Sprintf("%f", tx.Amount))...)
		}
		headers := bytes.Join([][]byte{currentBlock.PrevBlockHash, txs, timestamp}, []byte{})
		hash := sha256.Sum256(headers)

		if !bytes.Equal(currentBlock.Hash, hash[:]) {
			return false
		}

		if !bytes.Equal(currentBlock.PrevBlockHash, prevBlock.Hash) {
			return false
		}
	}

	bc.Blocks = newBlockchain.Blocks
	bc.Balances = newBlockchain.Balances
	return true
}

func NewBlockchain() *Blockchain {
	genesisBlock := NewGenesisBlock()
	return &Blockchain{
		Blocks:   []*Block{genesisBlock},
		Balances: map[string]float64{"genesis": 0},
	}
}

func (bc *Blockchain) Validate() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		prevBlock := bc.Blocks[i-1]

		timestamp := []byte(time.Unix(currentBlock.Timestamp, 0).String())
		var txs []byte
		for _, tx := range currentBlock.Transactions {
			txs = append(txs, []byte(tx.Sender+tx.Receiver+fmt.Sprintf("%f", tx.Amount))...)
		}
		headers := bytes.Join([][]byte{currentBlock.PrevBlockHash, txs, timestamp}, []byte{})
		hash := sha256.Sum256(headers)

		if !bytes.Equal(currentBlock.Hash, hash[:]) {
			return false
		}

		if !bytes.Equal(currentBlock.PrevBlockHash, prevBlock.Hash) {
			return false
		}
	}
	return true
}
