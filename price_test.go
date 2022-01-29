package gondlr

import (
	"bytes"
	"github.com/joshualawson/gondlr/wallet"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGondlr_Price(t *testing.T) {
	c := &MockHTTPClient{}
	c.On("Do", mock.AnythingOfType("*http.Request")).
		Return(&http.Response{StatusCode: http.StatusOK, Body: ioutil.NopCloser(bytes.NewReader([]byte("3574332")))}, nil)

	w := &MockWallet{}
	w.On("PublicKey").Return(wallet.PublicKey("*publicKey*"))

	n := &MockNetwork{}
	n.On("Name").Return("arweave")

	g, err := New(
		WithClient(c),
		WithHost("http://localhost/"),
		WithWallet(w), WithNetwork(n),
	)
	assert.NoError(t, err)

	p, err := g.Price("1000")
	assert.NoError(t, err)

	assert.Equal(t, "3574332", p.String())
}
