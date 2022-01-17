package main

import (
	"github.com/kadragon/nomadcoin/blockchain"
	"github.com/kadragon/nomadcoin/cli"
)

func main() {
	blockchain.Blockchain()
	cli.Start()
}
