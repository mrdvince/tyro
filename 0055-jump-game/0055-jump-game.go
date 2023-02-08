func canJump(nums []int) bool {
	nos := len(nums)
	if nos == 0 {
		return false
	}

	lastPos := nos - 1
	for idx := nos - 1; idx >= 0; idx-- {
		if idx+nums[idx] >= lastPos {
			lastPos = idx
		}
	}
	return lastPos == 0
}
