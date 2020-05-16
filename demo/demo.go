package main

import (
	"fmt"
	"github.com/Quantumoffices/ethrpc"
)

func main() {
	rpc := ethrpc.NewEthRPC("http://192.168.0.220:8545")
	version, _ := rpc.Web3ClientVersion()
	fmt.Println(version)
	count, _ := rpc.NetPeerCount()
	fmt.Println(count)
	sync, _ := rpc.EthSyncing()
	fmt.Println(fmt.Sprintf("EthSyncing=%#v", sync))
	block, err := rpc.EthGetBlockByNumber(sync.CurrentBlock, false)
	if err != nil {
		fmt.Println("error get EthGetBlockByNumber = ", err)
	}
	//t, err := rpc.EthGetTransactionByHash(block.Hash)
	//fmt.Println("block=", block)
	fmt.Println(fmt.Sprintf("transaction=%#v", block.Transactions))
	ethblocknumber, _ := rpc.EthBlockNumber()
	fmt.Println("EthBlockNumber=", ethblocknumber)
	addr, _ := rpc.EthCoinbase()
	fmt.Println("EthCoinbase=", addr)
	blockNum, _ := rpc.EthBlockNumber()
	fmt.Println("blockNum=", blockNum)
	isMining, _ := rpc.EthMining()
	fmt.Println(isMining)
	adds, _ := rpc.EthAccounts()
	fmt.Println(adds)
	balance, err := rpc.EthGetBalance(adds[1], ethrpc.BlockTag_Latest)
	if err != nil {
		fmt.Println("error get EthGetBalance = ", err)
	}
	fmt.Println("balance=", balance.Int64(), 1e+18)
	price, _ := rpc.EthGasPrice()
	fmt.Println("price=", price)
	prvVersion, _ := rpc.EthProtocolVersion()
	fmt.Println("prvVersion=", prvVersion)

	blockHight, err := rpc.EthBlockNumber()
	if err != nil {
		panic(err)
		return
	}

	curBlock, err := rpc.EthGetBlockByNumber(blockHight, false)
	if err != nil {
		panic(err)
	}
	fmt.Println(curBlock)
}
