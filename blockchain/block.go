package blockchain


type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// CreateBlock initializes a new block using the provided data and the previous block's hash.
// It calculates the block's hash and nonce using the Proof of Work system,
// and returns the newly created block.
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// AddBlock adds a new block to the BlockChain, given some data.
// It uses the last block in the BlockChain as the previous block,
// and appends the newly created block to the BlockChain.
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

// Genesis returns the Genesis block, which is the first block in any BlockChain.
// The Genesis block has a data of "Genesis", and a previous hash of nil.
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain returns a new BlockChain containing the Genesis block.
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}