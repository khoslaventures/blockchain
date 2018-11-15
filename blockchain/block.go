package blockchain

import (
	"bytes"
	"crypto/sha256"
)

// Block contains a hash of its data, along with data, and a hash to the
// previous block TODO: Add full header
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// Blockchain is a chain of Blocks
type Blockchain struct {
	Blocks []*Block
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
	lastBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, lastBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

// BeginGenesis creates the first block in the blockchain
func BeginGenesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockchain starts the chain at genesis
func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{BeginGenesis()}}
}
