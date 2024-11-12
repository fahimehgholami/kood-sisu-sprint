package sprint


func StrLength(s string) []int {
    runeCount := 0
    byteCount := len(s) 

    for range s {
        runeCount++ 
    }

    return []int{runeCount, byteCount} 
}