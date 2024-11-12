package sprint

func AlphaNumber(n int) string {
	sign := ""
	if n < 0 {
		sign = "-"
		n = -n 
	}

	alphaNum := ""
	for n > 0 {
		digit := n % 10                  
		alphaNum = string('a'+digit) + alphaNum 
		n /= 10                           
	}

	if alphaNum == "" {
		alphaNum = "a"
	}

	return sign + alphaNum
}
//    fmt.Println(sprint.AlphaNumber(-1280)) //  "-bcia"
