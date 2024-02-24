package utils

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBase64(h string) (string, error) {
	binary, err := hex.DecodeString(h)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(binary), nil
}
