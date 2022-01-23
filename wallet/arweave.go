package wallet

import (
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"github.com/everFinance/gojwk"
)

type ArweaveWallet struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

func Arweave(jwk []byte) (*ArweaveWallet, error) {
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

	return &ArweaveWallet{
		publicKey:  pub,
		privateKey: prv,
	}, nil
}

// PrivateKeyBytes ...
func (a *ArweaveWallet) PrivateKeyBytes() []byte {
	return a.privateKey.D.Bytes()
}

// PublicKeyBytes ...
func (a *ArweaveWallet) PublicKeyBytes() []byte {
	return x509.MarshalPKCS1PublicKey(a.publicKey)
}

// PublicKey ...
func (a *ArweaveWallet) PublicKey() PublicKey {
	s := sha256.Sum256(a.publicKey.N.Bytes())
	return []byte(base64.RawURLEncoding.EncodeToString(s[:]))
}

// PrivateKey ...
func (a *ArweaveWallet) PrivateKey() PrivateKey {
	s := sha256.Sum256(a.PrivateKeyBytes())
	return []byte(base64.RawURLEncoding.EncodeToString(s[:]))
}
