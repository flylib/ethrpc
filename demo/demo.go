package main

import (
	"errors"
	"fmt"
	"github.com/zjllib/ethrpc"
	"math/big"
)

func Transfer(from, to, pwd string, amount float64) error {
	amount = amount * 1e+18
	balanceBig, err := rpcJson.EthGetBalance(from, "latest")
	if err != nil {
		return err
	}

	balance, _ := new(big.Float).SetInt(&balanceBig).Float64()
	if balance < amount {
		return errors.New("余额不足！")
	}
	ok, err := rpcJson.PersonalUnLockAccount(from, pwd)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("解锁钱包失败!")
	}
	nonce, err := rpcJson.EthGetTransactionCount(from, "pending")
	fmt.Println("nonce", nonce)
	price, err := rpcJson.EthGasPrice()
	if err != nil {
		panic(err)
	}
	//fmt.Println(:"",price.Int64())

	t := ethrpc.T{
		From:     from,
		To:       to,
		Gas:      600000, //600000  default:21000
		GasPrice: &price, //big.NewInt(4500000000) 最快到账 60000000000 普通：20000000000   default:1000000000
		Value:    big.NewInt(int64(amount)),
		Nonce:    nonce,
	}
	//transaction, err := rpcJson.EthSendTransaction(t)
	//if err != nil {
	//	panic(err)
	//	return err
	//}
	//fmt.Println(transaction)

	s, err := rpcJson.SignTransaction(t)
	//fmt.Println(s, err)
	fmt.Printf("%+v", s)
	fmt.Println()
	ss, err := rpcJson.EthSendRawTransaction(s.Raw)
	fmt.Println(ss, err)
	//bytes, _ := json.Marshal(t)
	//// 将 byte 装换为 16进制的字符串
	//hexStringData := hex.EncodeToString(bytes)
	//fmt.Println("hexStringData=", "0x"+hexStringData)
	//data, err := rpcJson.EthSign(from, "0x"+hexStringData)
	//s, err := rpcJson.EthSendRawTransaction(data)
	//if err != nil {
	//	return err
	//}
	//fmt.Println(s)
	return err
}

var rpcJson = ethrpc.NewEthRPC("http://192.168.0.122:39005")

//获取余额
func getBlance(from string) (float64, error) {
	wei, err := rpcJson.EthGetBalance(from, "latest")
	balance, _ := new(big.Float).SetInt(&wei).Float64()
	return balance / 1e+18, err
}

func main() {
	//ok, err := rpcJson.EthPerUnLockAccount("0xd5693661f30c834f63577380ed0f673b7ef9d3c9", "123456")
	//if err != nil {
	//	panic(err)
	//}
	//if ok {
	//	data, err := rpcJson.EthSign("0xd5693661f30c834f63577380ed0f673b7ef9d3c9", "hello world")
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(data)
	//}

	//rpc := ethrpc.NewEthRPC("http://47.57.69.47:8785")
	//rpcJson := ethrpc.NewEthRPC("http://192.168.0.122:39005")

	//threads, i3 := rpc.MinerSetThreads(2)
	//fmt.Println(threads, i3)
	//b, i2 := rpc.MinerStart(1)
	//fmt.Println(b, i2)
	//return
	//s, i := rpc.MinerSetEtherBase("0xd5693661f30c834f63577380ed0f673b7ef9d3c9")
	//fmt.Println(s, i)
	//wei, e := rpc.EthGetBalance("0xe7f344e7b95c8a19d7d77ea27b86ee2fd6d776cf", "latest")
	index, err2 := rpcJson.EthGetUncleByBlockNumberAndIndex(129, 0)
	fmt.Println(index, err2)
	gasPrice, i := rpcJson.EthGasPrice()
	fmt.Println(gasPrice.Int64(), i)

	//wei, e := rpcJson.EthGetBalance("0xd5693661f30c834f63577380ed0f673b7ef9d3c9", "latest")
	////wei, e := rpc.EthGetBalance("0x332d51875124dcabcf46bb4be7ec1cb81b1d1f32", "latest")
	//if e != nil {
	//	panic(e)
	//}
	//fmt.Println(wei.Int64())
	//fmt.Println(new(big.Float).SetInt(&wei).Float64() / 1e+18)
	f, e := getBlance("0xd5693661f30c834f63577380ed0f673b7ef9d3c9")
	if e != nil {
		panic(e)
	}
	fmt.Println(f)
	Transfer("0xd5693661f30c834f63577380ed0f673b7ef9d3c9", "0x332d51875124dcabcf46bb4be7ec1cb81b1d1f32", "123456", 1)
	return
	//t := ethrpc.T{
	//	From:     "0xe7f344e7b95c8a19d7d77ea27b86ee2fd6d776cf",
	//	To:       "0xe7f344e7b95c8a19d7d77ea27b86ee2fd6d776cf",
	//	Gas:      600000,                 //600000
	//	GasPrice: big.NewInt(4500000000), //big.NewInt(4500000000) 最快到账 60000000000 普通：20000000000
	//	Value:    big.NewInt(int64(10000000000000)),
	//}
	//bytes, e := json.Marshal(t)
	//if e != nil {
	//	panic(e)
	//}
	//fmt.Println(string(bytes))
	//fmt.Println(1e+18)
	//fmt.Println(wei.Int64())
	//fmt.Println(float64(wei.Int64()) / 0.02895)
	version, _ := rpcJson.Web3ClientVersion()
	fmt.Println(version)
	count, _ := rpcJson.NetPeerCount()
	fmt.Println(count)
	sync, _ := rpcJson.EthSyncing()
	fmt.Println(fmt.Sprintf("EthSyncing=%#v", sync))
	block, err := rpcJson.EthGetBlockByNumber(sync.CurrentBlock, false)
	if err != nil {
		fmt.Println("error get EthGetBlockByNumber = ", err)
	}
	//t, err := rpc.EthGetTransactionByHash(block.Hash)
	//fmt.Println("block=", block)
	fmt.Println(fmt.Sprintf("transaction=%#v", block.Transactions))
	ethblocknumber, _ := rpcJson.EthBlockNumber()
	acounts, _ := rpcJson.EthAccounts()
	fmt.Println("EthBlockNumber=", ethblocknumber, acounts)
	addr, _ := rpcJson.EthCoinbase()
	fmt.Println("EthCoinbase=", addr)
	blockNum, _ := rpcJson.EthBlockNumber()
	fmt.Println("blockNum=", blockNum)
	isMining, _ := rpcJson.EthMining()
	fmt.Println(isMining)
	adds, _ := rpcJson.EthAccounts()
	fmt.Println(adds)
	balance, err := rpcJson.EthGetBalance(adds[1], ethrpc.BlockTag_Latest)
	if err != nil {
		fmt.Println("error get EthGetBalance = ", err)
	}
	fmt.Println("balance=", balance.Int64(), 1e+18)
	price, _ := rpcJson.EthGasPrice()
	fmt.Println("price=", price)
	prvVersion, _ := rpcJson.EthProtocolVersion()
	fmt.Println("prvVersion=", prvVersion)

	blockHight, err := rpcJson.EthBlockNumber()
	if err != nil {
		panic(err)
		return
	}

	curBlock, err := rpcJson.EthGetBlockByNumber(blockHight, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(curBlock)
}
