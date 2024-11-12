package sprint

func ArrMap(f func(int) bool, a []int) []bool {
    results := make([]bool, len(a))
    for i, element := range a {
        results[i] = f(element)
    }
    return results
}

func IsNegative(n int) bool {
    return n < 0
}

func IsPrime(n int) bool {
    if n <= 1 {
        return false
    }
    if n == 2 || n == 3 {
        return true
    }
    if n%2 == 0 || n%3 == 0 {
        return false
    }
    for i := 5; i*i <= n; i += 6 {
        if n%i == 0 || n%(i+2) == 0 {
            return false
        }
    }
    return true
}
