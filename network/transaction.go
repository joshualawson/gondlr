package network

import "math/big"

type Transaction struct {
	From        string   `json:"from"`
	To          string   `json:"to"`
	Amount      *big.Int `json:"amount"`
	BlockHeight *big.Int `json:"blockHeight"`
	Pending     bool     `json:"pending"`
	Confirmed   bool     `json:"confirmed"`
}
