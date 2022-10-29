func combinationSum(candidates []int, target int) [][]int {
    res := [][]int{}
    currComb := []int{}
    backtrack(&res, candidates, &currComb, 0, len(candidates), target)
    return res
}

func backtrack(res *[][]int, candidates []int, curr *[]int, start, end, target int) {
    if target == 0 {
        currCopy := make([]int, len(*curr))
        copy(currCopy, *curr)
        *res = append(*res, currCopy)
    } else if target > 1{
        for i := start; i < end; i++ {
            *curr = append(*curr, candidates[i])
            backtrack(res, candidates, curr, i, end, target - candidates[i])
            *curr = (*curr)[:len(*curr) - 1]
        }
    }
}