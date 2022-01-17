package network

import "math/big"

type Ethereum struct {
	signer     Signer
	privateKey string
}

func (a *Ethereum) Tx(transactionID string) (Transaction, error) {
	return Transaction{}, nil
}

func (a *Ethereum) CurrentHeight() *big.Int {
	return nil
}

func (a *Ethereum) Fee() *big.Int {
	return nil
}

func (a *Ethereum) SendTransaction(data interface{}) error {
	return nil
}

func (a *Ethereum) CreateTransaction(amount *big.Int, to string) error {
	return nil
}
