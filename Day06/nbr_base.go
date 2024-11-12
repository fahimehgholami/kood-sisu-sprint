package sprint

func NbrBase(n int, base string) string {
	if !isValidBase(base) {
		return "NV" 
	}

	baseLen := len(base)
	result := ""
	isNegative := n < 0

	if isNegative {
		n = -n
	}

	for n > 0 {
		remainder := n % baseLen
		result = string(base[remainder]) + result
		n /= baseLen
	}

	if result == "" {
		result = string(base[0])
	}

	if isNegative {
		result = "-" + result
	}

	return result
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	seen := make(map[rune]bool)
	for _, char := range base {
		if char == '+' || char == '-' {
			return false
		}
		if seen[char] {
			return false 
		}
		seen[char] = true
	}

	return true
}
