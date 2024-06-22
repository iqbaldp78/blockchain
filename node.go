package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
)

type Node struct {
	Address    string
	Blockchain *Blockchain
	Peers      []string
}

func NewNode(address string) *Node {
	return &Node{
		Address:    address,
		Blockchain: NewBlockchain(),
		Peers:      []string{},
	}
}

func (node *Node) AddPeer(address string) {
	node.Peers = append(node.Peers, address)
}

func (node *Node) Start() {
	listener, err := net.Listen("tcp", node.Address)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go node.HandleConnection(conn)
	}
}

func (node *Node) HandleConnection(conn net.Conn) {
	defer conn.Close()

	var buf bytes.Buffer
	_, err := buf.ReadFrom(conn)
	if err != nil {
		panic(err)
	}

	msg := DeserializeMessage(buf.Bytes())
	switch msg.Type {
	case "block":
		block := DeserializeBlock(msg.Payload)
		if node.Blockchain.AddBlockFromBlock(block) {
			node.BroadcastBlock(block)
		}
	case "sync":
		node.SendBlockchain(conn)
	case "blockchain":
		blockchain := DeserializeBlockchain(msg.Payload)
		if node.Blockchain.Update(blockchain) {
			fmt.Println("Blockchain synchronized")
		} else {
			fmt.Println("Failed to synchronize blockchain")
		}
	}
}

func (node *Node) SendBlock(conn net.Conn, block *Block) {
	msg := NewMessage("block", block.Serialize())
	conn.Write(msg.Serialize())
}

func (node *Node) BroadcastBlock(block *Block) {
	for _, peer := range node.Peers {
		conn, err := net.Dial("tcp", peer)
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer conn.Close()
		node.SendBlock(conn, block)
	}
}

func (node *Node) AddBlock(transactions []*Transaction) {
	node.Blockchain.AddBlock(transactions)
	block := node.Blockchain.Blocks[len(node.Blockchain.Blocks)-1]
	node.BroadcastBlock(block)
}

func (node *Node) RequestSync() {
	for _, peer := range node.Peers {
		conn, err := net.Dial("tcp", peer)
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer conn.Close()

		msg := NewMessage("sync", nil)
		conn.Write(msg.Serialize())
	}
}

func (node *Node) SendBlockchain(conn net.Conn) {
	msg := NewMessage("blockchain", node.Blockchain.Serialize())
	conn.Write(msg.Serialize())
}

type Message struct {
	Type    string
	Payload []byte
}

func NewMessage(msgType string, payload []byte) *Message {
	return &Message{
		Type:    msgType,
		Payload: payload,
	}
}

func (msg *Message) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(msg)
	if err != nil {
		panic(err)
	}
	return result.Bytes()
}

func DeserializeMessage(d []byte) *Message {
	var msg Message
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&msg)
	if err != nil {
		panic(err)
	}
	return &msg
}
