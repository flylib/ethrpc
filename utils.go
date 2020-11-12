package ethrpc

import (
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"
)

/**
 *@Project     openchain
 *@Author      king
 *@CreateTime  2020/4/18 1:42 下午
 *@ClassName   utils
 *@Description TODO  数据操作工具类
 */

// 将十六进制字符串值解析为int
func ParseInt(value string) (int, error) {
	i, err := strconv.ParseInt(strings.TrimPrefix(value, "0x"), 16, 64)
	if err != nil {
		log.Panicln("err:", err.Error())
		return 0, err
	}
	return int(i), nil
}

// 将十六进制字符串值解析为big.Int
func ParseBigInt(value string) (big.Int, error) {
	i := big.Int{}
	_, err := fmt.Sscan(value, &i)
	return i, err
}

// 将int转换为十六进制表示
func IntToHex(i int) string {
	return fmt.Sprintf("0x%x", i)
}

// BigToHex covert big.Int to hexadecimal representation
//到十六进制 长整型到十六进制表示
func BigToHex(bigInt big.Int) string {
	if bigInt.BitLen() == 0 {
		return "0x0"
	}

	return "0x" + strings.TrimPrefix(fmt.Sprintf("%x", bigInt.Bytes()), "0")
}

/**
将给定的以太金额转换为以wei为单位的数值。
web3.toWei(number, unit)

参数：
	number：Number|String|BigNumber - 金额
	unit ： String - 字符串单位
		noether: ‘0’
		wei: ‘1’
		kwei: ‘1000’
		Kwei: ‘1000’
		babbage: ‘1000’
		femtoether: ‘1000’
		mwei: ‘1000000’
		Mwei: ‘1000000’
		lovelace: ‘1000000’
		picoether: ‘1000000’
		gwei: ‘1000000000’
		Gwei: ‘1000000000’
		shannon: ‘1000000000’
		nanoether: ‘1000000000’
		nano: ‘1000000000’
		szabo: ‘1000000000000’
		microether: ‘1000000000000’
		micro: ‘1000000000000’
		finney: ‘1000000000000000’
		milliether: ‘1000000000000000’
		milli: ‘1000000000000000’
		ether: ‘1000000000000000000’
		kether: ‘1000000000000000000000’
		grand: ‘1000000000000000000000’
		mether: ‘1000000000000000000000000’
		gether: ‘1000000000000000000000000000’
		tether: ‘1000000000000000000000000000000’
返回：
	String|BigNumber - 根据传入参数的不同，分别是字符串形式的字符串，或者是BigNumber。
示例：
	var value = web3.toWei('1', 'ether');
	console.log(value); // "1000000000000000000"
注意，wei是最小的以太单位，应当总是使用wei进行计算，仅在需要显示 时进行转换。

*/
func (rpc *EthRPC) ToWei(number int64, unit string) (string, error) {
	var data string
	err := rpc.request("web3_", &data, number, unit)
	return data, err
}
