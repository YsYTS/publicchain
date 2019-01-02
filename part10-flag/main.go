package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

//Go 提供了一个 flag 包，支持基本的命令行标志解析。

func main() {
	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Blockdata")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		fmt.Println("No addblock and printchain")
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		fmt.Println("Data:" + *addBlockData)
	}

	if printChainCmd.Parsed() {
		fmt.Println("printchain, printchain, printchain")
	}

}
