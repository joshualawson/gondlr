package signers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPublicKey_String(t *testing.T) {
	var pk PublicKey = []byte("*publicKey*")

	assert.Equal(t, "*publicKey*", pk.String())
}

func TestPublicKey_Byte(t *testing.T) {
	var pk PublicKey = []byte("*publicKey*")

	assert.Equal(t, []byte("*publicKey*"), pk.Byte())
}

func TestPrivateKey_String(t *testing.T) {
	var pk PrivateKey = []byte("*privateKey*")

	assert.Equal(t, "*privateKey*", pk.String())
}

func TestPrivateKey_Byte(t *testing.T) {
	var pk PublicKey = []byte("*privateKey*")

	assert.Equal(t, []byte("*privateKey*"), pk.Byte())
}
