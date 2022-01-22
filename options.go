package gondlr

import (
	"github.com/joshualawson/gondlr/network"
	"github.com/joshualawson/gondlr/signers"
	"net/url"
)

type Option func(g *Gondlr) error

// WithClient allows the user to implement their own HTTP client and supply it to Gondlr
func WithClient(client HTTPClient) Option {
	return func(g *Gondlr) error {
		g.httpClient = client

		return nil
	}
}

func WithNetwork(nw string, privateKey string) Option {
	return func(g *Gondlr) error {
		switch nw {
		case "ethereum":
		case "arweave":
			signer, err := signers.Arweave([]byte(privateKey))
			if err != nil {
				panic(err)
			}
			g.network = network.Arweave(signer)
		case "solana":
		}

		return nil
	}
}

func WithHost(host string) Option {
	return func(g *Gondlr) error {
		u, err := url.Parse(host)
		if err != nil {
			return err
		}

		g.host = u

		return nil
	}
}

func WithWallet(wallet string) Option {
	return func(g *Gondlr) error {
		g.wallet = wallet
		return nil
	}
}
