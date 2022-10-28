func search(nums []int, target int) int {
    leftPtr, rightPtr := 0, len(nums) - 1
    
    for leftPtr <= rightPtr {
        mid := (leftPtr + rightPtr) // 2
        if nums[mid] == target {
            return mid
        }
        mid_num := nums[mid]
        if nums[leftPtr] <= mid_num {
            if target < nums[leftPtr] || target > mid_num {
                // got right
                leftPtr = mid + 1
            } else {
                rightPtr = mid - 1
            }
        } else {
            if target > nums[rightPtr] || target < mid_num {
                // go left
                rightPtr = mid - 1
            } else {
                leftPtr = mid + 1
            }
        }
    }
    
    return -1
}