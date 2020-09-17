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
返回指定UTF-8字符串的16进制表示。

参数：
	string - String: ·UTF-8字符串
返回：
	String: 16进制字符串
*/
func (rpc *EthRPC) UTF8ToHex(str string) (string, error) {
	var data string
	err := rpc.call("utils_utf8ToHex", &data, str)
	return data, err
}
