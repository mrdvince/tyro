func groupAnagrams(strs []string) [][]string {
    res := [][]string{}
    
    sortvals := func(x string) string {
        s := []rune(x)
        sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
        return string(s)
    }
    
    newstrs := map[string][]string{}
    for idx, word := range strs {
        newstrs[sortvals(word)] = append(newstrs[sortvals(word)], strs[idx])
    }
    
    for _, words := range newstrs {
        res = append(res, words)
    }
    return res
}