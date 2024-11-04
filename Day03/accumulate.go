package sprint

// Accumulate returns the sum of all digits from 0 to n. If n is negative, it returns 0.
func Accumulate(n int) int {
	// Check if n is negative
	if n < 0 {
		return 0
	}

	// Initialize sum
	sum := 0

	// Loop through numbers from 0 to n
	for i := 0; i <= n; i++ {
		sum += i // Add i to the sum
	}

	return sum // Return the calculated sum
}
//accumulate.go:
	// Test cases
	/*
	   fmt.Println(sprint.Accumulate(4))  // Expected output: 10
	   fmt.Println(sprint.Accumulate(-1)) // Expected output: 0
	   fmt.Println(sprint.Accumulate(0))  // Expected output: 0
	   fmt.Println(sprint.Accumulate(5))  // Expected output: 15
	*/