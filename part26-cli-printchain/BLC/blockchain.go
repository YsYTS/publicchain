package BLC

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"os"

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

func (bc *Blockchain) FindUnspentTransactions(address string, txs []*Transaction) []Transaction {

	//存储未花费输出的交易
	var unspentTXs []Transaction
	//存储所有区块交易的输入
	spentTXOs := make(map[string][]int)

	blockchainIterator := bc.Iterator()

	var hashInt big.Int

	for _, transaction := range txs {

		if transaction.IsCoinbase() == false {
			for _, in := range transaction.Vin {
				if in.CanUnlockOutputWith(address) {
					inTxID := hex.EncodeToString(in.Txid)
					spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Vout)

				}
			}
		}
	}

	for _, transaction := range txs {
		txID := hex.EncodeToString(transaction.ID)

	Outputs:
		for outIdx, out := range transaction.Vout {
			//是否已经被花费
			if spentTXOs[txID] != nil {
				for _, spentOut := range spentTXOs[txID] {
					if spentOut == outIdx {
						continue Outputs
					}
				}
			}

			if out.CanBeUnlockedWith(address) {
				unspentTXs = append(unspentTXs, *transaction)
			}
		}
	}

	for {
		err := blockchainIterator.DB.View(func(tx *bolt.Tx) error {
			//获取表
			b := tx.Bucket([]byte(blocksBucket))
			//通过Hash获取区块字节数组
			blockBytes := b.Get(blockchainIterator.CurrentHash)
			//反序列化
			block := DeserializeBlock(blockBytes)

			for _, transaction := range block.Transactions {
				fmt.Printf("TransactionHash: %x\n", transaction.ID)
				//将byte array 类型转为 string
				txID := hex.EncodeToString(transaction.ID)
			Outputs:
				for outIdx, out := range transaction.Vout {
					//是否已经被花费？
					if spentTXOs[txID] != nil {
						for _, spentOut := range spentTXOs[txID] {
							if spentOut == outIdx {
								continue Outputs
							}
						}
					}

					if out.CanBeUnlockedWith(address) {
						unspentTXs = append(unspentTXs, *transaction)
					}
					if transaction.IsCoinbase() == false {
						for _, in := range transaction.Vin {
							if in.CanUnlockOutputWith(address) {
								inTxID := hex.EncodeToString(in.Txid)
								spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Vout)
							}
						}
					}
				}

			}

			fmt.Println()
			return nil

		})

		if err != nil {
			log.Panic(err)
		}

		//获取下一个迭代器
		blockchainIterator = blockchainIterator.Next()

		//将迭代器中的hash存储到hashInt
		hashInt.SetBytes(blockchainIterator.CurrentHash)

		if hashInt.Cmp(big.NewInt(0)) == 0 {
			break
		}
	}
	return unspentTXs

}

//查找可用的未消费的输出信息
//16 10
func (bc *Blockchain) FindSpendableOutPuts(address string, amount int, txs []*Transaction) (int, map[string][]int) {
	//{"11111":[1,2,3], "00000":[2,3,5]}
	//字典，存储交易id， Vout中的未花费TXOutput的index
	unspentOutputs := make(map[string][]int)

	//查看未花费
	unspentTXs := bc.FindUnspentTransactions(address, txs)

	accumulated := 0 //统计unspentOutputs里面对应的TXOutput所对应的总量

Work:
	//遍历交易数组
	for _, tx := range unspentTXs {
		txID := hex.EncodeToString(tx.ID)
		//遍历交易中的Vout
		for outIdx, out := range tx.Vout {
			if out.CanBeUnlockedWith(address) && accumulated < amount {
				accumulated += out.Value
				unspentOutputs[txID] = append(unspentOutputs[txID], outIdx)

				if accumulated >= amount {
					break Work
				}
			}
		}
	}

	//12, {"11111":[1,2,3]}
	return accumulated, unspentOutputs
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

//根据交易数据，打包新的区块
func (bc *Blockchain) MineBlock(txs []*Transaction) {
	err := bc.DB.Update(func(tx *bolt.Tx) error {
		//新建区块
		newBlock := NewBlock(txs, bc.Tip)

		//将区块存储到数据库
		b := tx.Bucket([]byte(blocksBucket))

		if b != nil {
			//key: newBlock.Hash
			//value: newBlock.Serialize()
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}

			//key: []byte{"l"}
			//value: newBlock.Hash
			err = b.Put([]byte("l"), newBlock.Hash)
			if err != nil {
				log.Panic(err)
			}

			//更新blockchain最新区块的hash
			bc.Tip = newBlock.Hash

		}

		return nil

	})

	if err != nil {
		log.Panic(err)
	}
}

//判断数据库是否存在
func dbExists() bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}

	return true
}
