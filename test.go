package ethrpc

//
//import (
//	"crypto/ecdsa"
//	"encoding/hex"
//	"github.com/ethereum/go-ethereum/accounts/keystore"
//	"github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/common/hexutil"
//	"github.com/ethereum/go-ethereum/core/types"
//	"github.com/ethereum/go-ethereum/crypto"
//	"github.com/ethereum/go-ethereum/rlp"
//	"math/big"
//)
//
//const (
//	GAS_LIMIT = 21000
//	GAS_PRICE = 500000000000
//)
//
//func SignTransaction(chainID int64, tx *types.Transaction, privateKeyStr string) (string, error) {
//	privateKey, err := StringToPrivateKey(privateKeyStr)
//	if err != nil {
//		return "", err
//	}
//	signTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(chainID)), privateKey)
//	if err != nil {
//		return "", nil
//	}
//	b, err := rlp.EncodeToBytes(signTx)
//	if err != nil {
//		return "", err
//	}
//	return hex.EncodeToString(b), nil
//}
//
////离线转账
////@chainID int64 网络ID,1-主网ID
//func OfflineTransfer(chainID int64, nonce uint64, toAddress string, value *big.Int, privateKey string) (string, error) {
//	tx := types.NewTransaction(nonce, common.HexToAddress(toAddress), value, GAS_LIMIT, big.NewInt(GAS_PRICE), nil)
//	return SignTransaction(chainID, tx, privateKey)
//}
//
//
//

//私钥==>ecdsa.PrivateKey
//func StringToPrivateKey(privateKeyStr string) (*ecdsa.PrivateKey, error) {
//	privateKeyByte, err := hexutil.Decode(privateKeyStr)
//	if err != nil {
//		return nil, err
//	}
//	privateKey, err := crypto.ToECDSA(privateKeyByte)
//	if err != nil {
//		return nil, err
//	}
//	return privateKey, nil
//}
//
//keystore+password===>ecdsa.PrivateKey
//func KeystoreToPrivateKey(keystoreContent []byte, password string) (*ecdsa.PrivateKey, error) {
//	unlockedKey, err := keystore.DecryptKey(keystoreContent, password)
//	if err != nil {
//		return nil, err
//	}
//	return unlockedKey.PrivateKey, nil
//}
