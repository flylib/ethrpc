package ethrpc

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

const (
	GAS_LIMIT = 21000
	GAS_PRICE = 500000000000
)

//admin.nodeInfo
//{
//enode: "enode://e8766c279de0a4fc05a185d3afd1ed45ca5ad62a43c566c60b7909fbc4dcc47c8c1ca945c53eef55dc9adde27a7f41894ccca9a5a66fdc25ff5dbb781d77d874@127.0.0.1:36601",
//enr: "enr:-Jq4QFcSSDMJP7GZ6JZzMmmBLxwanuhqlHlsKx63h3bJGydXctK0roYSQf-O2SSA1FeRM2lGckv0WatLgyyIBBBgGbggg2V0aMrJhLWzns6DPIblgmlkgnY0gmlwhH8AAAGJc2VjcDI1NmsxoQLodmwnneCk_AWhhdOv0e1FylrWKkPFZsYLeQn7xNzEfIN0Y3CCjvmDdWRwgo75",
//id: "d20858ef353474ec813d9d4b5d3a8ec307c853a43e87442f5d4e71c6db9db9c7",
//ip: "127.0.0.1",
//listenAddr: "[::]:36601",
//name: "GBChian/v1.0.0-stable/linux-amd64/go1.14.3",
//ports: {
//discovery: 36601,
//listener: 36601
//},
//protocols: {
//cdt: {
//config: {
//chainId: 1, 对应找个chainId
//scrypt: {},
//singularityBlock: 3966693
//},
//consensus: "scrypt",
//difficulty: 9896789062,
//genesis: "0xe3fa5726a2f675b18df6e4d9cb092f186d3f0642bd4bc3eb7becbdd2dacf16d0",
//head: "0xeac6240b22590e30d8b9eccb0079b72c3f2e1cd7bf244b59927fc13c75cf60d4",
//network: 1
//}
//}
//}

//私钥签名转账
//@chainID  对应上面的chainId,一般默认为1

func SignTransaction(chainID int64, tx *types.Transaction, privateKeyStr string) (string, error) {
	privateKey, err := Private2Key(privateKeyStr)
	if err != nil {
		return "", err
	}
	signTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(chainID)), privateKey)
	if err != nil {
		return "", nil
	}
	b, err := rlp.EncodeToBytes(signTx)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
