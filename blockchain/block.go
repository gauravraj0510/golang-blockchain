package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

type Block struct {
	Hash         []byte
	Transactions []*Transaction
	PrevHash     []byte
	Nonce        int
}

// HashTransactions takes all the transaction hashes in a block, concatenates
// them, and hashes them into a single hash.
func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}

// CreateBlock constructs a new block using the provided transactions and previous hash.
// It initializes a ProofOfWork for the block, performs the proof-of-work to find a valid nonce and hash,
// and assigns these to the block before returning it.
func CreateBlock(txs []*Transaction, prevHash []byte) *Block {
	block := &Block{[]byte{}, txs, prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// Genesis creates and returns a Block with the given Transaction as its only
// Transaction. This block has no previous block (its PrevHash is empty) and
// represents the first block in the Blockchain.
func Genesis(coinbase *Transaction) *Block {
	return CreateBlock([]*Transaction{coinbase}, []byte{})
}

// Serialize converts the Block into a byte slice using gob encoding.
// It returns the encoded bytes which represent the serialized form of the block.
func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)

	Handle(err)

	return res.Bytes()
}

// Deserialize takes a byte slice representing a block and decodes it into
// a Block object. It will panic if there is an error in the decoding process.
func Deserialize(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)

	Handle(err)

	return &block
}

// Handle is a simple function for panicking the program if an error occurs.
// It's used throughout the codebase to handle unexpected errors.
func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}