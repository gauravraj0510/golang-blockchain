package main

import(
	"fmt"
	"strconv"
	"github.com/gauravraj0510/golang-blockchain/blockchain"
)

// main initializes a new BlockChain, adds a few blocks to it, and then prints
// each block's Previous Hash, Data, and Hash.
func main() {
	fmt.Printf("==========\n")
	chain := blockchain.InitBlockChain()
	chain.AddBlock("1st Block after Genesis!")
	chain.AddBlock("2nd Block after Genesis!")
	chain.AddBlock("3rd Block after Genesis!")
	fmt.Printf("==========\n")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println("----------------")
	}
}