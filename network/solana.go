package network

import "math/big"

type Solana struct {
	signer     Signer
	privateKey string
}

func (a *Solana) Tx(transactionID string) (Transaction, error) {
	return Transaction{}, nil
}

func (a *Solana) CurrentHeight() *big.Int {
	return nil
}

func (a *Solana) Fee() *big.Int {
	return nil
}

func (a *Solana) SendTransaction(data interface{}) error {
	return nil
}

func (a *Solana) CreateTransaction(amount *big.Int, to string) error {
	return nil
}
