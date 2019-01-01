package main

import (
	"fmt"
	"publicchain/part3-serialize-deserialize/BLC"
)

func main() {
	block := BLC.Block{[]byte("send 34 BTC to xiaoming"), 1000}

	fmt.Printf("%s\n", block.Data)
	fmt.Printf("%d\n", block.Nonce)

	bytes := block.Serialize()
	fmt.Println(bytes)

	blc := BLC.DeserializeBlock(bytes)
	fmt.Printf("%s\n", blc.Data)
	fmt.Printf("%d\n", blc.Nonce)

}
