package main

import (
	"publicchain/part28-newpart-transaction/BLC"
)

func main() {

	//创建Cli对象
	cli := BLC.CLI{}

	//调用Cli的Run方法
	cli.Run()
}
