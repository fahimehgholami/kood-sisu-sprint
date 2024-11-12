package sprint

// CountDivisible counts how many numbers in the range [from, to) that are divisible by divisor
// and checks for exceptions based on the input parameters.
func CountDivisible(from, to, step, divisor int) int {
    // Check for exceptions
    if step <= 0 || divisor == 0 {
        return 0
    }

    count := 0
    for i := from; i < to; i += step {
        if i%divisor == 0 {
            count++
        }
    }
    
    return count
}
	//count_divisible.go:
	/*
	fmt.Println(sprint.CountDivisible(5, 17, 2, 3)) // Expected output: 2
    fmt.Println(sprint.CountDivisible(5, 20, 3, 4)) // Expected output: 1
    fmt.Println(sprint.CountDivisible(10, 50, 5, 10)) // Expected output: 4
    fmt.Println(sprint.CountDivisible(5, 5, 1, 1)) // Expected output: 0 (empty range)
    fmt.Println(sprint.CountDivisible(5, 17, 0, 3)) // Expected output: 0 (invalid step)
    fmt.Println(sprint.CountDivisible(5, 17, 2, 0)) // Expected output: 0 (invalid divisor)
	*/