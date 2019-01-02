package main

import (
	"fmt"
	"publicchain/part13-persistence-and-cli/BLC"
)

func main() {
	blockchain := BLC.NewBlockchain()
	fmt.Println(blockchain)

	fmt.Printf("tip: %x\n", blockchain.Tip)
	fmt.Println(blockchain.DB)
	fmt.Printf("\n")
	blockchain.AddBlock("send 20 BTC to xiaomin from wangkun")
	fmt.Printf("tip: %x\n", blockchain.Tip)
	fmt.Printf("\n")
	blockchain.AddBlock("send 33 BTC to xiaowang from wangkun")
	fmt.Printf("tip: %x\n", blockchain.Tip)
	fmt.Printf("\n")
	blockchain.AddBlock("send 40 BTC to xiaoxiao from kunkun")
	fmt.Printf("tip: %x\n", blockchain.Tip)
	fmt.Printf("\n")

	var blockchainIterator *BLC.BlockchainIterator
	for {
		blockchainIterator = blockchain.Iterator()
		fmt.Printf("%x\n", blockchainIterator.CurrentHash)

		blockchainIterator = blockchainIterator.Next()

	}

}
