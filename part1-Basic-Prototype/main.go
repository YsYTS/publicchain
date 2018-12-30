package main

import (
	"fmt"
	"publicchain/part1-Basic-Prototype/BLC"
	"time"
)

//hash:16机制————64个字符，32个字节，256位

func main() {
	blockchain := BLC.NewBlockChain()
	blockchain.AddBlock("send 20 btc to wangkun from xiaoming")
	blockchain.AddBlock("send 30 btc to alvin from xiaoming")
	blockchain.AddBlock("send 40 btc to kun from xiaoming")
	for _, block := range blockchain.Blocks {
		fmt.Printf("Data:%s\n", string(block.Data))
		fmt.Printf("PrevBlockHase:%x\n", block.PrevBlockHash)
		fmt.Printf("Timestamp:%s\n", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Printf("Hash:%x\n", block.Hash)

		fmt.Println()

	}
}
