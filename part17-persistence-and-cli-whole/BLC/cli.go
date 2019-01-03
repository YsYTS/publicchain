package BLC

import (
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/boltdb/bolt"
)

type Cli struct {
	BC *Blockchain
}

//打印参数信息
func (cli *Cli) printUsage() {
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("\taddblock -data BLOCK_DATA - add a block to the blockchain")
	fmt.Println("\tprintchain - print all the blocks of the blockchain")
}

//判断终端参数的个数
func (cli *Cli) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *Cli) printChain() {
	var blockchainIterator *BlockchainIterator

	blockchainIterator = cli.BC.Iterator()
	var hashInt big.Int

	for {

		err := blockchainIterator.DB.View(func(tx *bolt.Tx) error {
			//获取表
			b := tx.Bucket([]byte(blocksBucket))
			//通过hash获取区块字节数组
			blockBytes := b.Get(blockchainIterator.CurrentHash)
			//反序列化
			block := DeserializeBlock(blockBytes)

			fmt.Printf("Data: %s\n", string(block.Data))
			fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
			fmt.Printf("Timestamp: %s\n", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
			fmt.Printf("Hash: %x\n", block.Hash)
			fmt.Printf("Nonce: %d\n", block.Nonce)

			fmt.Println()

			return nil
		})
		if err != nil {
			log.Panic(err)
		}

		//获取下一个迭代器
		blockchainIterator = blockchainIterator.Next()
		//将迭代器中的hash转化为int大数，存储到hashInt
		hashInt.SetBytes(blockchainIterator.CurrentHash)

		if hashInt.Cmp(big.NewInt(0)) == 0 {
			break

		}

	}
}

func (cli *Cli) addBlock(data string) {
	cli.BC.AddBlock(data)
}

func (cli *Cli) Run() {
	//判断终端参数的个数，如果没有参数，直接打印Usage信息并退出
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	addBlockData := addBlockCmd.String("data", "", "Block data")

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
		cli.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			cli.printUsage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)

	}

	if printChainCmd.Parsed() {
		//通过迭代器遍历所有区块信息
		cli.printChain()

	}

}
