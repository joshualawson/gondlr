package signers

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"math/big"
)

type ArweaveSigner struct {
	privateKey []byte
	publicKey  []byte
}

// Arweave creates an instance of the Arweave signer
func Arweave(privateKey []byte, publicKey []byte) (*ArweaveSigner, error) {

	return &ArweaveSigner{
		privateKey: privateKey,
		publicKey:  publicKey,
	}, nil
}

// Sign calculated the signature of a digest
func (a *ArweaveSigner) Sign(data []byte) ([]byte, error) {
	hashed := sha256.Sum256(data)

	pubk, err := x509.ParsePKCS1PublicKey(a.publicKey)
	if err != nil {
		return nil, err
	}

	pk := &rsa.PrivateKey{
		PublicKey: *pubk,
		D:         &big.Int{},
	}

	pk.D.SetBytes(a.privateKey)

	return rsa.SignPSS(rand.Reader, pk, crypto.SHA256, hashed[:], &rsa.PSSOptions{
		SaltLength: rsa.PSSSaltLengthAuto,
		Hash:       crypto.SHA256,
	})
}

// Verify verifies a signature
func (a *ArweaveSigner) Verify(data []byte, sig []byte) bool {
	hashed := sha256.Sum256(data)

	pk, err := x509.ParsePKCS1PublicKey(a.publicKey)
	if err != nil {
		return false
	}

	if err := rsa.VerifyPSS(pk, crypto.SHA256, hashed[:], sig, &rsa.PSSOptions{
		SaltLength: rsa.PSSSaltLengthAuto,
		Hash:       crypto.SHA256,
	}); err != nil {
		return false
	}

	return true
}
