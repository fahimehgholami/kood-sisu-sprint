package sprint

func Sqrt(n int) int {
    if n < 0 {
        return 0
    }
    low, high := 1, n
    for low <= high {
        mid := (low + high) / 2
        square := mid * mid
        if square == n {
            return mid
        } else if square < n {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return 0
}
