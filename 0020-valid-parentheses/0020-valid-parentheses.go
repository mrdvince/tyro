func isValid(s string) bool {
    
    brackets := map[rune]rune{
        ')':'(', 
        '}':'{', 
        ']':'[',
    }
    stack := make([]rune,0)
    for _, bracket := range s {
        switch bracket {
            case '(', '{', '[':
            stack = append(stack, bracket)
            
            case ')', '}', ']':
            if len(stack) == 0 || stack[len(stack)-1] != brackets[bracket] {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }
    
    return len(stack) == 0
    
}