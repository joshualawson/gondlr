package gondlr

import "net/http"

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Gondlr struct {
	httpClient HTTPClient
	signer     Signer
}

func New(opts ...Option) *Gondlr {
	g := &Gondlr{
		httpClient: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(g)
	}

	return g
}
