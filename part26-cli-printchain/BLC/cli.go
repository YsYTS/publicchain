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

	fmt.Println("\tgetbalance -address ADDRESS - Get balance of ADDRESS")
	fmt.Println("\tcreateblockchain -address ADDRESS -Create a blockchain and send genesis block reward to ADDRESS")
	fmt.Println("\tprintchian - Print all the blocks of the blockchain")
	fmt.Println("\tsend -from FROM -to TO -amount AMOUNT - Send AMOUNT of coins from FROM address to TO")

}

//判断终端参数的个数
func (cli *Cli) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *Cli) printChain() {

	//判断数据库是否存在
	if dbExists() == false {
		fmt.Println()
		fmt.Println("Database is not exist.")
		cli.printUsage()
		return
	}

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

			fmt.Printf("Transactions: %x\n", block.HashTransactions())
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

func (cli *Cli) sendToken() {

	// 1. 10 -> wangkun
	// 2. 3 -> 转给tushen
	// (1)新建一个交易

	var inputs []*Transaction

	tx1 := NewUTXOTransaction("wangkun", "tushen", 3, cli.BC, inputs)
	fmt.Println("第一笔交易")
	fmt.Println(tx1)

	inputs = append(inputs, tx1)

	tx2 := NewUTXOTransaction("wangkun", "alvin", 2, cli.BC, inputs)
	fmt.Println("第二笔交易")
	fmt.Println(tx2)
	inputs = append(inputs, tx2)

	tx3 := NewUTXOTransaction("tushen", "alvin", 1, cli.BC, inputs)
	fmt.Println("第三笔交易")
	fmt.Println(tx3)

	tx4 := NewUTXOTransaction("wangkun", "akun", 3, cli.BC, inputs)
	fmt.Println("第四笔交易")
	fmt.Println(tx4)

	fmt.Println("第二笔交易")
	fmt.Println(tx2)
	cli.BC.MineBlock([]*Transaction{tx1, tx2, tx3, tx4})

}

func (cli *Cli) addBlock(data string) {

	// count, outputMap := cli.BC.FindSpendableOutPuts("wangkun", 5)
	// fmt.Println(count)
	// fmt.Println(outputMap)

	// cli.sendToken()
	// tx := cli.BC.FindUnspentTransactions("wangkun")
	// fmt.Println("print wangkun:")
	// fmt.Println(tx)

	// tx = cli.BC.FindUnspentTransactions("tushen")
	// fmt.Println("pirnt tushen:")
	// fmt.Println(tx)

	// count, outputMap := cli.BC.FindSpendableOutPuts("wangkun", 5)
	// fmt.Println(count)
	// fmt.Println(outputMap)
	cli.sendToken()

}

//Run function————终端运行
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
