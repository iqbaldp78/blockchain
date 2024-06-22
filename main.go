package main

import (
	"fmt"
	"time"
)

func main() {
	node1 := NewNode("localhost:5000")
	node2 := NewNode("localhost:5001")

	node1.AddPeer(node2.Address)
	node2.AddPeer(node1.Address)

	// Initialize balances
	node1.Blockchain.Balances["Alice"] = 10.0
	node1.Blockchain.Balances["Bob"] = 5.0
	node1.Blockchain.Balances["Charlie"] = 3.0
	node1.Blockchain.Balances["Dave"] = 2.0

	node2.Blockchain.Balances["Alice"] = 10.0
	node2.Blockchain.Balances["Bob"] = 5.0
	node2.Blockchain.Balances["Charlie"] = 3.0
	node2.Blockchain.Balances["Dave"] = 2.0

	go node1.Start()
	go node2.Start()

	time.Sleep(2 * time.Second)

	// Print initial balances
	fmt.Println("Initial balances:")
	for k, v := range node1.Blockchain.Balances {
		fmt.Printf("%s: %f BTC\n", k, v)
	}

	fmt.Println("Node 1 adds a block with multiple transactions")
	transactions := []*Transaction{
		NewTransaction("Alice", "Bob", 1.0),
		NewTransaction("Bob", "Charlie", 2.0),
		NewTransaction("Charlie", "Dave", 3.0),
	}
	node1.AddBlock(transactions)

	time.Sleep(2 * time.Second)

	fmt.Println("Node 2 requests synchronization")
	node2.RequestSync()

	time.Sleep(2 * time.Second)

	fmt.Println("Node 2 Blockchain after synchronization:")
	for _, block := range node2.Blockchain.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		for _, tx := range block.Transactions {
			fmt.Printf("Transaction: %s -> %s: %f BTC\n", tx.Sender, tx.Receiver, tx.Amount)
		}
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}

	fmt.Println("Balances after first synchronization:")
	for k, v := range node2.Blockchain.Balances {
		fmt.Printf("%s: %f BTC\n", k, v)
	}

	// Validate blockchain in Node 2
	if node2.Blockchain.Validate() {
		fmt.Println("Node 2 Blockchain is valid")
	} else {
		fmt.Println("Node 2 Blockchain is not valid")
	}

	// Node 2 creates a new transaction from Charlie to Mike
	fmt.Println("Node 2 adds a new block with a transaction from Charlie to Mike")
	newTransaction := []*Transaction{
		NewTransaction("Charlie", "Mike", 1.5),
	}
	node2.Blockchain.Balances["Mike"] = 0.0 // Initialize Mike's balance
	node2.AddBlock(newTransaction)

	time.Sleep(2 * time.Second)

	fmt.Println("Node 1 requests synchronization")
	node1.RequestSync()

	time.Sleep(2 * time.Second)

	// Print only the new transaction added with prev and current hash
	fmt.Println("Incoming next transaction in Node 1:")
	lastBlock := node1.Blockchain.Blocks[len(node1.Blockchain.Blocks)-1]
	fmt.Printf("Prev. hash: %x\n", lastBlock.PrevBlockHash)
	for _, tx := range lastBlock.Transactions {
		fmt.Printf("Transaction: %s -> %s: %f BTC\n", tx.Sender, tx.Receiver, tx.Amount)
	}
	fmt.Printf("Hash: %x\n", lastBlock.Hash)
	fmt.Println()

	fmt.Println("Balances after second synchronization:")
	for k, v := range node1.Blockchain.Balances {
		fmt.Printf("%s: %f BTC\n", k, v)
	}
}
