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

// Take the data from the block

// create a counter (nonce) which starts at 0

// create a hash of the data plus the counter

// check the hash to see if it meets a set of requirements

// Requirements:
// The First few bytes must contain 0s

const Difficulty = 12

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

// NewProof takes a block and returns a ProofOfWork. It sets the target of the
// proof of work to a number that is very large, but still very slightly less
// than the maximum number that can be expressed in 256 bits. The target is
// calculated by shifting 1 to the left by 256-Difficulty bits.
func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))

	pow := &ProofOfWork{b, target}

	return pow
}

// InitData takes a nonce and returns a slice of bytes that is to be
// hashed in the proof of work. The slice of bytes is created by joining
// the previous hash of the block, the hash of the transactions of the
// block, the nonce, and the difficulty of the proof of work.
func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.HashTransactions(),
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)

	return data
}

// Run performs the proof of work by continuously hashing the data from the block
// plus a nonce until a hash is found that meets the requirements. It prints the
// hash on each iteration and returns the nonce and the hash when finished.
func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}

	}
	fmt.Println()

	return nonce, hash[:]
}

// Validate takes the nonce from the block and runs the proof of work
// calculation again. It checks if the hash from the calculation is
// less than the target from the proof of work. If it is, the proof of
// work is valid and the function returns true. Otherwise it returns
// false.
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}

// ToHex takes an int64 and returns a slice of bytes representing the number
// in Big Endian order. It will panic if there is an error in the conversion.
func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)

	}

	return buff.Bytes()
}