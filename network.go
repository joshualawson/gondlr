package gondlr

import (
	"github.com/joshualawson/gondlr/network"
	"math/big"
)

type Network interface {
	Name() string
	Currency() []string
	Base() float64
	Tx(transactionID string) (network.Transaction, error)
	CurrentHeight() int
	Fee() *big.Int
	SendTransaction(data interface{}) error
	CreateTransaction(amount *big.Int, to string) error
}
