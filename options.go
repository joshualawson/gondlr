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

func WithNetwork(nw string, wallet Wallet) Option {
	return func(g *Gondlr) error {
		switch nw {
		case "ethereum":
		case "arweave":
			s, err := signers.Arweave(wallet.PrivateKeyBytes(), wallet.PublicKeyBytes())
			if err != nil {
				panic(err)
			}
			g.network = network.Arweave()
			g.signer = s
			g.wallet = wallet
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
