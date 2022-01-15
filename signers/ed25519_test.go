package signers

import (
	"crypto/ed25519"
	"crypto/rand"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEd25519_Sign(t *testing.T) {
	_, pk, err := ed25519.GenerateKey(rand.Reader)
	assert.NoError(t, err)
	signer := Ed25519(string(pk))

	data := []byte("test")

	sig, err := signer.Sign(data)
	assert.NoError(t, err)

	signer.Verify(data, sig)
}

func TestEd25519Signer_PublicKey(t *testing.T) {
	pub, pk, err := ed25519.GenerateKey(rand.Reader)
	assert.NoError(t, err)
	signer := Ed25519(string(pk))

	assert.Equal(t, string(pub), signer.PublicKey().String())
}
