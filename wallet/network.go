package wallet

import (
	"github.com/joshualawson/gondlr/network"
	"math/big"
)

type Network interface {
	Name() string
	Tx(transactionID string) (network.Transaction, error)
	CurrentHeight() int
	Fee() *big.Int
	SendTransaction(data interface{}) error
	CreateTransaction(amount *big.Int, to string) error
	PublicKey() string
}
