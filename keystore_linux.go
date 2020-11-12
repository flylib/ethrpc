package ethrpc

import (
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

//私钥==>ecdsa.PrivateKey
func Private2Key(privateKeyStr string) (*ecdsa.PrivateKey, error) {
	privateKeyByte, err := hexutil.Decode(privateKeyStr)
	if err != nil {
		return nil, err
	}
	privateKey, err := crypto.ToECDSA(privateKeyByte)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

//keystore+password===>ecdsa.PrivateKey
func KeystoreToPrivateKey(keystoreContent []byte, password string) (*ecdsa.PrivateKey, error) {
	key, err := keystore.DecryptKey(keystoreContent, password)
	if err != nil {
		return nil, err
	}
	return key.PrivateKey, nil
}

//keystore+password===>ecdsa.PrivateKey
func KeystoreToPrivateKeyString(keystoreContent []byte, password string) (string, error) {
	key, err := keystore.DecryptKey(keystoreContent, password)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(crypto.FromECDSA(key.PrivateKey)), nil
}

//keystore+password===>ecdsa.PrivateKey
func ParseKeystore(keystoreContent []byte, password string) (addr, priKey string, err error) {
	key, err := keystore.DecryptKey(keystoreContent, password)
	if err != nil {
		return "", "", err
	}
	addr = key.Address.Hex()
	priKey = hex.EncodeToString(crypto.FromECDSA(key.PrivateKey))
	return
}
