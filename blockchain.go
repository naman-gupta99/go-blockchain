package main

type BlockChain struct {
	blocks []*Block
}

func (chain *BlockChain) AddBlock(data string) {
	prevHash := chain.blocks[len(chain.blocks)-1].Hash
	block := NewBlock(data, prevHash)

	chain.blocks = append(chain.blocks, block)
}

func NewBlockchain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}