func findDisappearedNumbers(nums []int) []int {
	numsList := make([]int, len(nums))
	for _, num := range nums {
		numsList[num-1] = num
	}
	result := make([]int, 0)
	for i := 0; i < len(numsList); i++ {
		if numsList[i] == 0 {
			result = append(result, i+1)
		}
	}

	return result
}