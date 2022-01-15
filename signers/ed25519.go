package signers

import (
	"crypto"
	"crypto/ed25519"
	"crypto/rand"
)

type Ed25519Signer struct {
	privateKey []byte
}

func Ed25519(privateKey string) *Ed25519Signer {
	return &Ed25519Signer{
		privateKey: []byte(privateKey),
	}
}

// PublicKey ...
func (e *Ed25519Signer) PublicKey() PublicKey {
	p := ed25519.PrivateKey(e.privateKey)
	pubKey := p.Public().(ed25519.PublicKey)

	return []byte(pubKey)
}

// Sign ...
func (e *Ed25519Signer) Sign(data []byte) ([]byte, error) {
	p := ed25519.PrivateKey(e.privateKey)
	signature, err := p.Sign(rand.Reader, data, crypto.Hash(0))
	if err != nil {
		return nil, err
	}

	return signature[:64], nil
}

// Verify ...
func (e *Ed25519Signer) Verify(data []byte, sig []byte) bool {
	return ed25519.Verify(ed25519.PublicKey(e.PublicKey()), data, sig)
}
