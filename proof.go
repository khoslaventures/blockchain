package blockchain

// Take data from the block
// Create a counter (nonce) which starts at 0
// Create a hash of the data plus the counter
// Check the hash to see if it meets a set of requirements

// Requirements:
// The First few bytes must contain 0s (hashcash)

// Difficulty represents fixed level of hardness unlike actual blockchain
// by requiring a certain amount of leading zeroes..
const Difficulty = 12

// ProofofWork is a proof of work that is used to submit a valid block
type ProofofWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(b *Block) *ProofofWork {
	target := big.NewInt(1)
	// Get leading number of zeros
	target.Lsh(target, uint(256-Difficulty))

	pow := &ProofofWork{b, target}
	return pow
}

func (pow *ProofOfWork) InitData(nonce int) []byte {
	// TODO: Uniform bytes of the block data
}
