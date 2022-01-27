package gondlr

import (
	"fmt"
	"net/http"
)

var (
	ErrorHTTP = func(expected int, received int) error {
		return fmt.Errorf("request returned [ %s ] not [ %s ]", http.StatusText(received), http.StatusText(expected))
	}

	ErrorJSONUnmarshal = func(err error) error {
		return fmt.Errorf("unable to unmarshal json: %w", err)
	}

	ErrorUnmarshalTextToBigInt = func(e error) error {
		return fmt.Errorf("unable to unmarshal text to big.Int: %w", e)
	}
)
