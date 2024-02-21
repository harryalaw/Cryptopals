package utils

import (
	"fmt"
	"testing"
)

func TestHexToBase64(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	output, err := HexToBase64(input)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if output != expected {
		t.Fatalf("\nExpected: %s\nGot     : %s", expected, output)
	}
}

func TestFixedXor(t *testing.T) {
	input := "1c0111001f010100061a024b53535009181c"
	xor := "686974207468652062756c6c277320657965"

	expected := "746865206b696420646f6e277420706c6179"

	output, err := FixedXor(input, xor)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if output != expected {
		t.Fatalf("\nExpected: %s\nGot     : %s", expected, output)
	}

}

func TestFindXorKey(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

    out, key := FindXorKey(input)

    if out == "" {
        t.Fatalf("We failed to do anything")
    }
    fmt.Println(out)
    fmt.Println(key)
}
