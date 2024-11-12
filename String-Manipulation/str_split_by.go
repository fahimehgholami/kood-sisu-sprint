package sprint
//test
func StrSplitBy(s, sep string) []string {
	if s == "" {
		return []string{} 
	}
	if sep == "" {
		return []string{s} 
	}

	var result []string
	start := 0
	for {
		index := indexOf(s, sep, start)
		if index == -1 {
			result = append(result, s[start:])
			break
		}
		result = append(result, s[start:index])
		start = index + len(sep)
	}

	return result
}

func indexOf(s, substr string, start int) int {
	if start < 0 || start >= len(s) {
		return -1
	}

	substrLen := len(substr)
	for i := start; i <= len(s)-substrLen; i++ {
		if s[i:i+substrLen] == substr {
			return i
		}
	}

	return -1
}
