package utils

import (
	"testing"
)

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

	out, _ := FindXorKey([]byte(input))

	if len(out) == 0 {
		t.Fatalf("We failed to do anything")
	}
	// fmt.Println(out)
	// fmt.Println(key)
}

type TC struct {
	input    string
	expected string
}

func TestEncryptRotatingXor(t *testing.T) {
	input := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`
	expected := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	actual := EncryptRotatingXor(input, "ICE")

	if actual != expected {
		t.Errorf("Encryption didn't succeed.\nExpected [%s]\ngot      [%s]", expected, actual)
	}
}
