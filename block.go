package ethrpc

import (
	"math/big"
)

// BlockHeader 区块头
type BlockHeader struct {
	Hash          string `json:"hash"`              //区块头
	Confirmations uint64 `json:"confirmations"`     //确认数量
	Merkleroot    string `json:"merkleroot"`        //Merkler 根
	PrevBlockHash string `json:"previousblockhash"` //上一个区块hash
	Height        uint64 `json:"height" storm:"id"` //区块高度
	Version       uint64 `json:"version"`           //版本
	Time          uint64 `json:"time"`              // 时间
	Fork          bool   `json:"fork"`              // 分叉
	Symbol        string `json:"symbol"`            //币种简称
}


// Block - block object
// 区块链对象
type Block struct {
	Number           int
	Hash             string
	ParentHash       string
	Nonce            string
	Sha3Uncles       string
	LogsBloom        string
	TransactionsRoot string
	StateRoot        string
	Miner            string
	Difficulty       big.Int
	TotalDifficulty  big.Int
	ExtraData        string
	Size             int
	GasLimit         int
	GasUsed          int
	Timestamp        int
	Uncles           []string
	Transactions     []Transaction
}

//代理区块链
type proxyBlock interface {
	toBlock() Block
}
