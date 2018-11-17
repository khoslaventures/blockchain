package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

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

// NewProof creates a new ProofofWork struct to be used for Mine
func NewProof(b *Block) *ProofofWork {
	target := big.NewInt(1)
	// Get leading number of zeros
	target.Lsh(target, uint(256-Difficulty))

	pow := &ProofofWork{b, target}
	return pow
}

// InitData gets uniform bytes of the block data
func (pow *ProofofWork) InitData(nonce int) []byte {
	data := bytes.Join([][]byte{
		pow.Block.PrevHash,
		pow.Block.Data,
		ToHex(int64(nonce)),
		ToHex(int64(Difficulty)),
	}, []byte{})
	return data
}

// Mine runs a standard proof of work algorithm
func (pow *ProofofWork) Mine() (int, []byte) {
	var intHash big.Int
	var hash [32]byte
	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			// Hash less than the target, block found
			break
		} else {
			// Continue searching for valid block
			nonce++
		}
	}
	fmt.Println()

	return nonce, hash[:]
}

func (pow *ProofofWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)
	hash := sha256.Sum256(data)

	intHash.SetBytes(hash[:])
	return intHash.Cmp(pow.Target) == -1
}

// ToHex converts an int64 to a slice of bytes
func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
