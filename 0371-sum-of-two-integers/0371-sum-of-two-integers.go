func getSum(a int, b int) int {
    
    for b != 0 {
        carry := a & b
        fmt.Println(carry)
        a = a ^ b
        b = carry << 1
    }
    return a
}