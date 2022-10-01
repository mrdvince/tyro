func romanToInt(s string) int {
    numerals := map[string]int{
        "I":1,
        "V":5,
        "X":10,
        "L":50,
        "C":100,
        "D":500,
        "M":1000,
    }
    result := 0
    for idx := 0; idx < len(s); idx ++ {
        if idx + 1 < len(s) && numerals[string(s[idx])] < numerals[string(s[idx+1])] {
            result -= numerals[string(s[idx])]
        } else {
            result += numerals[string(s[idx])]
        }
    } 
    return result
    
}