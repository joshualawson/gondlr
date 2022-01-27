package gondlr

import (
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

func WithNetwork(n Network) Option {
	return func(g *Gondlr) error {
		g.network = n
		return nil
	}
}

func WithWallet(w Wallet) Option {
	return func(g *Gondlr) error {
		g.wallet = w
		return nil
	}
}

func WithSigner(s Signer) Option {
	return func(g *Gondlr) error {
		g.signer = s
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
