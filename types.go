package ethrpc

//一个Promise对象，其解析值为RLP编码的交易对象。该对象的raw属性可以用来通过web3.eth.sendSignedTransaction() 方法来发送交易。
//{raw: '0xf86c808504a817c800825208943535353535353535353535353535353535353535880de0b6b3a76400008025a04f4c17305743700648bc4f6cd3038ec6f6af0df73e31757007b7f59df7bee88da07e1941b264348e80c78c4027afc65a87b0a5e43e86742b8ca0823584c6788fd0',
//tx: {
//nonce: '0x0',
//gasPrice: '0x4a817c800',
//gas: '0x5208',
//to: '0x3535353535353535353535353535353535353535',
//value: '0xde0b6b3a7640000',
//input: '0x',
//v: '0x25',
//r: '0x4f4c17305743700648bc4f6cd3038ec6f6af0df73e31757007b7f59df7bee88d',
//s: '0x7e1941b264348e80c78c4027afc65a87b0a5e43e86742b8ca0823584c6788fd0',
//hash: '0xda3be87732110de6c1354c83770aae630ede9ac308d9f7b399ecfba23d923384'
// }}
type SignPromise struct {
	Raw string `json:"raw"`
	Tx  struct {
		Nonce    string `json:"nonce"`
		Gas      string `json:"gas"`
		GasPrice string `json:"gasPrice"`
		To       string `json:"to"`
		Value    string `json:"value"`
		Input    string `json:"input"`
		V        string `json:"v"`
		R        string `json:"r"`
		Hash     string `json:"hash"` //交易hash
	} `json:"tx"`
}
