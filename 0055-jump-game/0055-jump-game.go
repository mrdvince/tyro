func canJump(nums []int) bool {
    end := len(nums) - 1
    for idx := len(nums) -1; idx >= 0; idx-- {
        if idx + nums[idx] >= end {
            end = idx
        }
    }
    if end == 0 {
        return true
    } else {
        return false
    }
}