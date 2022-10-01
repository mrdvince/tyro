func romanToInt(s string) int {
    numerals := map[byte]int{
        'I':1,
        'V':5,
        'X':10,
        'L':50,
        'C':100,
        'D':500,
        'M':1000,
    }
    result := 0
    for idx := 0; idx < len(s); idx ++ {
        if idx + 1 < len(s) && numerals[s[idx]] < numerals[s[idx+1]] {
            result -= numerals[s[idx]]
        } else {
            result += numerals[s[idx]]
        }
    } 
    return result
    
}