package BLC

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

//数据库名字
const dbFile = "blockchain.db"

//仓库————表的名字
const blocksBucket = "blocks"

type Blockchain struct {
	Tip []byte   //区块链里最后一个区块的hash
	DB  *bolt.DB //数据库
}

// //新增区块
// func (blockchain *Blockchain) AddBlock(data string) {
// 	//1.创建新的Block
// 	preBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
// 	newBlock := NewBlock(data, preBlock.Hash)

// 	//2.将区块添加到Blocks里面
// 	blockchain.Blocks = append(blockchain.Blocks, newBlock)

// }

//创建一个带有创世区块节点的区块链
func NewBlockchain() *Blockchain {
	var tip []byte //获取最后一个区块的hash

	//1.尝试打开或者创建数据库
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	//2.db.update更新数据
	//	1)表是否存在，如果不存在，需要创建表
	//	2）创建创世区块
	//	3)需要将创世去区块序列化
	//	4）将创世区块的Hash作为key，Block的序列化数据作为value存储的表中
	//	5）设置一个key，l，将hash作为value再次存储到数据里面

	err = db.Update(func(tx *bolt.Tx) error {
		//判断这一张表是否存在于数据库中
		b := tx.Bucket([]byte(blocksBucket))

		//表不存在
		if b == nil {
			fmt.Println("No existing blockchain found. Creating a new one...")
			//创建创世区块
			genesisBlock := NewGenesisBlock()
			//创建表
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Panic(err)
			}

			//将创世区块序列化以后的数据存储到表中
			err = b.Put(genesisBlock.Hash, genesisBlock.Serialize())

			if err != nil {
				log.Panic(err)
			}

			//存储 Hash
			err = b.Put([]byte("1"), genesisBlock.Hash)
			if err != nil {
				log.Panic(err)
			}

			tip = genesisBlock.Hash

		} else {
			//表存在
			//key:1
			//value:最后一个区块的hash
			tip = b.Get([]byte("1"))
		}
		return nil

	})

	if err != nil {
		log.Panic(err)
	}
	//defer db.Close()
	return &Blockchain{tip, db}

}
