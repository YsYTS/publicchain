package main

import (
	"publicchain/part24-transactions-newutxotransaction/BLC"
)

func main() {

	//创建区块链
	blockchain := BLC.NewBlockchain()

	//创建Cli对象
	cli := BLC.Cli{blockchain}

	//调用Cli的Run方法
	cli.Run()
}
