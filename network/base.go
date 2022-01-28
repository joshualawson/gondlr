package network

type Currency []string
type Base float64

var (
	CurrencyArweave    = Currency{"AR", "winston"}
	CurrencyEthereum   = Currency{"ETH", "wei"}
	CurrencySolanaBase = Currency{"SOL", "lamports"}
)

const (
	BaseArweave  Base = 1e12
	BaseEthereum Base = 1e18
	BaseSolana   Base = 1e9
)
