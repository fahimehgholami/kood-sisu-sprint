package sprint

import (
    "strconv"
    "strings"
)

func EightQueensSolver() string {
    var results []string
    solve([]int{}, &results)
    return strings.Join(results, "\n")
}

// Recursive function to place queens on the board
func solve(queens []int, results *[]string) {
    n := 8
    row := len(queens)

    if row == n {
        *results = append(*results, formatSolution(queens))
        return
    }

    for col := 1; col <= n; col++ {
        if isValid(queens, row, col) {
            solve(append(queens, col), results)
        }
    }
}

//  check if placing a queen at (row, col) is safe
func isValid(queens []int, row, col int) bool {
    for r, c := range queens {
        if c == col || r+c == row+col || r-c == row-col {
            return false
        }
    }
    return true
}

//  format the solution as required by the problem
func formatSolution(queens []int) string {
    var solution strings.Builder
    for _, col := range queens {
        solution.WriteString(strconv.Itoa(col))
    }
    return solution.String()
}
