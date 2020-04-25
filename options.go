package ethrpc

type OptionFun func(rpc *EthRPC)

//版本
func WithVersion(version string) OptionFun {
	return func(rpc *EthRPC) {
		rpc.Version = version
	}
}
