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
	defer db.Close()
	//查询数据
	err = db.View(func(tx *bolt.Tx) error {
		//获取表
		b := tx.Bucket([]byte(blocksBucket))
		valueByte := b.Get([]byte("wangkun"))

		fmt.Printf("%s", valueByte)

		return nil
	})

}
