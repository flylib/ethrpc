package ethrpc

import (
	"encoding/json"
	"math/big"
)

// T - input transaction object
//输入交易事务结构体
type T struct {
	From     string   `json:"from"`
	To       string   `json:"to"`
	Gas      int      `json:"gas"`
	GasPrice *big.Int `json:"gas_price"`
	Value    *big.Int `json:"value"`
	Data     string   `json:"data"`
	Nonce    int      `json:"nonce"`
}

// MarshalJSON implements the json.Unmarshaler interface.
//实现了交易的 json数据接口。
func (t T) MarshalJSON() ([]byte, error) {
	params := map[string]interface{}{
		"from": t.From,
	}
	if t.To != "" {
		params["to"] = t.To
	}
	if t.Gas > 0 {
		params["gas"] = IntToHex(t.Gas)
	}
	if t.GasPrice != nil {
		params["gasPrice"] = BigToHex(*t.GasPrice)
	}
	if t.Value != nil {
		params["value"] = BigToHex(*t.Value)
	}
	if t.Data != "" {
		params["data"] = t.Data
	}
	if t.Nonce > 0 {
		params["nonce"] = IntToHex(t.Nonce)
	}

	return json.Marshal(params)
}
