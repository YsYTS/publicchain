package main

import (
	"fmt"
	"os"
)

func main() {
	//os.Args 提供原始命令行参数访问功能。
	argsWithProg := os.Args
	fmt.Println(argsWithProg)

	argsWithoutProg := os.Args[1:]
	fmt.Println(argsWithoutProg)

	arg := os.Args[3]
	fmt.Println(arg)

}
