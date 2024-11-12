package sprint

// Counts the number of elements in the slice that satisfy the given function f
func ArrCountIf(f func(string) bool, a []string) int {
    count := 0
    for _, element := range a {
        if f(element) {
            count++
        }
    }
    return count
}

// Checks if the string contains only uppercase letters
func IsUpper(s string) bool {
    if len(s) == 0 {
        return false
    }
    for _, r := range s {
        if r < 'A' || r > 'Z' {
            return false
        }
    }
    return true
}

// Checks if the string contains only lowercase letters
func IsLower(s string) bool {
    if len(s) == 0 {
        return false
    }
    for _, r := range s {
        if r < 'a' || r > 'z' {
            return false
        }
    }
    return true
}

// Checks if the string contains only numeric characters
func IsNumeric(s string) bool {
    if len(s) == 0 {
        return false
    }
    for _, r := range s {
        if r < '0' || r > '9' {
            return false
        }
    }
    return true
}

// Checks if the string contains only alphanumeric characters
func IsAlphanumeric(s string) bool {
    if len(s) == 0 {
        return false
    }
    for _, r := range s {
        if (r < 'A' || r > 'Z') && (r < 'a' || r > 'z') && (r < '0' || r > '9') {
            return false
        }
    }
    return true
}
