package ethrpc

import (
	"math/big"
)

//* QUANTITY|TAG
//The following options are possible for the defaultBlock parameter:
//
//HEX String - an integer block number
//String "earliest" for the earliest/genesis block
//String "latest" - for the latest mined block
//String "pending" - for the pending state/transactions
const (
	BlockTag_Earliest = "earliest" //创世纪块
	BlockTag_Latest   = "latest"   //最新块
	BlockTag_Pending  = "pending"  //待定
)

// BlockHeader 区块头
type BlockHeader struct {
	Hash          string `json:"hash"`              //区块头
	Confirmations uint64 `json:"confirmations"`     //确认数量
	MerkleRoot    string `json:"merkleroot"`        //Merkler 根
	PrevBlockHash string `json:"previousblockhash"` //上一个区块hash
	Height        uint64 `json:"height" storm:"id"` //区块高度
	Version       uint64 `json:"version"`           //版本
	Time          uint64 `json:"time"`              // 时间
	Fork          bool   `json:"fork"`              // 分叉
	Symbol        string `json:"symbol"`            //币种简称
}

// Block - block object
/*"number": "0x1b4", // 436
"hash": "0xe670ec64341771606e55d6b4ca35a1a6b75ee3d5145a99d05921026d1527331",
"parentHash": "0x9646252be9520f6e71339a8df9c55e4d7619deeb018d2a3f2d21fc165dde5eb5",
"nonce": "0xe04d296d2460cfb8472af2c5fd05b5a214109c25688d3704aed5484f9a7792f2",
"sha3Uncles": "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
"logsBloom": "0xe670ec64341771606e55d6b4ca35a1a6b75ee3d5145a99d05921026d1527331",
"transactionsRoot": "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
"stateRoot": "0xd5855eb08b3387c0af375e9cdb6acfc05eb8f519e419b874b6ff2ffda7ed1dff",
"miner": "0x4e65fda2159562a496f9f3522f89122a3088497a",
"difficulty": "0x027f07", // 163591
"totalDifficulty":  "0x027f07", // 163591
"extraData": "0x0000000000000000000000000000000000000000000000000000000000000000",
"size":  "0x027f07", // 163591
"gasLimit": "0x9f759", // 653145
"gasUsed": "0x9f759", // 653145
"timestamp": "0x54e34e8e" // 1424182926
"transactions": [{...},{ ... }]
"uncles": ["0x1606e5...", "0xd5145a9..."]*/
// 区块链对象
type Block struct {
	Number           int           //QUANTITY - 块编号，挂起块为null
	Hash             string        // DATA, 32 Bytes-块哈希，挂起块为null
	ParentHash       string        //  DATA, 32 Bytes-父哈希
	Nonce            string        //DATA, 8 Bytes - 生成的pow哈希，挂起块为null（随机值）
	Sha3Uncles       string        //DATA, 32 Bytes - 块中叔伯数据的SHA3哈希
	LogsBloom        string        //DATA, 256 Bytes-快日志的bloom过滤器，挂起块为null
	TransactionsRoot string        //DATA, 32 Bytes - 块中的交易树根节点
	StateRoot        string        //DATA, 32 Bytes - 块最终状态树的根节点
	ReceiptsRoot     string        //DATA, 32 Bytes - 块交易收据树的根节点
	Miner            string        //DATA, 20 Bytes - 挖矿奖励的接收账户
	Difficulty       big.Int       // QUANTITY - 块难度，整数
	TotalDifficulty  big.Int       //QUANTITY - 截止到本块的链上总难度
	ExtraData        string        // DATA - 块额外数据
	Size             int           //QUANTITY - 本块字节数
	GasLimit         int           //QUANTITY -本块允许的最大gas用量
	GasUsed          int           //QUANTITY -本块中所有交易使用的总gas用量
	Timestamp        int           //QUANTITY - 块时间戳
	Uncles           []string      //Array - 叔伯哈希数组
	Transactions     []Transaction //Array - 交易对象数组，或32字节长的交易哈希数组
}

//代理区块链
type proxyBlock interface {
	toBlock() Block
}
