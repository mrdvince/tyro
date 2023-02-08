func canJump(nums []int) bool {
	length := len(nums)
	if length == 0 {
		return false
	}

	lastReachableIndex := length - 1
	for i := length - 1; i >= 0; i-- {
		if i+nums[i] >= lastReachableIndex {
			lastReachableIndex = i
		}
	}

	return lastReachableIndex == 0
}