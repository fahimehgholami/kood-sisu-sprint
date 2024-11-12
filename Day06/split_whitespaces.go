package sprint

func SplitWhitespaces(s string) []string {
	var words []string
	var word string

	for i := 0; i < len(s); i++ {
		char := s[i]
		if char == ' ' || char == '\t' || char == '\n' {
			if word != "" {
				words = append(words, word)
				word = "" 
			}
		} else {
			word += string(char)
		}
	}

	if word != "" {
		words = append(words, word)
	}

	return words
}