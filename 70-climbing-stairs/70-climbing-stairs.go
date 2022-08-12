var memo = make(map[int]int)
func climbStairs(n int) int {
    if n == 1 {
        return 1
    }
    if n == 2 {
        return 2
    }
    if val, ok := memo[n]; ok {
        return val
    }
    memo[n] = climbStairs(n-1) + climbStairs(n-2)
    return memo[n]
}