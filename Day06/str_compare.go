package sprint

func StrCompare(a, b string) int {
	if a == b {
		return 0 
	} else if a < b {
		return -1 
	} else {
		return 1 
	}
}