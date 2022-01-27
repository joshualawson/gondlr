package signers

import (
	"crypto/ed25519"
	"crypto/rand"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEd25519_Sign(t *testing.T) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	assert.NoError(t, err)
	signer := Ed25519(priv, pub)

	data := []byte("test")

	sig, err := signer.Sign(data)
	assert.NoError(t, err)

	signer.Verify(data, sig)
}
