package sprint

func BalanceOut(arr []bool) []bool {
    trueCount := 0
    falseCount := 0

    for _, value := range arr {
        if value {
            trueCount++
        } else {
            falseCount++
        }
    }

    for trueCount < falseCount {
        arr = append(arr, true)
        trueCount++
    }

    for falseCount < trueCount {
        arr = append(arr, false)
        falseCount++
    }

    return arr 
}
