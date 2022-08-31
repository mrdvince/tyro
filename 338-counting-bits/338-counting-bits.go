func countBits(n int) []int {
    rep := make([]int, n+1)
    for idx := 1; idx < n + 1; idx ++ {
        rep[idx] = rep[idx / 2] + idx % 2
    }
    return rep
}