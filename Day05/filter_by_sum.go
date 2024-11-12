package sprint

func FilterBySum(arr [][]int, limit int) [][]int {
    result := [][]int{} 

    for _, subArray := range arr {
        sum := 0

        for _, num := range subArray {
            sum += num
        }

        if sum >= limit {
            result = append(result, subArray) 
        }
    }

    return result 
}
