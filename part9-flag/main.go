package main

import (
	"flag"
	"fmt"
)

//Go 提供了一个 flag 包，支持基本的命令行标志解析。

func main() {
	//基本的标记声明仅支持字符串、整数型和布尔值选项。这里声明一个默认值为“foo”的字符
	wordPtr := flag.String("word", "foo", "a string")
	//使用和声明 word 标志相同的方法来声明 numb 和 fork 标志。
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	//用程序中已有的参数来声明一个标志也是可以的。
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	//所有标志声明完后，调用 flag.Parse() 来执行命令解析。
	flag.Parse()
	//这里将仅输出解析的选项及后面的位置参数。
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())

	flag.Usage()

}
