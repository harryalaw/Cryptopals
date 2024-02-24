package metric

import "testing"

func TestHammingDistance(t *testing.T) {

	input1 := "this is a test"
	input2 := "wokka wokka!!!"

	distance := Hamming(input1, input2)

	if distance != 37 {
		t.Fatalf("Distance not calculated correctly, expected=%d, got=%d", 37, distance)
	}
}
