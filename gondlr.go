package gondlr

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Gondlr struct {
	httpClient HTTPClient
	host       *url.URL
	network    Network
	wallet     Wallet
	signer     Signer
}

func New(opts ...Option) (*Gondlr, error) {
	g := &Gondlr{
		httpClient: http.DefaultClient,
	}

	for _, opt := range opts {
		if err := opt(g); err != nil {
			return nil, fmt.Errorf("gondlr failed to initilaize: %w", err)
		}
	}

	return g, nil
}

func (g *Gondlr) hostUrl() string {
	return fmt.Sprintf(g.host.String())
}

func (g *Gondlr) Network() string {
	return strings.Title(g.network.Name())
}

func (g *Gondlr) Withdraw() {
	fmt.Println("unimplemented")

	os.Exit(0)
}
func (g *Gondlr) Upload() {
	fmt.Println("unimplemented")

	os.Exit(0)
}
func (g *Gondlr) UploadDir() {
	fmt.Println("unimplemented")

	os.Exit(0)
}
func (g *Gondlr) Fund() {
	fmt.Println("unimplemented")

	os.Exit(0)
}
