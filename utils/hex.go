package utils

import (
	"encoding/base64"
	"encoding/hex"

	"github.com/harryalaw/cryptopals/words"
)

func HexToBase64(h string) (string, error) {
	binary, err := hex.DecodeString(h)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(binary), nil
}

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

func SingleXor(input string, xor byte) (string, error) {
	binaryInput, err := hex.DecodeString(input)

	if err != nil {
		return "", err
	}

	out := make([]byte, len(binaryInput))

	for i, val := range binaryInput {
		out[i] = val ^ xor
	}

	return string(out), nil
}

// Given an input string that has been XOR'd with a single character
// identify and return the original string and the character it had been
// XOR'd with
func FindXorKey(input string) (string, byte) {
    bestScore := -1.0;
    var xorKey byte = 0;
    text := "";
	var i byte
	for i = 0; i < 255; i++ {
		decoded, err := SingleXor(input, i)
		if err != nil {
			continue
		}

        score := words.ScoreFrequency(decoded)

        if score > bestScore {
            xorKey = i;
            text = decoded;
            bestScore = score
        }
	}

    return text, xorKey
}
