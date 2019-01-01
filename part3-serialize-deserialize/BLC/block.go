package BLC

import (
	"bytes"
	"encoding/gob"
	"log"
)

type Block struct {
	//Data，交易数据
	Data []byte
	//Nonce 随机数
	Nonce int
}

//将Block对象（结构体）序列化为[]byte
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)

	}
	return result.Bytes()
}

//将字节数组反序列化为Block（结构体）
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}
