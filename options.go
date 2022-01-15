package gondlr

type Option func(g *Gondlr)

// WithClient allows the user to implement their own HTTP client and supply it to Gondlr
func WithClient(client HTTPClient) Option {
	return func(g *Gondlr) {
		g.httpClient = client
	}
}

func WithSigner(signer string) Option {
	return func(g *Gondlr) {
		switch signer {
		case "ethereum":
		case "arweave":
		case "solana":
		}
	}
}
