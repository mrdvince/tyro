func numTrees(n int) int {
    var ans float64 = 1
    for idx := 0; idx < n; idx++ {
       ans *= float64((4 * idx + 2)) / float64((idx + 2))
    }
    return int(ans)
}
