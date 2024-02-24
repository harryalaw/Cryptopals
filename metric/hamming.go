package metric

func Hamming(a, b string) int {
	aBytes := []byte(a)
	bBytes := []byte(b)

	return HammingBytes(aBytes, bBytes)
}

func HammingBytes(a, b []byte) int {
	byteDiffs := 0

	for i := range a {
		xor := a[i] ^ b[i]
		for xor > 0 {
			byteDiffs += int(xor & 1)
			xor >>= 1
		}
	}

	return byteDiffs

}
