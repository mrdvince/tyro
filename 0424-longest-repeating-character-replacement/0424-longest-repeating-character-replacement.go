func characterReplacement(s string, k int) int {
    res, leftPtr, maxF := 0, 0, 0
    count := map[byte]int{}

    for rightPtr, _ := range s {
        count[s[rightPtr]] = 1 + count[s[rightPtr]]
        maxF = max(maxF, count[s[rightPtr]])
        
        windowLen := (rightPtr - leftPtr + 1)
        if  windowLen - maxF > k {
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