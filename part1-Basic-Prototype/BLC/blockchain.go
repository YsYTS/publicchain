package BLC

type BlockChain struct {
	Blocks []*Block //存储有序的区块
}

//新增区块
func (blockchain *BlockChain) AddBlock(data string) {
	//1.创建新的区块
	preBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
	newBlock := NewBlock(data, preBlock.Hash)
	//2.将区块添加到Blocks里面
	blockchain.Blocks = append(blockchain.Blocks, newBlock)
}

//创建一个带有创世区块的区块链
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}
