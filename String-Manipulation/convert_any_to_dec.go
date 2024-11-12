package sprint

import (
	"strings"
)

func ConvertAnyToDec(s string, base string) int {
	if len(base) < 2 || strings.ContainsAny(base, "+-") {
		return 0
	}

	baseMap := make(map[rune]int)
	for i, ch := range base {
		baseMap[ch] = i
	}

	result := 0
	for _, ch := range s { 
		value, exists := baseMap[ch]
		if !exists {
			return 0
		}
		result = result*len(base) + value
	}

	return result
}
