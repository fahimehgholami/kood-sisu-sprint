package sprint

// IntVsFloat compares an integer and a float, returning a string based on the comparison.
func IntVsFloat(i int, f float32) string {
	if float32(i) > f {
		return "Integer"
	} else if float32(i) < f {
		return "Float"
	} else {
		return "Same"
	}
}
//int_vs_float.go
	/*
	   fmt.Println(sprint.IntVsFloat(5, 8.8))   // Expected output: "Float"
	   fmt.Println(sprint.IntVsFloat(10, 10.0)) // Expected output: "Same"
	   fmt.Println(sprint.IntVsFloat(15, 12.3)) // Expected output: "Integer"
	*/
