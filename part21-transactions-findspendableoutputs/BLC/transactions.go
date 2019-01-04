package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

const subsidy = 10

type Transaction struct {
	//1.交易ID
	ID []byte
	//2.交易输入
	Vin []TXInput
	//3.交易输出
	Vout []TXOutput
}

//1.判断当前交易是否是CoinBaseTx
func (tx *Transaction) IsCoinbase() bool {


	return len(tx.Vin) == 1 && tx.Vin[0].Vout == -1 && len(tx.Vin[0].Txid) == 0
}

//创建一个新的 coinbase 交易
func NewCoinbaseTx(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}
	//创建特殊的输入
	txin := TXInput{[]byte{}, -1, data}
	//创建输出
	txout := TXOutput{subsidy, to}
	//创建交易
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	tx.SetID()

	return &tx
}

//设置交易Hash————先序列化，后hash
func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	//将序列化以后的字节数组生成256hash
	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

//交易输入
type TXInput struct {
	//1.交易的ID
	Txid []byte
	//2.存储TXOutput在Vout里面的索引
	Vout int
	//3.用户名
	Scriptsig string
}

//检查账号地址
func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
	return in.Scriptsig == unlockingData
}

//交易输出
type TXOutput struct {
	Value        int    //最小单位
	ScriptPubKey string //公钥————用户名
}




//检查是否能够解锁账号
func (out *TXOutput) CanBeUnlockedWith(unlockingData string) bool {
	return out.ScriptPubKey == unlockingData
}
