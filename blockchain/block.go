package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/kadragon/nomadcoin/db"
	"github.com/kadragon/nomadcoin/utils"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

func createBlock(data string, prevHash string, height int) *Block {
	block := &Block{
		Data:     data,
		Hash:     "",
		PrevHash: prevHash,
		Height:   height,
	}

	payload := fmt.Sprintf("%s%s%d", block.Data, block.PrevHash, block.Height)
	block.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))

	block.persist()

	return block
}

var ErrNotFound = errors.New("block not found")

func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}

func FindBlock(hash string) (*Block, error) {
	blockByte := db.Block(hash)

	if blockByte == nil {
		return nil, ErrNotFound
	}

	block := &Block{}
	block.restore(blockByte)

	return block, nil
}
