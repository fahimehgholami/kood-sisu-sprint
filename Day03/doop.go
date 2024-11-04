package sprint

// Doop performs a mathematical operation based on the given operator.
func Doop(a int, op string, b int) int {
	// Switch statement to handle different operations
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		// Handle division by zero
		if b == 0 {
			return 0
		}
		return a / b
	case "%":
		// Handle modulo by zero
		if b == 0 {
			return 0
		}
		return a % b
	default:
		// Return 0 for invalid operator
		return 0
	}
}
//doop.go:
	/*
	fmt.Println(sprint.Doop(5, "+", 3))  // Expected output: 8
    fmt.Println(sprint.Doop(8, "/", 0))  // Expected output: 0
    fmt.Println(sprint.Doop(8, "-", 2))  // Expected output: 6
    fmt.Println(sprint.Doop(5, "*", 4))  // Expected output: 20
    fmt.Println(sprint.Doop(10, "%", 3)) // Expected output: 1
    fmt.Println(sprint.Doop(10, "invalid", 3)) // Expected output: 0
	*/