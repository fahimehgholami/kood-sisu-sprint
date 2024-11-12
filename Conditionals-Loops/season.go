package sprint

// Season returns the season for a given month name or an invalid input message.
func Season(month string) string {
	switch month {
	case "jan", "Jan", "feb", "Feb", "dec", "Dec":
		return "winter"
	case "mar", "Mar", "apr", "Apr", "may", "May":
		return "spring"
	case "jun", "Jun", "jul", "Jul", "aug", "Aug":
		return "summer"
	case "sep", "Sep", "oct", "Oct", "nov", "Nov":
		return "autumn"
	default:
		return "invalid input: " + month
	}
}

//season.go
	// Test cases
	/*
	   fmt.Println(sprint.Season("feb"))        // Expected output: "winter"
	   fmt.Println(sprint.Season("September"))   // Expected output: "invalid input: September"
	   fmt.Println(sprint.Season("jun"))        // Expected output: "summer"
	   fmt.Println(sprint.Season("Apr"))        // Expected output: "spring"
	   fmt.Println(sprint.Season("hello"))      // Expected output: "invalid input: hello"
	*/