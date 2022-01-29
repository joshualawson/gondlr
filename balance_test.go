package gondlr

import (
	"bytes"
	"github.com/joshualawson/gondlr/wallet"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"math/big"
	"net/http"
	"testing"
)

func TestGondlr_Balance(t *testing.T) {
	c := &MockHTTPClient{}
	c.On("Do", mock.AnythingOfType("*http.Request")).
		Return(&http.Response{StatusCode: http.StatusOK, Body: ioutil.NopCloser(bytes.NewReader([]byte(`{"balance":"100"}`)))}, nil)

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

	bal, err := g.Balance()
	assert.NoError(t, err)

	c.AssertCalled(t, "Do", mock.AnythingOfType("*http.Request"))
	assert.Equal(t, big.NewInt(100), bal)
}
