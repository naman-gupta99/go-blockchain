package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Timestamp int64
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	info := bytes.Join([][]byte{b.Data, b.PrevHash, timestamp}, []byte{})
	hash := sha256.Sum256(info)

	b.Hash = hash[:]
}

func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte{}, []byte(data), prevHash, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func (block *Block) Print() {
	fmt.Printf("Prev. hash: %x\n", block.PrevHash)
	fmt.Printf("Data: %s\n", block.Data)
	fmt.Printf("Hash: %x\n", block.Hash)
}