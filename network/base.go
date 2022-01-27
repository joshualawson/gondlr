package network

type Currency string
type Base int

const (
	CurrencyArweave    Currency = "winston"
	CurrencyEthereum   Currency = "wei"
	CurrencySolanaBase Currency = "lamports"
)

var NetworkToCurrency = map[string]Currency{
	"arweave":  CurrencyArweave,
	"ethereum": CurrencyEthereum,
	"solana":   CurrencySolanaBase,
}

var CurrencyBase = map[Currency]float64{
	CurrencyArweave: 1e12,
}
