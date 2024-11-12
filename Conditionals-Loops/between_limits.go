package sprint

// BetweenLimits returns a string containing all runes between two input runes in alphabetical order.
func BetweenLimits(from, to rune) string {
	// Initialize an empty string to hold the result
	result := ""

	// Determine the start and end runes
	start := from
	end := to

	// Swap if 'from' is greater than 'to'
	if from > to {
		start, end = to, from
	}

	// Loop through the range between start and end
	for r := start + 1; r < end; r++ {
		result += string(r) // Append the rune to the result string
	}

	return result // Return the final string
}
	//between_limits.go:
	/*
		fmt.Println(sprint.BetweenLimits('f', 'j')) // Expected output: "ghi"
	    fmt.Println(sprint.BetweenLimits('j', 'f')) // Expected output: "ghi"
	    fmt.Println(sprint.BetweenLimits('a', 'd')) // Expected output: "bc"
	    fmt.Println(sprint.BetweenLimits('z', 'v')) // Expected output: "xy"
	*/