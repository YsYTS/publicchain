package BLC

type Blockchain struct {
	Blocks []*Block //存储有序区块
}

//新增区块
func (blockchain *Blockchain) AddBlock(data string) {
	//1.创建新的Block
	preBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
	newBlock := NewBlock(data, preBlock.Hash)

	//2.将区块添加到Blocks里面
	blockchain.Blocks = append(blockchain.Blocks, newBlock)

}

//创建一个带有创世区块节点的区块链
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
