func search(nums []int, target int) int {
	lo, hi := 0, len(nums) -1

	for lo <= hi {
		mid := (lo + hi) // 2
		curr_num := nums[mid]
		if curr_num == target {
			return mid
		}
		if curr_num < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}

	}
	return -1
}