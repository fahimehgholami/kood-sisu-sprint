package sprint

func ReverseAlphabet(step int) string {
	//set defult step to 1 if step is 0 or negative
	if step <= 0 {
		step = 1
	}
	result := ""
	for r := 'z'; r >= 'a'; r -= rune(step){
		result += string(r)
	}
	return result
}
/*
   fmt.Println("Output with step 5:", sprint.ReverseAlphabet(5))
	fmt.Println("Output with step 3:", sprint.ReverseAlphabet(3))
	fmt.Println("Output with step 0:", sprint.ReverseAlphabet(0)) // Tests default step of 1
	fmt.Println("Output with step -2:", sprint.ReverseAlphabet(-2)) // Tests default step of 1
}*/