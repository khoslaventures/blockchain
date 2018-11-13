package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// Block contains a hash of its data, along with data, and a hash to the
// previous block TODO: Add full header
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// Blockchain is a chain of blocks
type Blockchain struct {
	blocks []*Block
}

// GetBlockHash generates a SHA-256 hash over block header and data
func (b *Block) GetBlockHash() {
	// Join our data
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	// Generate (simplified) hash
	hash := sha256.Sum256(info)
	// Convert array to slice
	b.Hash = hash[:]
}

// CreateBlock bundles data and gets a BlockHash and returns the Block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.GetBlockHash()
	return block
}

// AddBlock appends a new block to the end of the Blockchain
func (chain *Blockchain) AddBlock(data string) {
	lastBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := CreateBlock(data, lastBlock.Hash)
	chain.blocks = append(chain.blocks, newBlock)
}

// BeginGenesis creates the first block in the blockchain
func BeginGenesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockchain starts the chain at genesis
func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{BeginGenesis()}}
}

func main() {
	chain := InitBlockchain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
