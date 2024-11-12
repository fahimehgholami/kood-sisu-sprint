package sprint

func intSqrt(n int) int {
    if n < 2 {
        return n
    }
    low, high := 1, n
    for low <= high {
        mid := (low + high) / 2
        square := mid * mid
        if square == n {
            return mid
        }
        if square < n {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return high
}

func IsPrime(n int) bool {
    if n <= 1 {
        return false
    }
    if n == 2 {
        return true
    }
    if n%2 == 0 {
        return false
    }

    limit := intSqrt(n)
    for i := 3; i <= limit; i += 2 {
        if n%i == 0 {
            return false
        }
    }

    return true
}
