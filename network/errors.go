package network

import "fmt"

var (
	ErrorSDK = func(err error) error {
		return fmt.Errorf("SDK Error: %w", err)
	}
)
