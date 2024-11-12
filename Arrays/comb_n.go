package sprint

import "fmt"

func CombN(n int) []string {
    var result []string

    var generateComb func(current string, start int)
    generateComb = func(current string, start int) {
        if len(current) == n {
            result = append(result, current)
            return
        }
        
        for i := start; i <= 9; i++ {
            generateComb(current+fmt.Sprint(i), i+1) 
        }
    }

    generateComb("", 0)
    return result
}
