package network

import "math/big"

type SolanaNetwork struct {
	privateKey string
}

func Solana() *SolanaNetwork {
	return &SolanaNetwork{}
}

func (a *SolanaNetwork) Name() string {
	return "solana"
}

func (a *SolanaNetwork) Tx(transactionID string) (Transaction, error) {
	return Transaction{}, nil
}

func (a *SolanaNetwork) CurrentHeight() *big.Int {
	return nil
}

func (a *SolanaNetwork) Fee() *big.Int {
	return nil
}

func (a *SolanaNetwork) SendTransaction(data interface{}) error {
	return nil
}

func (a *SolanaNetwork) CreateTransaction(amount *big.Int, to string) error {
	return nil
}
