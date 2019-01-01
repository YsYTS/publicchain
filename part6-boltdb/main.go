package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

//数据库名字
const dbFile = "blockchain.db"

//仓库————表的名字
const blocksBucket = "blocks"

func main() {

	//--------------数据库创建-----------------
	//如果数据库存在，打开；如果不存在，创建一个数据库
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	//插入、更新数据库
	db.Update(func(tx *bolt.Tx) error {
		//判断表是否存在于数据库中
		b := tx.Bucket([]byte(blocksBucket))
		if b == nil {
			fmt.Println("No existing blockchain found. Creating a new on database ")
			//CreateBucket 创建表
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Panic(err)
			}
			// key []byte, value []byte
			//存储数据
			err = b.Put([]byte("wangkun"), []byte("wangkunno1"))
			if err != nil {
				log.Panic(err)
			}
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
}
