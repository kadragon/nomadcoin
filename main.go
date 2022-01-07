package main

import (
	"log"

	"github.com/kadragon/nomadcoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	log.Printf("%#v", chain)
}
