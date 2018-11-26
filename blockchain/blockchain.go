package blockchain

import (
	"fmt"

	"github.com/dgraph-io/badger"
)

const (
	dbPath = "/Users/akashkhosla/blockdata"
)

// Blockchain is a chain of Blocks
type Blockchain struct {
	LastHash []byte
	Database *badger.DB
}

// BlockchainIterator is an iterator for our persistent structure on disk
type BlockchainIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

// InitBlockchain starts the chain at genesis
func InitBlockchain() *Blockchain {
	var lastHash []byte

	opts := badger.DefaultOptions

	// Put keys and values in the same folder
	opts.Dir = dbPath
	opts.ValueDir = dbPath

	db, err := badger.Open(opts)
	Handle(err)

	// pass in a closure with a transaction
	err = db.Update(func(txn *badger.Txn) error {
		// lh is the last hash key
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			fmt.Println("No existing blockchain found")
			genesis := BeginGenesis()
			fmt.Println("Genesis proved")

			err = txn.Set(genesis.Hash, genesis.Serialize())
			Handle(err)
			err = txn.Set([]byte("lh"), genesis.Hash)

			lastHash = genesis.Hash

			return err
		} else {
			// DB already has a blockchain
			item, err := txn.Get([]byte("lh"))
			Handle(err)
			lastHash, err = item.Value()
			return err
		}
	})

	Handle(err)

	blockchain := Blockchain{lastHash, db}
	return &blockchain
}

// AddBlock appends a new block to the end of the Blockchain
func (chain *Blockchain) AddBlock(data string) {
	var lastHash []byte

	// Get the last blockhash
	err := chain.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		lastHash, err = item.Value()

		return err
	})
	// nb is a new block
	nb := CreateBlock(data, lastHash)
	err = chain.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(nb.Hash, nb.Serialize())
		Handle(err)
		err = txn.Set([]byte("lh"), nb.Hash)

		chain.LastHash = nb.Hash

		return err
	})
	Handle(err)
}

// Iterator creates an iterator for the DB, since we need a persistent way to
// iterate through the DB.
func (chain *Blockchain) Iterator() *BlockchainIterator {
	iter := &BlockchainIterator{chain.LastHash, chain.Database}

	return iter
}

// Next will get the next block in the DB (going backwards, via prevHash)
func (iter *BlockchainIterator) Next() *Block {
	var block *Block

	err := iter.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)
		encodedBlock, err := item.Value()
		block = Deserialize(encodedBlock)

		return err
	})
	Handle(err)

	iter.CurrentHash = block.PrevHash

	return block
}
