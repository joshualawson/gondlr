package network

import "math/big"

type SolanaNetwork struct {
	signer     Signer
	privateKey string
}

func Solana(signer Signer, privateKey string) *SolanaNetwork {
	return &SolanaNetwork{
		signer:     signer,
		privateKey: privateKey,
	}
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
