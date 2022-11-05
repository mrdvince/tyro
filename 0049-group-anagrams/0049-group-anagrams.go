func groupAnagrams(strs []string) [][]string {
	res := map[string][]string{}

	for _, s := range strs {
		count := make([]int, 26)
		for _, c := range s {
			count[c-'a'] += 1
		}
		key := strings.Trim(fmt.Sprint(count), "[]")
		res[key] = append(res[key], s)
	}

	values := make([][]string, 0, len(res))
	for _, v := range res {
		values = append(values, v)
	}
	fmt.Println(values)

	return values
}