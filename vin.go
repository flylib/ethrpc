package ethrpc

import (
	"encoding/json"
	"math/big"
)

//from Address - Address the transaction is send from.
//data Data - Compiled code of a contract OR the hash of the invoked method signature and encoded parameters.
//to Address - Address the transaction is send to.
//gas Integer - Gas provided for the transaction execution. It will return unused gas.
//gas_price Integer - Value of the gas for this transaction.
//value Integer - Value sent with the transaction.
//nonce Integer - Value of the nonce. This allows to overwrite your own pending transactions that use the same nonce.
//Gas:      600000,                  //600000  default:21000
//GasPrice: big.NewInt(20000000000), //big.NewInt(4500000000) 最快到账 60000000000 普通：20000000000   default:1000000000
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
