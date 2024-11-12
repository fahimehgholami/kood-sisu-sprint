package sprint

func AlphabetMastery(n int) string {
	//initialize empty resdualt
	result := ""
	//loop from 0 > n 

	for i := 0; i < n; i++ {
		//from a to result
		result += string('a' + rune(i))
	}
	return result
}
/*
func main() {
    result := sprint.AlphabetMastery(6)
    fmt.Println(result) // Output: "abcdef"
}
*/