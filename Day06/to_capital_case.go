package sprint

func ToCapitalCase(s string) string {
	result := []rune(s) 
	inWord := false     

	for i := 0; i < len(result); i++ {
		if isAlphanumeric(result[i]) { 
			if !inWord {
				result[i] = toUpper(result[i]) 
				inWord = true
			} else {
				result[i] = toLower(result[i]) 
			}
		} else {
			inWord = false 
		}
	}

	return string(result) 
}

func isAlphanumeric(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9')
}

func toUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - ('a' - 'A')
	}
	return r
}

func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + ('a' - 'A')
	}
	return r
}
