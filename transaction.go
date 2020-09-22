package ethrpc

import (
	"encoding/json"
	"math/big"
	"unsafe"
)

//合约交易码

const (
	TTtransferCode     = "0xa9059cbb" //转账签名编码
	TTbalanceOfCode    = "0x70a08231" //余额查询编码
	TTDecimalsCode     = "0x313ce567" //位數
	TTallowanceCode    = "0xdd62ed3e" //合约拥有者
	TTsymbolCode       = "0x95d89b41" //合约简称
	TTtotalSupplyCode  = "0x18160ddd" //发行总量
	TTnameCode         = "0x06fdde03" //合约名称
	TTapproveCode      = "0x095ea7b3" //认证
	TTtransferFromCode = "0x23b872dd" //交易
)

//交易状态
const (
	TStatusSuccess = "0x1" //成功
	TStatusFair    = "0x0" //失败
)

// Transaction - transaction object
/*"hash":"0xc6ef2fc5426d6ad6fd9e2a26abeab0aa2411b7ab17f30a99d3cb96aed1d1055b",
"nonce":"0x",
"blockHash": "0xbeab0aa2411b7ab17f30a99d3cb9c6ef2fc5426d6ad6fd9e2a26a6aed1d1055b",
"blockNumber": "0x15df", // 5599
"transactionIndex":  "0x1", // 1
"from":"0x407d73d8a49eeb85d32cf465507dd71d507100c1",
"to":"0x85h43d8a49eeb85d32cf465507dd71d507100c1",
"value":"0x7f110", // 520464
"gas": "0x7f110", // 520464
"gasPrice":"0x09184e72a000",
"input":"0x603880600c6000396000f300603880600c6000396000f3603880600c6000396000f360",*/
//交易结构体
type Transaction struct {
	Hash             string  //DATA, 32字节 - 交易哈希
	Nonce            int     //QUANTITY - 本次交易之前发送方已经生成的交易数量
	BlockHash        string  // DATA, 32字节 - 交易所在块的哈希，对于挂起块，该值为null
	BlockNumber      *int    // QUANTITY - 交易所在块的编号，对于挂起块，该值为null
	TransactionIndex *int    //QUANTITY - 交易在块中的索引位置，挂起块该值为null
	From             string  // DATA, 20字节 - 交易发送方地址
	To               string  // DATA, 20字节 - 交易接收方地址，对于合约创建交易，该值为null
	Value            big.Int // QUANTITY - 发送的以太数量，单位：wei
	Gas              int     // QUANTITY - 发送方提供的gas可用量
	GasPrice         big.Int //: QUANTITY - 发送方提供的gas价格，单位：wei
	Data             string  //
	Input            string  //DATA - 随交易发送的数据
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *Transaction) UnmarshalJSON(data []byte) error {
	proxy := new(proxyTransaction)
	if err := json.Unmarshal(data, proxy); err != nil {
		return err
	}

	*t = *(*Transaction)(unsafe.Pointer(proxy))

	return nil
}

type proxyBlockWithTransactions struct {
	Number           hexInt             `json:"number"`
	Hash             string             `json:"hash"`
	ParentHash       string             `json:"parentHash"`
	Nonce            string             `json:"nonce"`
	Sha3Uncles       string             `json:"sha3Uncles"`
	LogsBloom        string             `json:"logsBloom"`
	TransactionsRoot string             `json:"transactionsRoot"`
	StateRoot        string             `json:"stateRoot"`
	Miner            string             `json:"miner"`
	Difficulty       hexBig             `json:"difficulty"`
	TotalDifficulty  hexBig             `json:"totalDifficulty"`
	ExtraData        string             `json:"extraData"`
	Size             hexInt             `json:"size"`
	GasLimit         hexInt             `json:"gasLimit"`
	GasUsed          hexInt             `json:"gasUsed"`
	Timestamp        hexInt             `json:"timestamp"`
	Uncles           []string           `json:"uncles"`
	Transactions     []proxyTransaction `json:"transactions"`
}

func (proxy *proxyBlockWithTransactions) toBlock() Block {
	return *(*Block)(unsafe.Pointer(proxy))
}

//没有事务的代理块
type proxyBlockWithoutTransactions struct {
	Number           hexInt   `json:"number"`
	Hash             string   `json:"hash"`
	ParentHash       string   `json:"parentHash"`
	Nonce            string   `json:"nonce"`
	Sha3Uncles       string   `json:"sha3Uncles"`
	LogsBloom        string   `json:"logsBloom"`
	TransactionsRoot string   `json:"transactionsRoot"`
	StateRoot        string   `json:"stateRoot"`
	Miner            string   `json:"miner"`
	Difficulty       hexBig   `json:"difficulty"`
	TotalDifficulty  hexBig   `json:"totalDifficulty"`
	ExtraData        string   `json:"extraData"`
	Size             hexInt   `json:"size"`
	GasLimit         hexInt   `json:"gasLimit"`
	GasUsed          hexInt   `json:"gasUsed"`
	Timestamp        hexInt   `json:"timestamp"`
	Uncles           []string `json:"uncles"`
	Transactions     []string `json:"transactions"`
}

func (proxy *proxyBlockWithoutTransactions) toBlock() Block {
	block := Block{
		Number:           int(proxy.Number),
		Hash:             proxy.Hash,
		ParentHash:       proxy.ParentHash,
		Nonce:            proxy.Nonce,
		Sha3Uncles:       proxy.Sha3Uncles,
		LogsBloom:        proxy.LogsBloom,
		TransactionsRoot: proxy.TransactionsRoot,
		StateRoot:        proxy.StateRoot,
		Miner:            proxy.Miner,
		Difficulty:       big.Int(proxy.Difficulty),
		TotalDifficulty:  big.Int(proxy.TotalDifficulty),
		ExtraData:        proxy.ExtraData,
		Size:             int(proxy.Size),
		GasLimit:         int(proxy.GasLimit),
		GasUsed:          int(proxy.GasUsed),
		Timestamp:        int(proxy.Timestamp),
		Uncles:           proxy.Uncles,
	}

	block.Transactions = make([]Transaction, len(proxy.Transactions))
	for i := range proxy.Transactions {
		block.Transactions[i] = Transaction{
			Hash: proxy.Transactions[i],
		}
	}

	return block
}

//transactionHash: DATA, 32 Bytes - 交易的hash.
//transactionIndex: QUANTITY - 区块中交易index的位置。
//blockHash: DATA, 32 Bytes - 此交易所处的区块hash。
//blockNumber: QUANTITY - 此交易所处的区块号
//cumulativeGasUsed: QUANTITY - 当这笔交易已经在区块中执行完成，所使用的gas总量。
//gasUsed: QUANTITY - 此特定交易所使用的单个gas金额。
//contractAddress: DATA, 20 Bytes - 创建的合同地址（如果该交易是创建合约，* 否则为空。
//logs: Array - 此交易生成的日志对象数组。
// 交易收据 TransactionReceipt - transaction receipt object
type TransactionReceipt struct {
	TransactionHash   string
	TransactionIndex  int
	BlockHash         string
	BlockNumber       int
	CumulativeGasUsed int
	GasUsed           int
	ContractAddress   string
	Logs              []Log
	LogsBloom         string
	Root              string
	Status            string
}

//代理交易收据
type proxyTransactionReceipt struct {
	TransactionHash   string `json:"transactionHash"`
	TransactionIndex  hexInt `json:"transactionIndex"`
	BlockHash         string `json:"blockHash"`
	BlockNumber       hexInt `json:"blockNumber"`
	CumulativeGasUsed hexInt `json:"cumulativeGasUsed"`
	GasUsed           hexInt `json:"gasUsed"`
	ContractAddress   string `json:"contractAddress,omitempty"`
	Logs              []Log  `json:"logs"`
	LogsBloom         string `json:"logsBloom"`
	Root              string `json:"root"`
	Status            string `json:"status,omitempty"`
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *TransactionReceipt) UnmarshalJSON(data []byte) error {
	proxy := new(proxyTransactionReceipt)
	if err := json.Unmarshal(data, proxy); err != nil {
		return err
	}

	*t = *(*TransactionReceipt)(unsafe.Pointer(proxy))

	return nil
}
