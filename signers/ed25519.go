package signers

import (
	"crypto"
	"crypto/ed25519"
	"crypto/rand"
)

type Ed25519Signer struct {
	privateKey ed25519.PrivateKey
	publicKey  ed25519.PublicKey
}

func Ed25519(privateKey []byte, publicKey []byte) *Ed25519Signer {
	return &Ed25519Signer{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

// Sign calculated the signature of a digest
func (e *Ed25519Signer) Sign(data []byte) ([]byte, error) {
	signature, err := e.privateKey.Sign(rand.Reader, data, crypto.Hash(0))
	if err != nil {
		return nil, err
	}

	return signature[:64], nil
}

// Verify verifies a signature
func (e *Ed25519Signer) Verify(data []byte, sig []byte) bool {
	return ed25519.Verify(e.publicKey, data, sig)
}
