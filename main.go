package main

func main() {
	bc := NewBlockchain()
	
	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.blocks {
		block.Print()
	}
}