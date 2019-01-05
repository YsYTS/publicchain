package BLC

import (
	"log"

	"github.com/boltdb/bolt"
)

//迭代器
type BlockchainIterator struct {
	CurrentHash []byte   //当前正在遍历的区块的Hash
	DB          *bolt.DB //数据库
}

//迭代器
func (blockchain *Blockchain) Iterator() *BlockchainIterator {

	return &BlockchainIterator{blockchain.Tip, blockchain.DB}
}

//下一个迭代
func (bi *BlockchainIterator) Next() *BlockchainIterator {

	var nextHash []byte
	//查询数据
	err := bi.DB.View(func(tx *bolt.Tx) error {
		//获取表
		b := tx.Bucket([]byte(blocksBucket))

		//通过当前hash获取block
		currentBlockBytes := b.Get(bi.CurrentHash)

		//反序列化
		currentBlock := DeserializeBlock(currentBlockBytes)

		nextHash = currentBlock.PrevBlockHash

		return nil
	})

	if err != nil {
		log.Panic(err)
	}
	return &BlockchainIterator{nextHash, bi.DB}
}
