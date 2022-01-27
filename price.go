package gondlr

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/big"
	"net/http"
)

func (g *Gondlr) Price(bytes int) (*big.Int, error) {
	log.
		WithFields(log.Fields{
			"host":      g.host.String(),
			"network":   g.network.Name(),
			"publicKey": g.wallet.PublicKey(),
		}).
		Debugf("GET %sprice/%s/%n", g.host.String(), g.network.Name(), bytes)

	req, err := http.NewRequest("GET", fmt.Sprintf("%sprice/%s/%v", g.host.String(), g.network.Name(), bytes), nil)
	if err != nil {
		return nil, err
	}

	res, err := g.httpClient.Do(req)
	if res.StatusCode != http.StatusOK {
		return nil, ErrorHTTP(http.StatusOK, res.StatusCode)
	}
	body, err := ioutil.ReadAll(res.Body)

	p := new(big.Int)
	if err := p.UnmarshalText(body); err != nil {
		return nil, ErrorUnmarshalTextToBigInt(err)
	}

	log.
		WithFields(log.Fields{
			"host":    g.host.String(),
			"network": g.network.Name(),
			"bytes":   bytes,
			"price":   p.String(),
		}).
		Debugf("response body: %v", string(body))

	return p, nil
}
