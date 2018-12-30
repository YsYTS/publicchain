package BLC

type Block struct {
	//时间戳，创建区块时的时间
	Timestamp int64
	//上一个区块的hash——父hash
	PreBlockHash []byte
	//data，交易数据
	Data []byte
	//hash，当前区块的hash
	Hash []byte
}
