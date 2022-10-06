package main

import (
	"crypto/sha256"
	"fmt"
)

type blockinfo struct {
	Trx      string
	hashval  string
	hashval2 string
}
type blockchain struct {
	gens_block blockinfo
	blockchain []blockinfo
}

func createbc() blockchain {
	gens_block := blockinfo{
		Trx: "Transaction",
	}
	gens_block.hashval = gens_block.CalculateHash()
	return blockchain{
		gens_block,
		[]blockinfo{gens_block},
	}
}
func (block *blockchain) NewBlock(trxstr1 string, nonceval int) {
	tempvar := fmt.Sprint(trxstr1, nonceval)
	oldbc := block.blockchain[len(block.blockchain)-1]
	createdbc := blockinfo{
		Trx:      tempvar,
		hashval2: oldbc.CalculateHash(),
	}
	block.blockchain = append(block.blockchain, createdbc)
	block.blockchain[(len(block.blockchain) - 1)].hashval = block.blockchain[(len(block.blockchain) - 1)].CalculateHash()
}
func (block *blockchain) ListBlocks() string {
	for ctr := 0; ctr < len(block.blockchain); ctr++ {
		if ctr == 0 {
			fmt.Println("\n")
			fmt.Println("Genesis will b block 0 & data in it is:", block.blockchain[ctr])
			fmt.Println("\n")
		}
		if ctr > 0 {
			fmt.Println("\nThe blockchain \n\n", ctr)
			fmt.Println("Transaction (TRX) & rand val ", block.blockchain[ctr].Trx)
			fmt.Println("Hash of prev block: \n", block.blockchain[ctr].hashval2)
			fmt.Println("Hash of block: \n", block.blockchain[ctr].hashval)
		}
	}
	return ""
}
func (block *blockchain) ChangeBlock(bcID int, trxstr string) {
	block.blockchain[bcID].Trx = trxstr
	block.blockchain[bcID].hashval = block.blockchain[bcID].CalculateHash()
}
func (block blockinfo) CalculateHash() string {
	tempvar := block.Trx + block.hashval2
	hashval := sha256.Sum256([]byte(tempvar))
	return fmt.Sprintf("%x", hashval)
}
func (block *blockchain) VerifyChain() bool {
	for ctr := 1; ctr < len(block.blockchain); ctr++ {
		cblock := block.blockchain[ctr]
		prev := block.blockchain[ctr-1]
		if cblock.hashval != cblock.CalculateHash() || cblock.hashval2 != prev.hashval {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Blockchain Assignment 1 - Hassan Sohail 19i 1785\n")
	block := createbc()
	block.NewBlock("Hassan to Bill", 2)
	block.NewBlock("Robot to Dog", 3)
	fmt.Println(block.VerifyChain())
	block.ChangeBlock(1, "WhoAreU to bill")
	fmt.Println(block.VerifyChain())
	block.ListBlocks()
}
