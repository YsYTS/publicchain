package main

import (
	"publicchain/part26-cli-printchain/BLC"
)

func main() {

	//创建Cli对象
	cli := BLC.Cli{}

	//调用Cli的Run方法
	cli.Run()
}
