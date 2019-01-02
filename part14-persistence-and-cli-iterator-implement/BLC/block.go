package BLC

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"time"
)

type Block struct {
	//时间戳，创建区块的时间
	Timestamp int64
	//上一个区块的hash，父hash
	PrevBlockHash []byte
	//交易数据
	Data []byte
	//当前区块Hash
	Hash []byte
	//Nonce值
	Nonce int
}

//将Block对象序列化为[]byte
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

//工厂方法
func NewBlock(data string, PrevBlockHash []byte) *Block {
	//创建区块
	block := &Block{Timestamp: time.Now().Unix(), PrevBlockHash: PrevBlockHash, Data: []byte(data), Hash: []byte{}, Nonce: 0}
	//将block作为参数，创建一个pow对象
	pow := NewProofOfWork(block)
	//Run()执行一次工作量证明
	nonce, hash := pow.Run()
	//设置区块Hash
	block.Hash = hash
	//设置区块Nonce
	block.Nonce = nonce

	isValid := pow.Validate()

	fmt.Println(isValid)

	//返回区块
	return block

}

//将字节数组反序列化成Block
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}

//创建创世区块，并返回创世区块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}
