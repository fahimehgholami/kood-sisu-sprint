package sprint

func Countdown(n int) string {
	// Start with an empty result string
	result := ""
	for i := n; i > 0; i -= 2 {
		// Add the current number followed by ", " if there are more numbers in the sequence
		result += string(rune('0' + i)) + ", "
	}
	result += "0!" // Add the final "0!" to the result
	return result
}

//	fmt.Println(sprint.Countdown(7)) 
