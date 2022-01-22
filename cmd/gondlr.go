package main

import (
	"fmt"
	"github.com/joshualawson/gondlr"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
)

var version = "development"

type config struct {
	host       string
	network    string
	wallet     string
	timeout    int
	noConfirm  bool
	multiplier float64
	batchSize  int
	debug      bool
	indexFile  string
}

type command int

const (
	Balance command = iota
	Withdraw
	Upload
	UploadDir
	Fund
	Price
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	var cfg config

	app := &cli.App{
		Name:    "gondlr",
		Usage:   "bundlr cli built in golang",
		Version: version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "host",
				Value:       "http://node1.bundlr.network",
				Usage:       "Bundlr node hostname/URL",
				Destination: &cfg.host,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "network",
				Aliases:     []string{"n"},
				Value:       "arweave",
				Usage:       "network to use",
				Destination: &cfg.network,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "wallet",
				Aliases:     []string{"w"},
				Usage:       "Path to keyfile or the private key itself",
				Destination: &cfg.wallet,
				Required:    true,
			},
			&cli.IntFlag{
				Name:        "timeout",
				Usage:       "The timeout (in ms) for API HTTP requests - increase if you get timeouts for upload",
				Destination: &cfg.timeout,
			},
			&cli.BoolFlag{
				Name:        "no-confirmation",
				Usage:       "Disable confirmations for certain actions",
				Destination: &cfg.noConfirm,
			},
			&cli.Float64Flag{
				Name:        "multiplier",
				Usage:       "Disable confirmations for certain actions",
				Value:       1.0,
				Destination: &cfg.multiplier,
			},
			&cli.IntFlag{
				Name:        "batch-size",
				Usage:       "Adjust the upload-dir batch size (process more items at once - uses more resources (network, memory, cpu) accordingly!)",
				Destination: &cfg.batchSize,
			},
			&cli.BoolFlag{
				Name:        "debug",
				Aliases:     []string{"d"},
				Usage:       "Increases verbosity of errors and logs additional debug information. Used for troubleshooting.",
				Value:       false,
				Destination: &cfg.debug,
			},
			&cli.StringFlag{
				Name:        "index-file",
				Usage:       "Name of the file to use as an index for upload-dir manifests (relative to the path provided to upload-dir).",
				Destination: &cfg.indexFile,
			},
		},
		Commands: []*cli.Command{
			{
				Name:        "balance",
				Description: "Gets the specified user's balance for the current Bundlr node",
				Usage:       "",
				Action: func(c *cli.Context) error {
					run(cfg, Balance)
					return nil
				},
			},
			{
				Name:        "withdraw",
				Description: "Sends a fund withdrawal request",
				Usage:       "Upload a file",
				Action: func(c *cli.Context) error {
					run(cfg, Withdraw)
					return nil
				},
			},
			{
				Name:        "upload",
				Description: "Uploads a specified file",
				Usage:       "Upload a file",
				Action: func(c *cli.Context) error {
					run(cfg, Upload)
					return nil
				},
			},
			{
				Name:        "upload-dir",
				Description: "Uploads a folder (with a manifest)",
				Usage:       "Upload a file",
				Action: func(c *cli.Context) error {
					run(cfg, UploadDir)
					return nil
				},
			},
			{
				Name:        "fund",
				Description: "Funds your account with the specified amount of atomic units",
				Usage:       "Upload a file",
				Action: func(c *cli.Context) error {
					run(cfg, Fund)
					return nil
				},
			},
			{
				Name:        "price",
				Description: "Check how much of a specific currency is required for an upload of <amount> bytes",
				Usage:       "Upload a file",
				Action: func(c *cli.Context) error {
					run(cfg, Price)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func run(cfg config, cmd command) {
	switch cfg.network {
	case "arweave":
		w, err := os.ReadFile(cfg.wallet)
		if err != nil {
			panic(err)
		}
		cfg.wallet = string(w)
	}

	g, err := gondlr.New(
		gondlr.WithNetwork(cfg.network, cfg.wallet),
		gondlr.WithHost(cfg.host),
		gondlr.WithWallet(cfg.wallet),
	)
	if err != nil {
		panic(err)
	}

	switch cmd {
	case Balance:
		bal, err := g.Balance()
		if err != nil {
			panic(err)
		}

		fmt.Printf("Balance: %s\n", bal.String())
	case Withdraw:
		g.Withdraw()
	case Upload:
		g.Upload()
	case UploadDir:
		g.UploadDir()
	case Fund:
		g.Fund()
	case Price:
		g.Price()
	}
}
