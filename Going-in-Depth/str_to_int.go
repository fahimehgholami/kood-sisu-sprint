package sprint

func StrToInt(s string) int {
	result := 0     
	sign := 1        
	i := 0           

	for i < len(s) && (s[i] == ' ' || s[i] == '\t') {
		i++
	}

	if i < len(s) {
		if s[i] == '+' {
			sign = 1  
			i++
		} else if s[i] == '-' {
			sign = -1 
			i++
		}
	}

	for i < len(s) {
		if s[i] < '0' || s[i] > '9' { 
			return 0 
		}

		result = result*10 + int(s[i]-'0')
		i++
	}

	return result * sign
}
