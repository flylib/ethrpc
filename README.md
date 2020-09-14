# 以太坊JsonRpc协议


## 文档参考
* [英文文档](https://github.com/ethrpc/wiki/wiki/JSON-RPC)
* [中文文档](http://cw.hubwiz.com/card/c/ethereum-json-rpc-api/)
* [Geth管理API文档](http://cw.hubwiz.com/card/c/geth-rpc-api/1/4/2/)
* [用Go来做以太坊开发](https://goethereumbook.org/zh/)
* https://www.jsonrpc.org/specification

##代币价格 

|Unit|	Wei Value	|Wei|
| --- | --- |---
wei	|1	|1| wei
Kwei (babbage)|	1e3 wei|	1,000
Mwei (lovelace)|	1e6 wei	|1,000,000
Gwei (shannon)	|1e9 wei|	1,000,000,000
microether (szabo)	|1e12 wei|	1,000,000,000,000
milliether (finney)	|1e15 wei|	1,000,000,000,000,000
ether	|1e18 wei	|1,000,000,000,000,000,000

##Gas-燃料
* Gas limit:燃料限制，是用户愿意为执行某个操作或确认交易支付的最大Gas量（最少21,000）、
* Gas Price:燃料价格，用户愿意花费于每个 Gas 单位的价钱。
* 1Gwei(Gas Price)≈0.00000002 ETH
* 0.00000002*21000=0.00042ETH
* 目前为止，确认交易使用 1 Gwei 需要大约30分钟，而用 40 Gwei 大约1-2分钟。

##常用命令
* eth.accounts 查看现有账户，会返回一个数组，数组的每一项是现有的账号地址。在未创建账户的情况下，初始状态默认应该是返回空数组。eth.accounts[0] 获取第0个账户的地址
* personal.newAccount("123") 创建一个账户，括号里面的参数是所创建的账户的密码
* personal.unlockAccount(eth.accounts[0]) 解锁第0个账户，会提示输入密码以进行解锁，一个账户只有解锁后，才能转出其中的币
* eth.getBalance(eth.accounts[0]) 查看第0个账户的余额，初始应该是0


##参考
* https://www.jianshu.com/p/b56552b1d1a0以太坊中代币数量的计量单位说明