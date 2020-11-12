package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/zjllib/ethrpc"
	"io/ioutil"
	"math/big"
	"os"
)

//var rpcJson = ethrpc.NewEthRPC("http://192.168.0.122:39005")
var rpcJson = ethrpc.NewEthRPC("http://explorer.gbcoin.gold/node")

//外网ETH节点：http://www.auecoin.com/eth
//用户名：gethnode88669s1 密码：Gh875ds22kws03232wq
func init() {
	return
	//无需认证：http://www.auecoin.com/ethyyuyewnkks7e453sxsa76w3sakanode
	//act := ethrpc.Account{
	//	User: "gethnode88669s1",
	//	PWD:  "Gh875ds22kws03232wq",
	//}
	////rpcJson = ethrpc.NewEthRPC("http://www.auecoin.com/eth", ethrpc.SetAuthType(ethrpc.AuthBasicAuth),
	//rpcJson = ethrpc.NewEthRPC("http://www.auecoin.com/ethyyuyewnkks7e453sxsa76w3sakanode", ethrpc.SetAuthType(ethrpc.AuthNone),
	//	ethrpc.SetBasicAuth(act), ethrpc.Debug())
	number, err := rpcJson.EthBlockNumber()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(number)
	byNumber, err := rpcJson.EthGetBlockByNumber(0, false)
	fmt.Println(byNumber, err)
}

func Transfer(from, to, pwd string, amount float64) error {
	bigAmount := big.NewInt(0).Mul(big.NewInt(int64(amount*1e+8)), big.NewInt(1e+10))
	//amount = amount * 1e+18
	//balanceBig, err := rpcJson.EthGetBalance(from, "latest")
	//if err != nil {
	//	return err
	//}
	//
	//balance, _ := new(big.Float).SetInt(bigAmount).Float64()
	//if balance < amount {
	//	return errors.New("余额不足！")
	//}
	fmt.Println(len("0xcf06ad2594ce2a7315d3c8bfec1dc76d2992ec6ac14a4dce8fc3546afc5addcc"))
	blance, err2 := GetBalance(from)
	fmt.Println(blance, err2)
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
	//ethrpc.OfflineTransfer(1, uint64(nonce), to, bigAmount, "")

	fmt.Println(int64(amount))
	t := ethrpc.T{
		From:     from,
		To:       to,
		Gas:      21000,  //600000  default:21000
		GasPrice: &price, //big.NewInt(4500000000) 最快到账 60000000000 普通：20000000000   default:1000000000
		Value:    bigAmount,
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

//获取余额
//获取余额
func GetBalance(from string) (float64, error) {
	wei, err := rpcJson.EthGetBalance(from, "latest")
	balance, _ := big.NewFloat(0).Quo(new(big.Float).SetInt(&wei), new(big.Float).SetFloat64(1e+18)).Float64()
	return balance, err
}

var (
	file     = flag.String("file", "./UTC--2020-09-19T06-00-03.502591131Z--d5693661f30c834f63577380ed0f673b7ef9d3c9", "file")
	password = flag.String("password", "123456", "password")
)

func init() {
	//flag.Parse()
	//if _, err := os.Stat(*file); os.IsNotExist(err) {
	//	flag.Usage()
	//	os.Exit(1)
	//}
	//keyjson, err := ioutil.ReadFile(*file)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(keyjson))
}

func DecryptKeystore() {
	if _, err := os.Stat(*file); os.IsNotExist(err) {
		flag.Usage()
		os.Exit(1)
	}

	keyjson, err := ioutil.ReadFile(*file)
	if err != nil {
		panic(err)
	}
	addr, key, err := ethrpc.ParseKeystore(keyjson, *password)
	if err != nil {
		panic(err)
	}
	fmt.Println(addr, "----", key)
}

func main() {
	DecryptKeystore()
	//file, err := ioutil.ReadFile("./demo/UTC--2020-09-19T06-00-03.502591131Z--d5693661f30c834f63577380ed0f673b7ef9d3c9")
	//if err != nil {
	//	panic(err)
	//}
	////key, err := ethrpc.KeystoreToPrivateKey(file, "123456")
	////if err != nil {
	////	panic(err)
	////}
	////fmt.Println(key)
	//fmt.Println(file)
	//return
	//fmt.Println(big.NewInt(0).Mul(big.NewInt(1e+8), big.NewInt(1e+10)).Int64())
	//fmt.Println(len("1000000000000000000"))
	//fmt.Println(fmt.Sprintf("%f", 1e+18))
	//fmt.Println()
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
	//index, err2 := rpcJson.EthGetUncleByBlockNumberAndIndex(129, 0)
	//fmt.Println(index, err2)
	//gasPrice, i := rpcJson.EthGasPrice()
	//fmt.Println(gasPrice.Int64(), i)

	//wei, e := rpcJson.EthGetBalance("0xd5693661f30c834f63577380ed0f673b7ef9d3c9", "latest")
	////wei, e := rpc.EthGetBalance("0x332d51875124dcabcf46bb4be7ec1cb81b1d1f32", "latest")
	//if e != nil {
	//	panic(e)
	//}
	//fmt.Println(wei.Int64())
	//fmt.Println(new(big.Float).SetInt(&wei).Float64() / 1e+18)
	//f, e := GetBlance("0x07fd99da07693d14d72bed5fee61f0f706cbaa74")
	//if e != nil {
	//	panic(e)
	//}
	//fmt.Println(f)
	//线上0xe4287734d19147259a4307bea5282b6648a79bb3
	//
	//Transfer(
	//	"0x5922a85090f1d2a476ccbb4659fbf93fe2943716",
	//	"0xe4287734d19147259a4307bea5282b6648a79bb3",
	//	"gbchin2020", 800)
	//
	////质押
	//Transfer(
	//	"0xd5693661f30c834f63577380ed0f673b7ef9d3c9",
	//	"0xbdaf3976466e0531b377aab2432c9645506afb46",
	//	"123456", 10)
	////质押
	//Transfer(
	//	"0xd5693661f30c834f63577380ed0f673b7ef9d3c9",
	//	"0xbdaf3976466e0531b377aab2432c9645506afb46",
	//	"123456", 10) //质押
	//Transfer(
	//	"0xd5693661f30c834f63577380ed0f673b7ef9d3c9",
	//	"0xbdaf3976466e0531b377aab2432c9645506afb46",
	//	"123456", 10)
	//燃烧池
	//Transfer(
	//	"0xd5693661f30c834f63577380ed0f673b7ef9d3c9",
	//	"0x46f8178169312bce2405283bc9be2b3e3bc489f9",
	//	"123456", 200)
	//balance, err2 := rpcJson.GetBalance("0x46f8178169312bce2405283bc9be2b3e3bc489f9")
	//fmt.Println(balance, err2)
	return
}
