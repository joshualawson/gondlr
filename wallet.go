package gondlr

import "github.com/joshualawson/gondlr/wallet"

type Wallet interface {
	PrivateKeyBytes() []byte
	PublicKeyBytes() []byte
	PublicKey() wallet.PublicKey
	PrivateKey() wallet.PrivateKey
}
