package ethrpc

import "fmt"

/**
Web3ClientVersion returns the current client version.
返回当前客户端版本号
参数：
none
返回：
String - 当前客户端版本号
*/
func (rpc *EthRPC) Web3ClientVersion() (string, error) {
	var clientVersion string
	err := rpc.request("web3_clientVersion", &clientVersion)
	return clientVersion, err
}

/**
	Web3Sha3 returns Keccak-256 (not the standardized SHA3-256) of the given data.
	返回指定数据的Keccak-256（不是标准化的SHA3-256）哈希值。
	参数：
	String - 将数据转化为SHA3哈希
	params: [
	  '0x68656c6c6f20776f726c64'
    ]
	返回：
	DATA - 返回SHA3
*/
func (rpc *EthRPC) Web3Sha3(data []byte) (string, error) {
	var hash string

	err := rpc.request("web3_sha3", &hash, fmt.Sprintf("0x%x", data))
	return hash, err
}
