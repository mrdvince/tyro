func maxArea(height []int) int {
    maxArea, leftPtr, rightPtr := 0, 0, len(height) - 1
    
    for leftPtr < rightPtr {
        gap, left_val, right_val := (rightPtr - leftPtr), height[leftPtr], height[rightPtr]
        maxArea = max(maxArea, min(left_val, right_val) * gap)
        
        if left_val < right_val {
            leftPtr ++
        } else {
            rightPtr --
        }
    }
    return maxArea
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}