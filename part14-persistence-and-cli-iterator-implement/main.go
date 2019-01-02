package main

import (
	"fmt"
	"math/big"
	"publicchain/part14-persistence-and-cli/BLC"
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
	blockchainIterator = blockchain.Iterator()
	var hashInt big.Int

	for {

		fmt.Printf("%x\n", blockchainIterator.CurrentHash)
		//获取下一个迭代器
		blockchainIterator = blockchainIterator.Next()
		//将迭代器中的hash转化为int大数，存储到hashInt
		hashInt.SetBytes(blockchainIterator.CurrentHash)

		if hashInt.Cmp(big.NewInt(0)) == 0 {
			break

		}

	}

}
