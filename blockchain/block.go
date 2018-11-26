package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

// Block contains a hash of its data, along with data, and a hash to the
// previous block TODO: Add full header
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// CreateBlock bundles data and gets a BlockHash and returns the Block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Mine()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// BeginGenesis creates the first block in the blockchain
func BeginGenesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// Serialize outputs a byte representation of Block for the DB
func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)
	Handle(err)

	return res.Bytes()
}

// Deserialize takes in a serialized byte slice and returns a block
func Deserialize(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)

	Handle(err)

	return &block
}

func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
