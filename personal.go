package ethrpc

// EthGetBalance returns the balance of the account of given address in wei.
// 返回创建的地址
func (rpc *EthRPC) PersonalNewAccount(password string) (address string, err error) {
	var response string
	if err := rpc.call("personal_newAccount", &response, password); err != nil {
		return "创建失败", err
	}

	return response, err
}

//personal.unlockAccount-
//个人解锁账户
func (rpc *EthRPC) PersonalUnLockAccount(address, password string, duration ...int) (bool, error) {
	//默认解锁30秒
	second := 30
	if len(duration) > 0 {
		second = duration[0]
	}
	var lock bool
	err := rpc.call("personal_unlockAccount", &lock, address, password, second)
	return lock, err
}

/**
	sendTransaction方法验证指定的密码并提交交易，该方法的交易参数 与eth_sendTransaction一样，同时包含from账户地址。
    如果密码可以成功解密交易中from地址对应的私钥，那么该方法将验证交易、 签名并广播到以太坊网络中。
    由于在sendTransaction方法调用时，from账户并未在节点中全局解锁 （仅在该调用内解锁），因此from账户不能用于其他RPC调用
*/
func (rpc *EthRPC) PersonalSendTransaction(transaction T, password string) (string, error) {
	var hex string

	err := rpc.call("personal_sendTransaction", &hex, transaction, password)

	return hex, err
}

//personal_listAccounts
//个人账户列表
func (rpc *EthRPC) PersonalListAccounts() ([]string, error) {
	var list []string

	err := rpc.call("personal_listAccounts", &list)
	return list, err
}
