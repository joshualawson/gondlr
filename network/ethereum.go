package network

import "math/big"

type EthereumNetwork struct {
	signer     Signer
	privateKey string
}

func Ethereum(signer Signer, privateKey string) *EthereumNetwork {
	return &EthereumNetwork{
		signer:     signer,
		privateKey: privateKey,
	}
}

func (a *EthereumNetwork) Name() string {
	return "ethereum"
}

func (a *EthereumNetwork) Tx(transactionID string) (Transaction, error) {
	return Transaction{}, nil
}

func (a *EthereumNetwork) CurrentHeight() *big.Int {
	return nil
}

func (a *EthereumNetwork) Fee() *big.Int {
	return nil
}

func (a *EthereumNetwork) SendTransaction(data interface{}) error {
	return nil
}

func (a *EthereumNetwork) CreateTransaction(amount *big.Int, to string) error {
	return nil
}
