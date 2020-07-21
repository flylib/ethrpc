package ethrpc

import (
	"encoding/json"
	"unsafe"
)

// Log - log object
//日志对象
type Log struct {
	Removed          bool
	LogIndex         int
	TransactionIndex int
	TransactionHash  string
	BlockNumber      int
	BlockHash        string
	Address          string
	Data             string
	Topics           []string
}


//代理日志
type proxyLog struct {
	Removed          bool     `json:"removed"`
	LogIndex         hexInt   `json:"logIndex"`
	TransactionIndex hexInt   `json:"transactionIndex"`
	TransactionHash  string   `json:"transactionHash"`
	BlockNumber      hexInt   `json:"blockNumber"`
	BlockHash        string   `json:"blockHash"`
	Address          string   `json:"address"`
	Data             string   `json:"data"`
	Topics           []string `json:"topics"`
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (log *Log) UnmarshalJSON(data []byte) error {
	proxy := new(proxyLog)
	if err := json.Unmarshal(data, proxy); err != nil {
		return err
	}

	*log = *(*Log)(unsafe.Pointer(proxy))

	return nil
}

//日志接口
type logger interface {
	Println(v ...interface{})
}




