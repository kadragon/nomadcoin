package main

import (
	"github.com/kadragon/nomadcoin/explorer"
	"github.com/kadragon/nomadcoin/rest"
)

func main() {
	go rest.Start(3000)
	explorer.Start(3001)
}
