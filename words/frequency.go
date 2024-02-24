package words

var englishFreq = map[byte]float64{
	'e': 12.702, 't': 9.056, 'a': 8.167, 'o': 7.507, 'i': 6.966,
	'n': 6.749, 's': 6.327, 'h': 6.094, 'r': 5.987, 'd': 4.253,
	'l': 4.025, 'c': 2.782, 'u': 2.758, 'm': 2.406, 'w': 2.360,
	'f': 2.228, 'g': 2.015, 'y': 1.974, 'p': 1.929, 'b': 1.492,
	'v': 0.978, 'k': 0.772, 'j': 0.153, 'x': 0.150, 'q': 0.095,
	'z': 0.074,
}

func ScoreFrequency(text []byte) float64 {
	// Calculate character frequency distribution of the text
	textFreq, specialCharCount := charFrequency(text)

	// Calculate the score based on the frequency difference
	score := 0.0
	totalChars := len(text)
	for char, freq := range textFreq {
		expectedFreq, exists := englishFreq[char]
		if !exists {
			continue
		}
		score += float64(freq) * (expectedFreq / 100)
	}

	// Normalize the score by the number of characters
	if totalChars > 0 {
		score /= float64(totalChars)
	}

	// Penalise based on number of special characters
	specialCharPenalty := 1.0 - (float64(specialCharCount) / float64(totalChars))

	score *= specialCharPenalty

	return score
}

// Function to calculate the frequency distribution of characters in text
func charFrequency(text []byte) (map[byte]int, int) {
	specialCharCount := 0
	freq := make(map[byte]int)
	for _, char := range text {
		if isLetter(char) {
			char = toLower(char)
			freq[char]++
		} else if char != ' ' {
			specialCharCount++
		}
	}
	return freq, specialCharCount
}

func isLetter(char byte) bool {
	return 65 <= char && char <= 90 || 97 <= char && char <= 122
}

func toLower(char byte) byte {
	switch true {
	case 65 <= char && char <= 90:
		return char + 32
	default:
		return char
	}
}
