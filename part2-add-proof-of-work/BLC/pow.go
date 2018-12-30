package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

var (
	//定义Nonce最大值
	maxNonce = math.MaxInt64
)

//难度控制
const targetBits = 16

//00000001 难度3，左移8 - 3 = 5，00100000

type ProofOfWork struct {
	block  *Block   //当前需要验证的区块
	target *big.Int //大数存储，区块难度
}

//数据拼接，返回字节数组
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

//ProofOfWork对象的方法
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)

		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])
		//00000010000
		//00000001000
		//  if x < y, return -1
		// if x == y, return 0
		// if x > y, return +1
		// if hashInt < pow.target  return -1
		//如果hashInt < pow.target, 返回
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}

	}
	fmt.Printf("\n\n")
	return nonce, hash[:]

}

//创建一个ProofOfWork的结构体对象,工厂方法
func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	//Lsh，左移
	target.Lsh(target, uint(256-targetBits))
	// fmt.Println("--------------------------")
	// fmt.Println(target)
	pow := &ProofOfWork{block, target}

	return pow
}

//验证当前工作量证明的有效性
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid

}
