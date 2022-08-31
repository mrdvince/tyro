func countBits(n int) []int {
    rep := make([]int, 0)
    for idx := 0; idx < n + 1; idx ++ {
        rep = append(rep, strings.Count(strconv.FormatInt(int64(idx),2), "1"))
    }
    return rep
}