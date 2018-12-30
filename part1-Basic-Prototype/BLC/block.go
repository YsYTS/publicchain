package BLC

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	//时间戳，创建区块时的时间
	Timestamp int64
	//上一个区块的hash——父hash
	PrevBlockHash []byte
	//data，交易数据
	Data []byte
	//hash，当前区块的hash
	Hash []byte
}

func (block *Block) SetHash() {

	//1.将时间戳转化为字节数组
	//a.将int64转字符串

	//第二个参数: 2~36,转化为2进制~36进制
	timeString := strconv.FormatInt(block.Timestamp, 2)
	//b.将字符串转字节数组
	timestamp := []byte(timeString)
	//2.将除了Hash以外的其他属性，以字节数组的形式全拼接起来

	headers := bytes.Join([][]byte{block.PrevBlockHash, block.Data, timestamp}, []byte{})

	//3.将拼接起来的数据进行256 hash
	hash := sha256.Sum256(headers)
	//4.将hash赋给Hash属性字节
	block.Hash = hash[:]

}

//创建新区块，工厂方法
func NewBlock(data string, prevBlockHash []byte) *Block {
	//创建区块，.Unix()将时间戳转化为int64类型
	block := &Block{Timestamp: time.Now().Unix(), PrevBlockHash: prevBlockHash, Data: []byte(data), Hash: []byte{}}
	//设置当前区块的Hash值
	block.SetHash()
	//返回区块
	return block
}

//创建创世区块，并返回创世区块
func NewGenesisBlock() *Block {
	return NewBlock("Genenis Block", []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

}
