package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type block struct {
	Data     string
	Hash     string
	PrevHash string
}

type blockchain struct {
	blocks []*block
}

var b *blockchain
var once sync.Once // 무조건 한번만 실행하는것을 보장하기 위해서 추가

// calculateHash calculate hash with block data, prevHash
func (b *block) calculateHash() {
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
func createBlock(data string) *block {
	newblock := block{data, "", getLastHash()}
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

func (b *blockchain) Allblocks() []*block {
	return b.blocks
}
