package sprint

import (
	"fmt"
)

func Combinations() string {
	var result string // Initialize an empty string 

	// first digit of the triplet
	for i := 0; i <= 7; i++ { // i can go from 0 to 7
		//second digit, greater than the first
		for j := i + 1; j <= 8; j++ { // j can go from i+1 to 8
			//  third digit, greater than the second
			for k := j + 1; k <= 9; k++ { // k can go from j+1 to 9
				// Format the triplet
				triplet := fmt.Sprintf("%d%d%d", i, j, k)

				if result != "" { // result is not empty
					result += ", " 
				}
				result += triplet 
			}
		}
	}

	return result 
}

/*
	fmt.Println(sprint.Combinations())
*/