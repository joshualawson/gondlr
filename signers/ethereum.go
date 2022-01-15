package signers

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"log"
)

type EthereumSigner struct {
	privateKey *ecdsa.PrivateKey
}

func Ethereum(privateKey string) (*EthereumSigner, error) {
	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}

	return &EthereumSigner{
		privateKey: pk,
	}, nil
}

// PublicKey ...
func (e *EthereumSigner) PublicKey() PublicKey {
	publicKey := e.privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])

	return []byte(hexutil.Encode(hash.Sum(nil)[12:]))
}

// Sign ...
func (e *EthereumSigner) Sign(data []byte) ([]byte, error) {
	return crypto.Sign(data, e.privateKey)
}

// Verify ...
func (e *EthereumSigner) Verify(data []byte, sig []byte) bool {
	return crypto.VerifySignature(e.PublicKey(), data, sig)
}
