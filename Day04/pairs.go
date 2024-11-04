package sprint

import (
	"fmt"
)

// Pairs generates all possible combinations of two two-digit numbers.
func Pairs() string {
	var result string // Initialize an empty string to build the output

	for i := 0; i < 100; i++ { // Outer loop for the first number
		for j := i + 1; j < 100; j++ { // Inner loop for the second number
			// Format numbers with leading zeros and space between them
			pair := fmt.Sprintf("%02d %02d", i, j)
			
			if result != "" { // Check if result is not empty
				result += ", " // Append comma and space before new pair
			}
			result += pair // Append the formatted pair to the result
		}
	}

	return result // Return the constructed string
}
/*
	fmt.Println(sprint.Pairs())
*/