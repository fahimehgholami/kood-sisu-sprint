package sprint

func GenerateRange(min, max int) []int {
    if min >= max {
        return nil 
    }

    rangeSlice := make([]int, max-min)

    for i := min; i < max; i++ {
        rangeSlice[i-min] = i 
    }

    return rangeSlice 
}
