package main

import (
	"fmt"
	"log"
	"math/big"
	"publicchain/part15-persistence-and-cli/BLC"
	"time"

	"github.com/boltdb/bolt"
)

const bolcksBucket = "blocks"

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
		err := blockchainIterator.DB.View(func(tx *bolt.Tx) error {
			//获取表
			b := tx.Bucket([]byte(bolcksBucket))
			//通过hash获取区块字节数组
			blockBytes := b.Get(blockchainIterator.CurrentHash)
			//反序列化
			block := BLC.DeserializeBlock(blockBytes)

			fmt.Printf("Data: %s\n", string(block.Data))
			fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
			fmt.Printf("Timestamp: %s\n", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
			fmt.Printf("Hash: %x\n", block.Hash)
			fmt.Printf("Nonce: %d\n", block.Nonce)

			fmt.Println()

			return nil
		})
		if err != nil {
			log.Panic(err)
		}

		//获取下一个迭代器
		blockchainIterator = blockchainIterator.Next()
		//将迭代器中的hash转化为int大数，存储到hashInt
		hashInt.SetBytes(blockchainIterator.CurrentHash)

		if hashInt.Cmp(big.NewInt(0)) == 0 {
			break

		}

	}

}
