package sprint

func StrReverse(s string) string {
    runes := []rune(s) 
    n := len(runes)    

	reversedRunes := make([]rune, n)

    for i := 0; i < n; i++ {
        reversedRunes[i] = runes[n-1-i] 
    }

    return string(reversedRunes) 
}
