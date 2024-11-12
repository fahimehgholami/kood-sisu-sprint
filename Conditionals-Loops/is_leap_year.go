package sprint

// IsLeapYear checks if the given year is a leap year.
func IsLeapYear(year int) bool {
	// A year is a leap year if:
	// 1. It is divisible by 4
	// 2. It is not divisible by 100 unless it is also divisible by 400
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

//is_leap_year.go:
/*
	fmt.Println(sprint.IsLeapYear(2020)) // Expected output: true
    fmt.Println(sprint.IsLeapYear(2021)) // Expected output: false
    fmt.Println(sprint.IsLeapYear(1900)) // Expected output: false
    fmt.Println(sprint.IsLeapYear(2000)) // Expected output: true
    fmt.Println(sprint.IsLeapYear(2024)) // Expected output: true
*/
