package sprint

func StrToInt(s string) int {
    result := 0
    sign := 1
    start := 0

    if len(s) > 0 && s[0] == '-' {
        sign = -1
        start = 1
    }

    for i := start; i < len(s); i++ {
        if s[i] < '0' || s[i] > '9' {
            return 0
        }
        result = result*10 + int(s[i]-'0')
    }

    return sign * result
}

func BulkAtoi(arr []string) []int {
    result := make([]int, len(arr))
    for i, str := range arr {
        result[i] = StrToInt(str)
    }
    return result
}
