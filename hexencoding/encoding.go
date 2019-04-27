package hexencoding

import (
	"encoding/base64"
	"encoding/hex"
)

// HexToBase64 converts a hex input to a base64 output
func HexToBase64(input string) (string, error) {
	b, err := hex.DecodeString(input)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}
