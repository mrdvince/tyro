func groupAnagrams(strs []string) [][]string {
	res := map[[26]int][]string{}
	for _, s := range strs {
        count := [26]int{}
		for _, c := range s {
			count[c-'a'] += 1
		}
		res[count] = append(res[count], s)
	}

	values := make([][]string, 0, len(res))
	for _, v := range res {
		values = append(values, v)
	}

	return values
}