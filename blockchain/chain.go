package blockchain

import (
	"fmt"
	"sync"

	"github.com/kadragon/nomadcoin/db"
	"github.com/kadragon/nomadcoin/utils"
)

type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height     int    `json:"height"`
}

var b *blockchain
var once sync.Once // 무조건 한번만 실행하는것을 보장하기 위해서 추가

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)

	b.NewestHash = block.Hash
	b.Height = block.Height

	b.persist()
}

func (b *blockchain) Blocks() []*Block {
	var blocks []*Block

	hashCursor := b.NewestHash
	for {
		block, err := FindBlock(hashCursor)
		utils.HandleErr(err)
		blocks = append(blocks, block)

		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}

	return blocks
}

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}

			// search for checkpoint on the db
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				b.AddBlock("Genesis")
			} else {
				// restore b form bytes
				b.restore(checkpoint)
			}
		})
	}

	fmt.Println(b.NewestHash)

	return b
}
