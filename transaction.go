package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Transaction struct {
	ID        string
	Sender    string
	Receiver  string
	Amount    float64
	Timestamp int64
}

func NewTransaction(sender, receiver string, amount float64) *Transaction {
	tx := &Transaction{Sender: sender, Receiver: receiver, Amount: amount, Timestamp: time.Now().Unix()}
	tx.ID = tx.Hash()
	return tx
}

func (tx *Transaction) Hash() string {
	record := string(rune(tx.Timestamp)) + tx.Sender + tx.Receiver + fmt.Sprintf("%f", tx.Amount)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}
