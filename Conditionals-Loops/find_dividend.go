package sprint

// FindDividend returns the first number in the range [from, to) that is divisible by divisor.
// If no such number exists, it returns -1.
func FindDividend(from, to, divisor int) int {
    for i := from; i < to; i++ {
        if i%divisor == 0 {
            return i // Return the first divisible number found
        }
    }
    return -1 // Return -1 if no number is found
}
	//find_dividend.go:
	/*
	fmt.Println(sprint.FindDividend(5, 17, 4))  // Expected output: 8
    fmt.Println(sprint.FindDividend(10, 18, 9)) // Expected output: -1
    fmt.Println(sprint.FindDividend(1, 10, 3))  // Expected output: 3
    fmt.Println(sprint.FindDividend(15, 20, 5)) // Expected output: 15
    fmt.Println(sprint.FindDividend(20, 25, 6)) // Expected output: -1
	*/