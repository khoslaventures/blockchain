# Blockchain in Golang

Some practice for getting better at Go. This is a persistent (BadgerDB)
blockchain implementation with Proof of Work, Merkle Trees and command line
interface. This leverages the new Go modules dependency management system.

## Usage

```bash
# Will print out blockchain, if no blocks, will create genesis
go run main.go print

# Add a block, with a string of BLOCK_DATA
go run main.go add -block BLOCK_DATA
```

### Example
```bash
go run main.go add -block "send 10 BTC to bob"
```

## TODOs

- [x] Proof of Work
- [x] Basic block structure
- [x] Begin command line interface
- [x] Add a persistence layer like BadgerDB or BoltDB
- [ ] Transactions
- [ ] Accounts/Private Key Management
- [ ] Basic Wallet CLI
- [ ] Digital Signatures
- [ ] UTXO Persistence
- [ ] Merkle Trees (fraud proofs, light client support)
- [ ] Networking Between Clients
- [ ] Dynamic difficulty (Optional)
- [ ] Light client (Optional)
