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

//新增区块
func (blockchain *Blockchain) AddBlock(data string) {

	//1.创建区块

	newBlock := NewBlock(data, blockchain.Tip)

	//2.Update数据

	err := blockchain.DB.Update(func(tx *bolt.Tx) error {
		//获取数据库表
		b := tx.Bucket([]byte(blocksBucket))
		//存储新区块数据
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		//更新l对应的hash
		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}

		//将最新区块hash存储到blockchain的tip中
		blockchain.Tip = newBlock.Hash

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

}

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
			err = b.Put([]byte("l"), genesisBlock.Hash)
			if err != nil {
				log.Panic(err)
			}

			tip = genesisBlock.Hash

		} else {
			//表存在
			//key:1
			//value:最后一个区块的hash
			tip = b.Get([]byte("l"))
		}
		return nil

	})

	if err != nil {
		log.Panic(err)
	}
	//defer db.Close()
	return &Blockchain{tip, db}

}
