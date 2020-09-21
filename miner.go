package ethrpc

/**
miner的setEtherBase方法用来设置接收挖矿奖励的以太坊账号。
sets the etherbase, where mining rewards will go.

参数：
	"params": [address] Address to send the rewards when mining.

返回：
	bool - 是否设置成功、
	Boolean - true if changed, otherwise false.
*/
func (rpc *EthRPC) MinerSetEtherBase(addr string) (bool, error) {
	var data bool
	err := rpc.call("miner_setEtherbase", &data, addr)
	return data, err
}

/**
miner的setGasPrice方法用来设置挖矿时交易的最小可接受 gas价格，任何低于此设置值的交易将被排除在挖矿过程之外。
sets the minimal accepted gas price when mining transactions. Any transactions that are below this limit are excluded from the mining process.

参数：
	price:Integer - Value of the gas price.

返回：
	Boolean - true if changed, otherwise false.
示例:
	miner_setGasPrice(10)
*/
func (rpc *EthRPC) MinerSetGasPrice(price int64) (bool, error) {
	var data bool
	err := rpc.call("miner_setGasPrice", &data, price)
	return data, err
}

/**
miner的start方法启动具有指定线程数量的CPU挖矿进程，并根据需要生成新的有向无环图。
starts the CPU mining process with the given number of threads.

参数：
	threads:Integer - Given number of threads used for mining.
返回：
	 Boolean
示例:
	miner.Start(1) bool
*/
func (rpc *EthRPC) MinerStart(threads int64) (bool, error) {
	var data bool
	err := rpc.call("miner_start", &data, threads)
	return data, err
}

/**
miner的stop方法用来停止CPU挖矿操作的执行。
Mining process to be stopped.

参数：
返回：
	Boolean - 是否停止成功
示例:
	miner.Stop() bool
*/
func (rpc *EthRPC) MinerStop() (bool, error) {
	var data bool
	err := rpc.call("miner_stop", &data)
	return data, err
}

/**
miner的getHashrate方法返回当前挖矿的哈希生成速率

参数：
返回：
   Integer- "result": 1000
示例:
	miner.GetHashrate()
*/
func (rpc *EthRPC) MinerGetHashRate() (int64, error) {
	var data int64
	err := rpc.call("miner_getHashrate", &data)
	return data, err
}

/**
Set the number of miner threads. By default, it is set to the number of available CPUs on the machine.
参数：
  Integer- threads
返回：
   Boolean-
   {
	  "id": 21,
	  "jsonrpc": "2.0",
	  "result": true
	}
示例:
	miner.setThreads(3)
*/
//func (rpc *EthRPC) MinerSetThreads(threads int64) (bool, error) {
//	var data bool
//	err := rpc.call("miner_setThreads", &data, threads)
//	return data, err
//}

/**
Checks whether the CPU miner is running.
参数：
	None
返回：
   Boolean-
   {
	  "id": 21,
	  "jsonrpc": "2.0",
	  "result": true
	}
示例:
	 miner.isMining()
*/
func (rpc *EthRPC) MinerIsMining() (bool, error) {
	var data bool
	err := rpc.call("miner_isMining", &data)
	return data, err
}
