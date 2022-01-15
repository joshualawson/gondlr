package signers

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/sha3"
	"log"
	"testing"
)

func TestEthereum_Sign(t *testing.T) {
	pk, err := crypto.GenerateKey()
	privateKeyBytes := crypto.FromECDSA(pk)

	assert.NoError(t, err)
	signer, err := Ethereum(hexutil.Encode(privateKeyBytes)[2:])
	assert.NoError(t, err)

	data := []byte("test1231231231231231231231231231")

	sig, err := signer.Sign(data)
	assert.NoError(t, err)

	signer.Verify(data, sig)
}

func TestEthereum_PublicKey(t *testing.T) {
	// Keys were generated on website for test -- DO NOT USE THESE KEYS FOR ANYTHING!
	signer, err := Ethereum("654397ed69104794e69cb4ff4692a6c40e554cf570bd71223bc0cc45dfe8a909")
	assert.NoError(t, err)

	assert.Equal(t, []byte("0xf60e51cb6ae11ee9981142f6b3de07a43efc8186"), signer.PublicKey())
}

func getEthereumPublicKey(privateKey *ecdsa.PrivateKey) []byte {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	return []byte(hexutil.Encode(hash.Sum(nil)[12:]))
}
