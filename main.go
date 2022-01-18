package main

import (
	"github.com/kadragon/nomadcoin/cli"
	"github.com/kadragon/nomadcoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
