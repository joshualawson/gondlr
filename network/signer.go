package network

import "github.com/joshualawson/gondlr/signers"

type Signer interface {
	PublicKey() signers.PublicKey
	Sign(data []byte) ([]byte, error)
	Verify(data []byte, sig []byte) bool
}
