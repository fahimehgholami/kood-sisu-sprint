package sprint

func ArrAny(f func(string) bool, a []string) bool {
    for _, element := range a {
        if f(element) {
            return true
        }
    }
    return false
}

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
