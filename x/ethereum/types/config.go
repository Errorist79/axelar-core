package types

// EthConfig contains all Ethereum module configuration values
type EthConfig struct {
	EthRpcAddr    string `mapstructure:"rpc_addr"`
	WithEthBridge bool   `mapstructure:"start-with-bridge"`
}

// DefaultConfig returns a configuration populated with default values
func DefaultConfig() EthConfig {
	return EthConfig{
		EthRpcAddr:    "http://127.0.0.1:7545",
		WithEthBridge: true,
	}
}
