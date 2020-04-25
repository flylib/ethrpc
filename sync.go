package ethrpc

import (
	"encoding/json"
	"unsafe"
)

// Syncing - object with syncing data info
// 区块同步结构体
type Syncing struct {
	IsSyncing     bool //同步状态 false  表示同步完成
	StartingBlock int  //开始同步区块
	CurrentBlock  int  //当前区块高度
	HighestBlock  int  //区块高度
}

// UnmarshalJSON implements the json.Unmarshaler interface.
//JSON数据实现
func (s *Syncing) UnmarshalJSON(data []byte) error {
	proxy := new(proxySyncing)
	if err := json.Unmarshal(data, proxy); err != nil {
		return err
	}

	proxy.IsSyncing = true
	*s = *(*Syncing)(unsafe.Pointer(proxy))

	return nil
}



//代理同步
type proxySyncing struct {
	IsSyncing     bool   `json:"-"`
	StartingBlock hexInt `json:"startingBlock"`
	CurrentBlock  hexInt `json:"currentBlock"`
	HighestBlock  hexInt `json:"highestBlock"`
}

//代理事务
type proxyTransaction struct {
	Hash             string  `json:"hash"`
	Nonce            hexInt  `json:"nonce"`
	BlockHash        string  `json:"blockHash"`
	BlockNumber      *hexInt `json:"blockNumber"`
	TransactionIndex *hexInt `json:"transactionIndex"`
	From             string  `json:"from"`
	To               string  `json:"to"`
	Value            hexBig  `json:"value"`
	Gas              hexInt  `json:"gas"`
	GasPrice         hexBig  `json:"gasPrice"`
	Input            string  `json:"input"`
}
