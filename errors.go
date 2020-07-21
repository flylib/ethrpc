package ethrpc

import "fmt"

//EthError- ethrpc error 错误
type EthError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//错误信息
func (err EthError) Error() string {
	return fmt.Sprintf("Error %d (%s)", err.Code, err.Message)
}
