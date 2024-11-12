package sprint

func IsSorted(f func(a, b string) int, arr []string) bool {
    n := len(arr)
    if n < 2 {
        return true
    }

    ascending := true
    descending := true

    for i := 1; i < n; i++ {
        if f(arr[i-1], arr[i]) > 0 {
            ascending = false
        }
        if f(arr[i-1], arr[i]) < 0 {
            descending = false
        }
    }

    return ascending || descending
}

func StrCompare(a, b string) int {
    if a < b {
        return -1
    }
    if a > b {
        return 1
    }
    return 0
}
