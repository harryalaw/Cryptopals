package utils

import (
	"encoding/hex"
	"fmt"

	"github.com/harryalaw/cryptopals/words"
)

func FixedXor(input string, xor string) (string, error) {
	binaryInput, err := hex.DecodeString(input)

	if err != nil {
		return "", err
	}

	xorValue, err := hex.DecodeString(xor)

	if err != nil {
		return "", err
	}

	out := make([]byte, len(binaryInput))

	for i, val := range binaryInput {
		out[i] = val ^ xorValue[i]
	}

	return hex.EncodeToString(out), nil
}

func SingleXorHex(input string, xor byte) (string, error) {
	binaryInput, err := hex.DecodeString(input)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(SingleXor(binaryInput, xor)), nil
}

func SingleXor(input []byte, xor byte) []byte {
	out := make([]byte, len(input))

	for i, val := range input {
		out[i] = val ^ xor
	}

	return out
}

// Given an input string that has been XOR'd with a single character
// identify and return the original bytes and the character it had been
// XOR'd with
func FindXorKey(input []byte) ([]byte, byte) {
	bestScore := -1.0
	var xorKey byte = 0
	var text []byte
	var i byte

    for i = 1; i < 255; i++ {
		decoded := SingleXor(input, i)

		score := words.ScoreFrequency(decoded)

		if score > bestScore {
			xorKey = i
			text = decoded
			bestScore = score
		}
	}

	return text, xorKey
}

func EncryptRotatingXor(input string, key string) string {
	inputBytes := []byte(input)
	out := make([]byte, len(inputBytes))

	keyBytes := []byte(key)

	for i, val := range inputBytes {
		out[i] = val ^ keyBytes[i%len(keyBytes)]
	}

	return hex.EncodeToString(out)
}
