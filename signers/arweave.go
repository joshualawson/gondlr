package signers

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/everFinance/gojwk"
)

type ArweaveSigner struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

func Arweave(jwk []byte) (*ArweaveSigner, error) {
	key, _ := gojwk.Unmarshal(jwk)

	pubKey, err := key.DecodePublicKey()
	if err != nil {
		return nil, err
	}
	pub, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		err = fmt.Errorf("pubKey type error")
		return nil, err
	}

	prvKey, err := key.DecodePrivateKey()
	if err != nil {
		return nil, err
	}
	prv, ok := prvKey.(*rsa.PrivateKey)
	if !ok {
		err = fmt.Errorf("prvKey type error")
		return nil, err
	}

	return &ArweaveSigner{
		publicKey:  pub,
		privateKey: prv,
	}, nil
}

// PublicKey ...
func (a *ArweaveSigner) PublicKey() PublicKey {
	s := sha256.Sum256(a.publicKey.N.Bytes())
	return []byte(base64.RawURLEncoding.EncodeToString(s[:]))
}

// Sign ...
func (a *ArweaveSigner) Sign(data []byte) ([]byte, error) {
	hashed := sha256.Sum256(data)

	return rsa.SignPSS(rand.Reader, a.privateKey, crypto.SHA256, hashed[:], &rsa.PSSOptions{
		SaltLength: rsa.PSSSaltLengthAuto,
		Hash:       crypto.SHA256,
	})
}

// Verify ...
func (a *ArweaveSigner) Verify(data []byte, sig []byte) bool {
	hashed := sha256.Sum256(data)

	if err := rsa.VerifyPSS(a.publicKey, crypto.SHA256, hashed[:], sig, &rsa.PSSOptions{
		SaltLength: rsa.PSSSaltLengthAuto,
		Hash:       crypto.SHA256,
	}); err != nil {
		return false
	}

	return true
}
