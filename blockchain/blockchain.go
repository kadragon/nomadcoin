package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"sync"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
}

type blockchain struct {
	blocks []*Block
}

var b *blockchain
var once sync.Once // 무조건 한번만 실행하는것을 보장하기 위해서 추가

// calculateHash calculate hash with block data, prevHash
func (b *Block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

// getLastHash block의 길이를 측정해서, block이 없을 경우 빈 hash를
// block이 있을 경우, 마지막 block의 hash를 반환
func getLastHash() string {
	totalBlocks := len(GetBlockchain().blocks)
	if totalBlocks == 0 {
		return ""
	}

	return GetBlockchain().blocks[totalBlocks-1].Hash
}

// createBlock data를 입력 받아, 새로운 block을 만든다.
// 만드는 과정에서 prevHash를 가져오며, hash를 계산한다.
func createBlock(data string) *Block {
	newblock := Block{data, "", getLastHash(), len(GetBlockchain().Allblocks()) + 1}
	newblock.calculateHash()

	return &newblock
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis")
		})
	}

	return b
}

func (b *blockchain) Allblocks() []*Block {
	return b.blocks
}

var ErrNotFound = errors.New("block not found")

func (b *blockchain) GetBlock(height int) (*Block, error) {
	if height > len(b.blocks) {
		return nil, ErrNotFound
	}

	return b.blocks[height-1], nil
}
