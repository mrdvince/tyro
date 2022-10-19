func characterReplacement(s string, k int) int {
    res := 0
    count := map[byte]int{}
    leftPtr := 0
    maxF := 0
    
    for rightPtr, _ := range s {
        count[s[rightPtr]] = 1 + count[s[rightPtr]]
        maxF = max(maxF, count[s[rightPtr]])
        if (rightPtr - leftPtr + 1) - maxF > k {
            count[s[leftPtr]] -= 1
            leftPtr +=1
        }
        res = max(maxF, rightPtr - leftPtr + 1)
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}