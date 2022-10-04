func maxArea(height []int) int {
    maxArea, leftPtr, rightPtr := 0, 0, len(height) - 1
    if len(height) <=2 {
        return min(height[0], height[1])
    }
    for leftPtr < rightPtr {        
        maxArea = max(maxArea, min(height[leftPtr], height[rightPtr]) * (rightPtr - leftPtr))
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