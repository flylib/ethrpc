package ethrpc

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//version
const (
	v1 = "1.0"
	V2 = "2.0"
)

////http客户端接口
//type httpClient interface {
//	Post(url string, contentType string, body io.Reader) (*http.Response, error)
//}

//auth
const (
	AuthNone      = iota //none
	AuthBasicAuth        //基本认证
	AuthMD5              //MD5加密
	AuthToken            //token
)

// EthRPC - Ethereum rpc client
//以太坊RPC客户端结构体
type EthRPC struct {
	url string //URL链接
	//client    httpClient //客户端对象
	Version   string //版本
	Auth      int    //认证方式
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
		url: url,
		//client:  http.DefaultClient,
		Version: V2, //default 2.0
	}
	for _, option := range options {
		option(rpc)
	}
	return rpc
}

func (rpc *EthRPC) request(method string, target interface{}, params ...interface{}) error {
	result, err := rpc.call(method, params...)
	if err != nil {
		return err
	}
	return json.Unmarshal(result, target)
}

// URL returns client url
func (rpc *EthRPC) URL() string {
	return rpc.url
}

//设置版本
func SetVersion(v string) func(rpc *EthRPC) {
	return func(rpc *EthRPC) {
		rpc.Version = v
	}
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

// 设置认证方式
func (rpc *EthRPC) SetAuthType(authType int) {
	rpc.Auth = authType
}

func (rpc *EthRPC) call(method string, params ...interface{}) (json.RawMessage, error) {
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
	newReq, err := http.NewRequest("POST", rpc.url, bytes.NewBuffer(reqData))
	if err != nil {
		return nil, err
	}
	switch rpc.Auth {
	case AuthBasicAuth: //基本认证
		newReq.SetBasicAuth(rpc.BasicAuth.User, rpc.BasicAuth.PWD)
	}
	newReq.Header.Set("Content-Type", "application/json")
	//client := &http.Client{}
	res, err := http.DefaultClient.Do(newReq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	response := new(Response)
	if err := json.Unmarshal(resData, response); err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, *response.Error
	}

	return response.Result, nil
}

//// RawCall returns raw response of method call (Deprecated) (弃用)
//func (rpc *EthRPC) RawCall(method string, params ...interface{}) (json.RawMessage, error) {
//	return rpc.call(method, params...)
//}
