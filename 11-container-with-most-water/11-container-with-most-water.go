func maxArea(height []int) int {
    maxArea, leftPtr, rightPtr := 0, 0, len(height) - 1
    
    for leftPtr < rightPtr {
        gap := rightPtr - leftPtr
        maxArea = max(maxArea, min(height[leftPtr], height[rightPtr]) * gap)
        if height[leftPtr] < height[rightPtr] {
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