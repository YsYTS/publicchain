package main

import (
	"publicchain/part29-newpart-transaction-send/BLC"
)

func main() {

	//创建Cli对象
	cli := BLC.CLI{}

	//调用Cli的Run方法 
	cli.Run()
}
