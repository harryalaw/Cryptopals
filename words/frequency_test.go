package words

import "testing"

func TestIsLower(t *testing.T) {
    var i byte = 0;
    caps := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ");

    for i = 0; i < 255; i++ {
        lower := toLower(i);

        if i != lower && !includes(caps, i) {
            t.Errorf("Expected byte to not change, got=%b, expected=%b", lower, i)
        }
    }

}

func TestIsLetter(t *testing.T) {
    var i byte = 0
    letters := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

    for i = 0; i < 255; i++ {
        letter := isLetter(i);

        if letter && !includes(letters, i) {
            t.Errorf("Expected byte %b (%c) to not be a letter", i, i)
        }
    }
}

func includes(arr []byte, x byte) bool {
    for _, val := range arr {
        if val == x {
            return true
        }
    }
    return false
}
