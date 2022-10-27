func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	for idx, num := range nums {
		if idx > 0 && num == nums[idx-1] {
			continue
		}

		leftPtr, rightPtr := idx+1, len(nums)-1
		for leftPtr < rightPtr {
			threeSum := num + nums[leftPtr] + nums[rightPtr]

			if threeSum > 0 {
				rightPtr -= 1
			} else if threeSum < 0 {
				leftPtr += 1
			} else {
				res = append(res, []int{num, nums[leftPtr], nums[rightPtr]})
				leftPtr += 1
				for nums[leftPtr] == nums[leftPtr-1] && leftPtr < rightPtr {
					leftPtr += 1
				}
			}
		}
	}

	return res
}
