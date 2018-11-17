package blockchain

// Block contains a hash of its data, along with data, and a hash to the
// previous block TODO: Add full header
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// Blockchain is a chain of Blocks
type Blockchain struct {
	Blocks []*Block
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
