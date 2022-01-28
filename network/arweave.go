package network

import (
	arweave "github.com/joshualawson/arweave-api"
	"math/big"
)

type ArweaveNetworkCaller interface {
	Transaction(id string) (arweave.Transaction, error)
	TransactionStatus(id string) (arweave.TransactionStatus, error)
	Info() (arweave.Info, error)
}

type ArweaveNetwork struct {
	ar ArweaveNetworkCaller
}

func Arweave() *ArweaveNetwork {
	return &ArweaveNetwork{}
}

func (a *ArweaveNetwork) Name() string {
	return "arweave"
}

func (a *ArweaveNetwork) Currency() []string {
	return CurrencyArweave
}

func (a *ArweaveNetwork) Base() float64 {
	return float64(BaseArweave)
}

func (a *ArweaveNetwork) Tx(transactionID string) (Transaction, error) {
	t, err := a.ar.Transaction(transactionID)
	if err != nil {
		return Transaction{}, ErrorSDK(err)
	}

	if t.Pending {
		return Transaction{Pending: true}, nil
	}

	ts, err := a.ar.TransactionStatus(transactionID)
	if err != nil {
		return Transaction{}, ErrorSDK(err)
	}

	confirmed := false
	if !t.Pending && ts.NumberOfConfirmations >= 10 {
		confirmed = true
	}

	amount, _ := new(big.Int).SetString(t.Quantity, 10)

	return Transaction{
		From:      t.Owner,
		To:        t.Target,
		Amount:    amount,
		Pending:   t.Pending,
		Confirmed: confirmed,
	}, nil
}

func (a *ArweaveNetwork) CurrentHeight() int {
	i, err := a.ar.Info()
	if err != nil {
		return 0
	}
	return i.Height
}

func (a *ArweaveNetwork) Fee() *big.Int {
	return nil
}

func (a *ArweaveNetwork) SendTransaction(data interface{}) error {
	return nil
}

func (a *ArweaveNetwork) CreateTransaction(amount *big.Int, to string) error {
	return nil
}
