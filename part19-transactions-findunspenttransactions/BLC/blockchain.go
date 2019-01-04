package BLC

import (
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/boltdb/bolt"
)

//数据库名字
const dbFile = "blockchain.db"

//仓库————表的名字
const blocksBucket = "blocks"

//创世区块中存储的信息
const genesisCoinbaseData = "The Time 03/Jan/2009 Chancellor on brink of second bailout for banks"

type Blockchain struct {
	Tip []byte   //区块链里最后一个区块的hash
	DB  *bolt.DB //数据库
}

// 先找到包含当前用户未花费输出的所有交易集合
//返回交易的数组

func (bc *Blockchain) FindUnspentTransactions(address string) []Transaction {
	blockchainIterator := bc.Iterator()

	var hashInt big.Int

	for {
		err := blockchainIterator.DB.View(func(tx *bolt.Tx) error {
			//获取表
			b := tx.Bucket([]byte(blocksBucket))
			//通过Hash获取区块字节数组
			blockBytes := b.Get(blockchainIterator.CurrentHash)
			//反序列化
			block := DeserializeBlock(blockBytes)

			fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
			fmt.Printf("Timestamp: %s \n", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
			fmt.Printf("Hash: %x \n", block.Hash)
			fmt.Printf("Nonce: %d\n", block.Nonce)

			for _, transaction := range block.Transactions {
				fmt.Printf("TransactionHash: %x\n", transaction.ID)

			}
			fmt.Println()
			return nil

		})
		if err != nil {
			log.Panic(err)
		}

		//获取下一个迭代器
		blockchainIterator = blockchainIterator.Next()

		//将迭代器中的Hash存储到hashInt中
		hashInt.SetBytes(blockchainIterator.CurrentHash)

		// 		Cmp compares x and y and returns:

		// -1 if x < y 0 if x == y +1 if x > y

		if hashInt.Cmp(big.NewInt(0)) == 0 {
			break

		}

	}

	return nil

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
			//创建创世区块的交易对象————coinbase
			cbtx := NewCoinbaseTx("wangkun", genesisCoinbaseData)
			//将创世区块
			genesisBlock := NewGenesisBlock(cbtx)
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
