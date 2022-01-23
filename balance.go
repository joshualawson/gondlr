package gondlr

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/big"
	"net/http"
)

type BalanceResponse struct {
	balance big.Int `json:"balance"`
}

func (g *Gondlr) Balance() (*big.Int, error) {
	log.
		WithFields(log.Fields{
			"host":      g.host.String(),
			"network":   g.network.Name(),
			"publicKey": g.wallet.PublicKey(),
		}).
		Debugf("GET %saccount/balance/%s?address=%s", g.host.String(), g.network.Name(), g.wallet.PublicKey())

	req, err := http.NewRequest("GET", fmt.Sprintf("%saccount/balance/%s?address=%s", g.host.String(), g.network.Name(), g.wallet.PublicKey()), nil)
	if err != nil {
		return nil, err
	}

	res, err := g.httpClient.Do(req)
	if res.StatusCode != http.StatusOK {
		return nil, ErrorHTTP(http.StatusOK, res.StatusCode)
	}
	body, err := ioutil.ReadAll(res.Body)

	var bal BalanceResponse
	if err := json.Unmarshal(body, &bal); err != nil {
		return nil, ErrorJSONUnmarshal(err)
	}

	log.
		WithFields(log.Fields{
			"host":      g.host.String(),
			"network":   g.network.Name(),
			"publicKey": g.wallet.PublicKey(),
			"balance":   bal.balance.String(),
		}).
		Debugf("response body: %v", string(body))

	return &bal.balance, nil
}
