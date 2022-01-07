package main

import (
	"log"

	"github.com/kadragon/nomadcoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()

	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	chain.AddBlock("Fourth Block")

	for _, block := range chain.Allblocks() {
		log.Println(block.Hash)
	}
}
