package ethrpc

/**
NetVersion returns the current network protocol version.
返回当前网络协议的版本。
参数：
none
返回：
String - 当前网络协议版本
*/
func (rpc *EthRPC) NetVersion() (string, error) {
	var version string
	err := rpc.request("net_version", &version)
	return version, err
}

/**
NetListening returns true if client is actively listening for network connections.
如果客户端正在主动监听网络连接，则返回true。
参数：
none
返回：
Boolean - 当监听时为true，否则为false。
*/
func (rpc *EthRPC) NetListening() (bool, error) {
	var listening bool

	err := rpc.request("net_listening", &listening)
	return listening, err
}

/**
NetPeerCount returns number of peers currently connected to the client.
返回当前连接到客户端的对等的数量。
参数：
none
返回：
QUANTITY - 已连接对等的数量的整数。
*/
func (rpc *EthRPC) NetPeerCount() (int, error) {
	var response string
	if err := rpc.request("net_peerCount", &response); err != nil {
		return 0, err
	}

	return ParseInt(response)
}
