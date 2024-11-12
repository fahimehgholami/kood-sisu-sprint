package sprint

import "strings"

func ConvertAnyToAny(nbr, baseFrom, baseTo string) string {
	if len(baseFrom) < 2 || len(baseTo) < 2 {
		return "Invalid base: must have at least 2 unique characters."
	}
	if strings.Contains(baseFrom, "+") || strings.Contains(baseFrom, "-") || strings.Contains(baseTo, "+") || strings.Contains(baseTo, "-") {
		return "Invalid base: base should not contain + or -."
	}

	decimalValue := 0
	for i := 0; i < len(nbr); i++ {
		digit := string(nbr[i])
		index := strings.Index(baseFrom, digit)
		if index == -1 {
			return "Invalid input: nbr contains invalid characters."
		}
		decimalValue = decimalValue*len(baseFrom) + index
	}

	if decimalValue == 0 {
		return string(baseTo[0]) 
	}

	var result []byte
	baseToLength := len(baseTo)

	for decimalValue > 0 {
		remainder := decimalValue % baseToLength
		result = append([]byte{baseTo[remainder]}, result...)
		decimalValue /= baseToLength
	}

	return string(result)
}
