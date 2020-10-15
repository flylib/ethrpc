package ethrpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"regexp"
)

/**
 *@Project     openchain
 *@Author      king
 *@CreateTime  2020/4/18 1:41 下午
 *@ClassName   api
 *@Description  ethereum jsonRPC  API 接口
 */

const (
	v1 = "1.0"
	V2 = "2.0"
)

//var _ EthereumAPI = (*EthRPC)(nil)

//http客户端接口
type httpClient interface {
	Post(url string, contentType string, body io.Reader) (*http.Response, error)
}

const (
	AuthNone      = iota //none
	AuthBasicAuth        //基本认证
	AuthMD5              //MD5加密
	AuthToken            //token
)

// EthRPC - Ethereum rpc client
//以太坊RPC客户端结构体
type EthRPC struct {
	url       string     //URL链接
	client    httpClient //客户端对象
	log       logger     //日志
	Debug     bool       //调试
	Version   string     //版本
	Auth      int        //认证方式
	BasicAuth Account
}

//基本认证
type Account struct {
	User string `json:"user"`
	PWD  string `json:"pwd"`
}

// NewEthRPC create new rpc client with given url
func NewEthRPC(url string, options ...func(rpc *EthRPC)) *EthRPC {
	rpc := &EthRPC{
		url:     url,
		client:  http.DefaultClient,
		log:     log.New(os.Stderr, "", log.LstdFlags),
		Version: V2,
	}
	for _, option := range options {
		option(rpc)
	}
	return rpc
}

func (rpc *EthRPC) call(method string, target interface{}, params ...interface{}) error {
	result, err := rpc.Call(method, params...)
	if err != nil {
		return err
	}

	if target == nil {
		return nil
	}

	return json.Unmarshal(result, target)
}

// URL returns client url
func (rpc *EthRPC) URL() string {
	return rpc.url
}

//设置授权方式
func SetAuthType(auth int) func(rpc *EthRPC) {
	return func(rpc *EthRPC) {
		rpc.Auth = auth
	}
}

//设置授权方式
func SetBasicAuth(act Account) func(rpc *EthRPC) {
	return func(rpc *EthRPC) {
		rpc.BasicAuth = act
	}
}

//设置授权方式
func Debug() func(rpc *EthRPC) {
	return func(rpc *EthRPC) {
		rpc.Debug = true
	}
}

// 设置认证方式
func (rpc *EthRPC) SetAuthType(authType int) {
	rpc.Auth = authType
}

// Call returns raw response of method call
func (rpc *EthRPC) Call(method string, params ...interface{}) (json.RawMessage, error) {
	req := Request{
		ID:      1,
		JsonRPC: rpc.Version,
		Method:  method,
		Params:  params,
	}
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	newRequest, err := http.NewRequest("POST", rpc.url, bytes.NewBuffer(reqData))
	if err != nil {
		return nil, err
	}
	//基本认证
	if rpc.Auth == AuthBasicAuth {
		newRequest.SetBasicAuth(rpc.BasicAuth.User, rpc.BasicAuth.PWD)
	}
	newRequest.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(newRequest)
	if err != nil {
		return nil, err
	}
	//res, err := rpc.client.Post(rpc.url, "application/json", bytes.NewBuffer(reqData))
	//if err != nil {
	//	return nil, err
	//}
	defer res.Body.Close()
	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if rpc.Debug {
		rpc.log.Println(fmt.Sprintf("%s\nRequest: %s\nResponse: %s\n", method, reqData, resData))
	}
	resETHObject := new(Response)
	if err := json.Unmarshal(resData, resETHObject); err != nil {
		return nil, err
	}

	if resETHObject.Error != nil {
		return nil, *resETHObject.Error
	}

	return resETHObject.Result, nil
}

// RawCall returns raw response of method call (Deprecated) (弃用)
func (rpc *EthRPC) RawCall(method string, params ...interface{}) (json.RawMessage, error) {
	return rpc.Call(method, params...)
}

/**
EthProtocolVersion returns the current ethrpc protocol version.
返回当前以太坊协议的版本
参数：
none
返回：
String - The current ethrpc protocol version
*/
func (rpc *EthRPC) EthProtocolVersion() (string, error) {
	var protocolVersion string

	err := rpc.call("eth_protocolVersion", &protocolVersion)
	return protocolVersion, err
}

/**
EthSyncing returns an object with data about the sync status or false.
返回同步状态,或返回FALSE。
参数：
none
返回：
Object|Boolean, 同步状态数据或FALSE的对象在不同步时：

startingBlock: QUANTITY - 导入开始的块（只有在同步达到它的头之后才会被重置）
currentBlock: QUANTITY - 当前块,与eth_blockNumber相同
highestBlock: QUANTITY - 估算的最高区块
*/
func (rpc *EthRPC) EthSyncing() (*Syncing, error) {
	result, err := rpc.RawCall("eth_syncing")
	if err != nil {
		return nil, err
	}
	syncing := new(Syncing)
	if bytes.Equal(result, []byte("false")) {
		return syncing, nil
	}
	err = json.Unmarshal(result, syncing)
	return syncing, err
}

/**
EthCoinbase returns the client coinbase address
返回客户端的coinbase地址
参数：
none
返回：
DATA, 20 bytes - 当前coinbase的地址.
*/
func (rpc *EthRPC) EthCoinbase() (string, error) {
	var address string

	err := rpc.call("eth_coinbase", &address)
	return address, err
}

/**
EthMining returns true if client is actively mining new blocks.
如果客户端正在主动挖掘新块，则返回true。
参数：
none
返回：
Boolean - 客户端主动的挖矿返回true，否则为false。
*/
func (rpc *EthRPC) EthMining() (bool, error) {
	var mining bool
	err := rpc.call("eth_mining", &mining)
	return mining, err
}

/**
EthHashrate returns the number of hashes per second that the node is mining with.
返回节点正在挖掘的每秒散列数。
参数：
none
返回：
QUANTITY - 每秒的哈希数。
*/
func (rpc *EthRPC) EthHashrate() (int, error) {
	var response string

	if err := rpc.call("eth_hashrate", &response); err != nil {
		return 0, err
	}

	return ParseInt(response)
}

/**
EthGasPrice returns the current price per gas in wei.
每个gas的当前价格,单位wei。
参数：
none
返回：
QUANTITY - 当前gas的价格整数。
*/
func (rpc *EthRPC) EthGasPrice() (big.Int, error) {
	var response string
	if err := rpc.call("eth_gasPrice", &response); err != nil {
		return big.Int{}, err
	}

	return ParseBigInt(response)
}

/**
EthAccounts returns a list of addresses owned by client.
返回客户端拥有的地址列表。
参数：
none
返回：
Array of DATA, 20 Bytes - 该地址拥有的地址列表
*/
func (rpc *EthRPC) EthAccounts() ([]string, error) {
	accounts := []string{}
	err := rpc.call("eth_accounts", &accounts)
	return accounts, err
}

/**
EthBlockNumber returns the number of most recent block.
返回最近块的数量。
参数：
none
返回：
QUANTITY - 客户端所在当前块号的整数。
*/
func (rpc *EthRPC) EthBlockNumber() (int, error) {
	var response string
	if err := rpc.call("eth_blockNumber", &response); err != nil {
		return 0, err
	}
	return ParseInt(response)
}

/**
	EthGetBalance returns the balance of the account of given address in wei.
	返回给定地址中的账户余额单位是wei。
	参数
	DATA, 20 Bytes - 通过地址来检索余额。

	QUANTITY|TAG - 区块号, 或使用 "latest", "earliest"，"pending"

	params: [
   '0x407d73d8a49eeb85d32cf465507dd71d507100c1',
   'latest']
	返回
	QUANTITY - 当前余额的整数，以wei(Ether = 10e18Wei)为单位。
*/
func (rpc *EthRPC) EthGetBalance(address, block string) (big.Int, error) {
	var response string
	if err := rpc.call("eth_getBalance", &response, address, block); err != nil {
		return big.Int{}, err
	}

	return ParseBigInt(response)
}

//tx pool content 交易池查询
func (rpc *EthRPC) EthTxContent() (content interface{}, err error) {
	var response interface{}
	if err := rpc.call("txpool_inspect", &response); err != nil {
		return "查询失败", err
	}

	return response, err
}

/**
	EthGetStorageAt returns the value from a storage position at a given address.
	返回指定地址的存储位置的值。
	参数
	DATA, 20 Bytes - 存储地址.
	QUANTITY - 存储位置的整数.
	QUANTITY|TAG - 区块号, 或 "latest", "earliest", "pending"

	params: [
 	  '0x407d73d8a49eeb85d32cf465507dd71d507100c1',
      '0x0', // storage position at 0
  	  '0x2' // state at block number 2]
	返回
	DATA - 在这个存储位置的值。
*/
func (rpc *EthRPC) EthGetStorageAt(data string, position int, tag string) (string, error) {
	var result string

	err := rpc.call("eth_getStorageAt", &result, data, IntToHex(position), tag)
	return result, err
}

/**
EthGetTransactionCount returns the number of transactions sent from an address.
返回指定地址发生的交易数量。
参数：
DATA, 20 Bytes - 地址.
QUANTITY|TAG - 区块号, 或"latest", "earliest" , "pending"

params: [
   '0x407d73d8a49eeb85d32cf465507dd71d507100c1',
   'latest' // state at the latest block]
返回
QUANTITY - 返回从该地址发送的交易数量。
*/
func (rpc *EthRPC) EthGetTransactionCount(address, block string) (int, error) {
	var response string

	if err := rpc.call("eth_getTransactionCount", &response, address, block); err != nil {
		return 0, err
	}

	return ParseInt(response)
}

/**
EthGetBlockTransactionCountByHash returns the number of transactions in a block from a block matching the given block hash.
从匹配给定块哈希的块中返回一个块中的事务数。
通过区块hash匹配区块中的交易数。
参数
DATA, 32 Bytes - 区块的hash
params: [ '0xb903239f8543d04b5dc1ba6579132b143087c68db1b2168786408fcbce568238' ]
返回
QUANTITY - 在该区块的交易数
*/
func (rpc *EthRPC) EthGetBlockTransactionCountByHash(hash string) (int, error) {
	var response string

	if err := rpc.call("eth_getBlockTransactionCountByHash", &response, hash); err != nil {
		return 0, err
	}

	return ParseInt(response)
}

/**
	EthGetBlockTransactionCountByNumber returns the number of transactions in a block from a block matching the given block
	通过区块号进行匹配，返回交易数。

	参数
	QUANTITY|TAG - 区块号, 或 "earliest", "latest" or "pending"

	params: [
   		'0xe8', // 232
	]
	返回
	QUANTITY - 在该区块中的交易数量。
*/
func (rpc *EthRPC) EthGetBlockTransactionCountByNumber(number int) (int, error) {
	var response string

	if err := rpc.call("eth_getBlockTransactionCountByNumber", &response, IntToHex(number)); err != nil {
		return 0, err
	}

	return ParseInt(response)
}

/**
	EthGetUncleCountByBlockHash returns the number of uncles in a block from a block matching the given block hash.
	通过指定的区块hash，返回uncle数量。
	参数
	DATA, 32 Bytes - 区块hash
	params: [
   '0xb903239f8543d04b5dc1ba6579132b143087c68db1b2168786408fcbce568238'
	]
	返回
	QUANTITY - 该区块的uncle数量
*/
func (rpc *EthRPC) EthGetUncleCountByBlockHash(hash string) (int, error) {
	var response string

	if err := rpc.call("eth_getUncleCountByBlockHash", &response, hash); err != nil {
		return 0, err
	}

	return ParseInt(response)
}

/**
返回具有指定哈希的块具有指定索引位置的叔伯。
参数
	DATA, 32字节 - 块哈希
	QUANTITY - 叔伯索引位置
	params: [
	   '0xc6ef2fc5426d6ad6fd9e2a26abeab0aa2411b7ab17f30a99d3cb96aed1d1055b',
	   '0x0' // 0
	]
返回
	响应结果请参考eth_getBlockByHash调用。
*/
func (rpc *EthRPC) EthGetUncleByBlockHashAndIndex(hash string, index int) (*Block, error) {
	return rpc.getBlock("eth_getUncleByBlockHashAndIndex", false, hash, IntToHex(index))
}

/**
返回具有指定哈希的块具有指定索引位置的叔伯。
参数
	QUANTITY|TAG - 整数块编号，或字符串"earliest"、"latest" 或"pending"
	QUANTITY - 叔伯在块内的索引序号
	params: [
	   '0x29c', // 668
	   '0x0' // 0
	]
返回
	响应结果请参考eth_getBlockByHash调用。
*/
func (rpc *EthRPC) EthGetUncleByBlockNumberAndIndex(number int, index int) (*Block, error) {
	return rpc.getBlock("eth_getUncleByBlockNumberAndIndex", false, IntToHex(number), IntToHex(index))
}

/**
	EthGetCode returns code at a given address.
	返回指定地址的code。
	参数
	DATA, 20 Bytes - 地址
	QUANTITY|TAG - 区块号, 或"latest", "earliest", "pending"
	params: [
   		'0xa94f5374fce5edbc8e2a8697c15331677e6ebf0b',
   		'0x2'  // 2
	]
	返回
	DATA - 指定地址的code。
*/
func (rpc *EthRPC) EthGetCode(address, block string) (string, error) {
	var code string

	err := rpc.call("eth_getCode", &code, address, block)
	return code, err
}

/**
EthSign signs data with a given address.
Calculates an Ethereum specific signature with: sign(keccak256("\x19Ethereum Signed Message:\n" + len(message) + message)))
用指定地址签名数据。
注意，签名地址必须解锁。
参数
DATA, 20 Bytes - 地址
DATA, 签名的数据
返回
DATA: 签名后的数据
*/
func (rpc *EthRPC) EthSign(address, data string) (string, error) {
	var signature string

	err := rpc.call("eth_sign", &signature, address, data)
	return signature, err
}

/**
使用web3.eth.signTransaction()方法对交易进行签名，用来签名的账户地址需要首先解锁。
参数:
	transactionObject：Object - 要签名的交易数据
	address：String - 用于签名的账户地址
	callback：Function - 可选的回调函数，其第一个参数为错误对象，第二个参数为结果
返回:
	一个Promise对象，其解析值为RLP编码的交易对象。该对象的raw属性可以用来通过web3.eth.sendSignedTransaction() 方法来发送交易。
*/
func (rpc *EthRPC) SignTransaction(t T) (SignPromise, error) {
	var res SignPromise
	err := rpc.call("eth_signTransaction", &res, t)
	return res, err
}

/**
	EthSendTransaction creates new message call transaction or a contract creation, if the data field contains code.
	如果数据字段包含code，则创建新的消息调用交易或创建合约。
	params
	Object - 交易对象
	from: DATA, 20 Bytes - 交易的发送地址。
	to: DATA, 20 Bytes -（创建新合约时可选）交易指向的地址。
	gas: QUANTITY -（可选，默认值：90000）为交易执行提供	的gas。它会返回未使用的gas。
	gasPrice: QUANTITY -（可选，默认：待确认）GasPrice就是你愿意为一个单位的Gas出多少Eth，一般用Gwei作单位。所以Gas Price 越高， 就表示交易中每运算一步，会支付更多的Eth。
	value: QUANTITY - (可选) 此交易发送的value。
	data: DATA - (可选) 合约的编译后的代码
	nonce: QUANTITY -（可选）一个整数。允许你使用相同的nonce覆盖自己处于等待中交易。

    example
	params: [{
  		"from": "0xb60e8dd61c5d32be8058bb8eb970870f07233155",
		"to": "0xd46e8dd67c5d32be8058bb8eb970870f072445675",
  		"gas": "0x76c0", // 30400,
  		"gasPrice": "0x9184e72a000", // 10000000000000
  		"value": "0x9184e72a", // 2441406250
  		"data": "0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675"
		}]

	返回：
	DATA, 32 Bytes - 交易hash或0hash(如果交易尚不可用)。
	当创建合约的时候，使用eth_getTransactionReceipt获取合约地址。
*/
func (rpc *EthRPC) EthSendTransaction(transaction T) (string, error) {
	var hash string
	err := rpc.call("eth_sendTransaction", &hash, transaction)
	return hash, err
}

func (rpc *EthRPC) SendTransaction(transaction Transaction) (string, error) {
	var hash string

	err := rpc.call("eth_sendTransaction", &hash, transaction)
	return hash, err
}

/**
	 EthSendRawTransaction creates new message call transaction or a contract creation for signed transactions.
	 为已签名的交易创建新的消息调用交易或合约创建。
	参数
		Object - 交易对象
		data: DATA, 已签名的交易数据。
		params: [{
  			"data": "0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675"
		}]
	返回：
		DATA, 32 Bytes - 交易hash，或0（如果交易还不可用）。

		当你创建的是一个合约时，使用eth_getTransactionReceipt来获取指定的合约地址。
*/
func (rpc *EthRPC) EthSendRawTransaction(data string) (string, error) {
	var hash string

	err := rpc.call("eth_sendRawTransaction", &hash, data)
	return hash, err
}

/**
	 EthSendRawTransaction creates new message call transaction or a contract creation for signed transactions.
	 为已签名的交易创建新的消息调用交易或合约创建。
	参数
		Object - 交易对象
		data: DATA, 已签名的交易数据。
		params: [{
  			"data": "0xd46e8dd67c5d32be8d46e8dd67c5d32be8058bb8eb970870f072445675058bb8eb970870f072445675"
		}]
	返回：
		DATA, 32 Bytes - 交易hash，或0（如果交易还不可用）。

		当你创建的是一个合约时，使用eth_getTransactionReceipt来获取指定的合约地址。
*/
func (rpc *EthRPC) EthSendSignedTransaction(data string) (string, error) {
	var hash string
	err := rpc.call("eth_sendSignedTransaction", &hash, data)
	return hash, err
}

/**
EthCall executes a new message call immediately without creating a transaction on the block chain.
立即执行新的消息调用，而不在区块链上创建交易。

参数：
	Object -交易调用的对象
	from: DATA, 20 Bytes - (可选) 交易发送的地址。
	to: DATA, 20 Bytes - 交易所针对的地址。
	gas: QUANTITY - (可选) 交易执行提供的gas。 eth_call消耗0gas，但某些执行可能需要此参数。
	gasPrice: QUANTITY - (可选) 为每个gas支付多少个gasPrice.
	value: QUANTITY - (可选) 此交易发送的value（整型）
	data: DATA - (可选) 合约的编译代码
	QUANTITY|TAG - 区块号, 或"latest", "earliest", "pending"
返回：
	DATA -已执行合约的返回价值.
*/
func (rpc *EthRPC) EthCall(transaction T, tag string) (string, error) {
	var data string

	err := rpc.call("eth_call", &data, transaction, tag)
	return data, err
}

/**
EthEstimateGas makes a call or transaction, which won't be added to the blockchain and returns the used gas, which can be used for estimating the used gas.
估算call或交易要使用的gas（这些call或交易不会添加到区块链中）。
参数
	所有eth_call的参数
返回
	QUANTITY - 使用gas的总金额。
*/
func (rpc *EthRPC) EthEstimateGas(transaction T) (int, error) {
	var response string

	err := rpc.call("eth_estimateGas", &response, transaction)
	if err != nil {
		return 0, err
	}

	return ParseInt(response)
}

/**
	 EthGetBlockByHash returns information about a block by hash.
	 通过hash，返回一个区块的信息。
	参数
		DATA, 32 Bytes - 区块hash。
		Boolean - 如果为true，则返回完整的交易对象，如果为false，则仅返回交易的散列。
		params: [
   			'0xe670ec64341771606e55d6b4ca35a1a6b75ee3d5145a99d05921026d1527331',
   			true
		]
	参考：http://cw.hubwiz.com/card/c/ethereum-json-rpc-api/1/3/21/
*/
func (rpc *EthRPC) EthGetBlockByHash(hash string, withTransactions bool) (*Block, error) {
	return rpc.getBlock("eth_getBlockByHash", withTransactions, hash, withTransactions)
}

/**
	EthGetBlockByNumber returns information about a block by block number.
	通过区块号返回有关块的信息。
	参数
	QUANTITY|TAG - 区块号, 或"earliest", "latest", "pending"
	Boolean - 如果为true，则返回完整的交易对象，如果为false，则仅返回事务的散列。
	params: [
  		 '0x1b4', // 436
 		  true
	]
	返回
	跟上面的eth_getBlockByHash一样
*/
func (rpc *EthRPC) EthGetBlockByNumber(number int, withTransactions bool) (*Block, error) {
	return rpc.getBlock("eth_getBlockByNumber", withTransactions, IntToHex(number), withTransactions)
}

/**
	EthGetTransactionByHash returns the information about a transaction requested by transaction hash.
	通过交易hash，返回交易信息。
	参数
		DATA, 32 Bytes - 交易hash
		params: [
  			 "0xb903239f8543d04b5dc1ba6579132b143087c68db1b2168786408fcbce568238"
		]
	参考：http://cw.hubwiz.com/card/c/ethereum-json-rpc-api/1/3/23/
*/
func (rpc *EthRPC) EthGetTransactionByHash(hash string) (*Transaction, error) {
	return rpc.getTransaction("eth_getTransactionByHash", hash)
}

func (rpc *EthRPC) getBlock(method string, withTransactions bool, params ...interface{}) (*Block, error) {
	result, err := rpc.RawCall(method, params...)
	if err != nil {
		return nil, err
	}
	if bytes.Equal(result, []byte("null")) {
		return nil, nil
	}

	var response proxyBlock
	if withTransactions {
		response = new(proxyBlockWithTransactions)
	} else {
		response = new(proxyBlockWithoutTransactions)
	}

	err = json.Unmarshal(result, response)
	if err != nil {
		return nil, err
	}

	block := response.toBlock()
	return &block, nil
}

func (rpc *EthRPC) getTransaction(method string, params ...interface{}) (*Transaction, error) {
	transaction := new(Transaction)

	err := rpc.call(method, transaction, params...)
	return transaction, err
}

/**
EthGetTransactionByBlockHashAndIndex returns information about a transaction by block hash and transaction index position.
通过区块hash和交易index位置，获取交易信息。
参数
	DATA, 32 Bytes - 区块的hash。
	QUANTITY - 交易index位置（整型）
返回
	跟eth_getBlockByHash一样。
参考：http://cw.hubwiz.com/card/c/ethereum-json-rpc-api/1/3/24/
*/
func (rpc *EthRPC) EthGetTransactionByBlockHashAndIndex(blockHash string, transactionIndex int) (*Transaction, error) {
	return rpc.getTransaction("eth_getTransactionByBlockHashAndIndex", blockHash, IntToHex(transactionIndex))
}

/**
	EthGetTransactionByBlockNumberAndIndex returns information about a transaction by block number and transaction index position.
	通过区块号和交易index位置，获取交易信息。
	参数
		QUANTITY|TAG - 区块号，或"earliest", "latest", "pending"
		QUANTITY - 交易的index位置。
		params: [
  		 '0x29c', // 668
  		 '0x0' // 0
		]
	返回
		跟eth_getBlockByHash一样。
*/
func (rpc *EthRPC) EthGetTransactionByBlockNumberAndIndex(blockNumber, transactionIndex int) (*Transaction, error) {
	return rpc.getTransaction("eth_getTransactionByBlockNumberAndIndex", IntToHex(blockNumber), IntToHex(transactionIndex))
}

/**
	EthGetTransactionReceipt returns the receipt of a transaction by transaction hash.
	Note That the receipt is not available for pending transactions.
	通过交易hash，接收交易结果。返回指定交易的收据，使用哈希指定交易。
    需要指出的是，挂起的交易其收据无效。
	注意，当交易处于pending时，接收不可用。
	参数
		DATA, 32 Bytes - 交易hash
		params: [
   			'0xb903239f8543d04b5dc1ba6579132b143087c68db1b2168786408fcbce568238'
		]

	返回:
		Object - 交易接收对象, 当接收没找到则为null:

			transactionHash: DATA, 32 Bytes - 交易的hash.
			transactionIndex: QUANTITY - 区块中交易index的位置。
			blockHash: DATA, 32 Bytes - 此交易所处的区块hash。
			blockNumber: QUANTITY - 此交易所处的区块号
			cumulativeGasUsed: QUANTITY - 当这笔交易已经在区块中执行完成，所使用的gas总量。
			gasUsed: QUANTITY - 此特定交易所使用的单个gas金额。
			contractAddress: DATA, 20 Bytes - 创建的合同地址（如果该交易是创建合约，* 否则为空。
			logs: Array - 此交易生成的日志对象数组。
*/
func (rpc *EthRPC) EthGetTransactionReceipt(hash string) (*TransactionReceipt, error) {
	transactionReceipt := new(TransactionReceipt)

	err := rpc.call("eth_getTransactionReceipt", transactionReceipt, hash)
	if err != nil {
		return nil, err
	}

	return transactionReceipt, nil
}

/**
EthGetCompilers returns a list of available compilers in the client.
返回客户端中可用编译器的列表。
参数
	none
返回
Array - 可以编译的数组列表。
*/
func (rpc *EthRPC) EthGetCompilers() ([]string, error) {
	compilers := []string{}

	err := rpc.call("eth_getCompilers", &compilers)
	return compilers, err
}

/**
	返回编译后的solidity代码。
	参数
		String - 源码
		params: [
  			 "contract test { function multiply(uint a) returns(uint d) {   return a * 7;   } }",
		]
	返回：
		DATA - 编译后的代码。
*/
func (rpc *EthRPC) EthGetCompileSolidity(codes string) (string, error) {
	var data string

	err := rpc.call("eth_compileSolidity", codes, &data)
	return data, err
}

/**
	返回编译后的LLL代码。
	参数
		String - 源码
		params: [
  				 "(returnlll (suicide (caller)))",
		]
	返回：
		DATA - 已编译的代码.
*/
func (rpc *EthRPC) EthGetCompileLLL(codes string) (string, error) {
	var data string

	err := rpc.call("eth_compileLLL", codes, &data)
	return data, err
}

/**
	返回编译后的LLL代码。
	参数
		String - 源码
		params: [
  				  " some serpent "，
		]
	返回：
		DATA - 已编译的代码.
*/
func (rpc *EthRPC) EthGetCompileSerpent(codes string) (string, error) {
	var data string

	err := rpc.call("eth_compileSerpent", codes, &data)
	return data, err
}

/**
	EthNewFilter creates a new filter object.
 	根据过滤器选项创建过滤器对象，以通知状态何时更改（日志）。要检查状态是否已更改，请调用eth_getFilterChanges。
	参数
		Object - 过滤器选项:
			fromBlock: QUANTITY|TAG - (optional, default: "latest") Integer block number, or "latest" for the last mined block or "pending", "earliest" for not yet mined transactions.
				（可选，默认值：“latest”）区块号，或最近一次挖掘块的“latest”或“pending”，“earliest”用于还未挖矿的交易。
			toBlock: QUANTITY|TAG - (optional, default: "latest") 区块号，或最近一次挖掘块的“latest”或“pending”，“earliest”用于还未挖矿的交易。
			address: DATA|Array, 20 Bytes -（可选）合约地址或日志应从其发出的地址列表。
			topics: DATA数组, - (可选) 32字节数据topic数组。
		params: [{
			"fromBlock": "0x1",
			"toBlock": "0x2",
			"address": "0x8888f1f195afa192cfee860698584c030f4c9db1",
			"topics": ["0x000000000000000000000000a94f5374fce5edbc8e2a8697c15331677e6ebf0b"]
		返回
			QUANTITY - 过滤id.
}]
*/
func (rpc *EthRPC) EthNewFilter(params FilterParams) (string, error) {
	var filterID string
	err := rpc.call("eth_newFilter", &filterID, params)
	return filterID, err
}

/**
EthNewBlockFilter creates a filter in the node, to notify when a new block arrives.
To check if the state has changed, call EthGetFilterChanges.
在节点中创建一个过滤器，以通知新块到达。要检查状态是否已更改，请调用eth_getFilterChanges。
参数
	None
返回
	QUANTITY - 过滤器id.
*/
func (rpc *EthRPC) EthNewBlockFilter() (string, error) {
	var filterID string
	err := rpc.call("eth_newBlockFilter", &filterID)
	return filterID, err
}

/**
EthNewPendingTransactionFilter creates a filter in the node, to notify when new pending transactions arrive.
To check if the state has changed, call EthGetFilterChanges.
在节点中创建过滤器，以通知新的待处理交易到达。要检查状态是否已更改，请调用eth_getFilterChanges。
参数
	None
返回
	QUANTITY - 过滤器id.
*/
func (rpc *EthRPC) EthNewPendingTransactionFilter() (string, error) {
	var filterID string
	err := rpc.call("eth_newPendingTransactionFilter", &filterID)
	return filterID, err
}

/**
	EthUninstallFilter uninstalls a filter with given id.
	在不再需要监控时，应停止调用，卸载指定ID的过滤器。 另外，过滤一段时间未使用eth_getFilterChanges请求的超时。
	参数
		QUANTITY - The filter id.
		params: [
  			"0xb" // 11
		]
	返回
		Boolean - 如果过滤器已成功卸载，则为true，否则为false。
*/
func (rpc *EthRPC) EthUninstallFilter(filterID string) (bool, error) {
	var res bool
	err := rpc.call("eth_uninstallFilter", &res, filterID)
	return res, err
}

/**
EthGetFilterChanges polling method for a filter, which returns an array of logs which occurred since last poll.
过滤器的poll方法，该方法返回自上次poll以来发生的日志数组。
参数
	QUANTITY - 过滤id。
		params: [
		  "0x16" // 22
		]

返回
	Array - 日志对象数组，或者如果自上次poll以来没有任何更改，则为空数组。
		对于用eth_newBlockFilter创建的过滤器，返回是块hahs（DATA，32字节），例如，[“0x3454645634534......”]。
		对于使用eth_newPendingTransactionFilter创建的过滤器，返回是事务hash（DATA，32字节），例如，[“0x6345343454645......”。
		对于使用eth_newFilter创建的过滤器，日志是具有以下参数：
		type: TAG - 等待日志处于待处理状态。如果日志已被开采，则开采。
		logIndex: QUANTITY - 块中日志索引位置。当状态为pending的日志时为null。
		transactionIndex: QUANTITY - 日志从交易index位置创建的整数。当其panding的日志时为null。
		transactionHash: DATA, 32 Bytes - 这个日志创建的交易的hash。 当其处于pending的日志时为null。
		blockHash: DATA, 32 Bytes - 该日志所在块的散列，当其处于pending，则为空。当其pending日志时也为null。
		blockNumber: QUANTITY - 此日志中的区块号，当它处于pending，则为null，当其日志处于pending，也为null。
		address: DATA, 20 Bytes - 该日志源的地址。
		data: DATA - 包含日志的一个或多个32字节未index的参数。
		topics: DATA的数组 - 索引日志参数的0到4 32字节数组数据。（在solidity中：第一个topic是事件签名的hash（例如Deposit(address，bytes32，uint256))，除非你使用匿名说明符声明该事件。）
*/
func (rpc *EthRPC) EthGetFilterChanges(filterID string) ([]Log, error) {
	var logs = []Log{}
	err := rpc.call("eth_getFilterChanges", &logs, filterID)
	return logs, err
}

/**
	EthGetFilterLogs returns an array of all logs matching filter with given id.
	指定id，返回匹配的所有日志数组。
	参数
		QUANTITY - 过滤器id。
		params: [
  			"0x16" // 22
		]
	返回
		参考 eth_getFilterChanges 返回
*/
func (rpc *EthRPC) EthGetFilterLogs(filterID string) ([]Log, error) {
	var logs = []Log{}
	err := rpc.call("eth_getFilterLogs", &logs, filterID)
	return logs, err
}

/**
	EthGetLogs returns an array of all logs matching a given filter object.
	通过过滤器对象，返回匹配的所有日志数组。
	参数
		Object - 过滤器对象, 参照 eth_newFilter参数.
		params: [{
  			"topics": ["0x000000000000000000000000a94f5374fce5edbc8e2a8697c15331677e6ebf0b"]
		}]
	返回
		参考 eth_getFilterChanges 返回
*/
func (rpc *EthRPC) EthGetLogs(params FilterParams) ([]Log, error) {
	var logs = []Log{}
	err := rpc.call("eth_getLogs", &logs, params)
	return logs, err
}

//eth<=>wei
var ethWeiPrice = new(big.Float).SetFloat64(1e+18)

//获取余额
func (rpc *EthRPC) GetBalance(from string) (float64, error) {
	wei, err := rpc.EthGetBalance(from, "latest")
	if err != nil {
		return 0, err
	}
	balance, _ := big.NewFloat(0).Quo(new(big.Float).SetInt(&wei), ethWeiPrice).Float64()
	return balance, err
}

//验证地址是否是ETH地址
func CheckETHAddr(addr string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(addr)
}
