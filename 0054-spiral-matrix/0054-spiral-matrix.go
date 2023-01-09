func spiralOrder(matrix [][]int) []int {
    res := []int{}
    left, right, top, bottom := 0, len(matrix[0]), 0, len(matrix)
    
    update := func(num int) {
        res = append(res, num)
    }
    
    for left < right && top < bottom {
        // left to right
        for i := left; i < right; i++ {
            update(matrix[top][i])
        }
        // done with top, shift down
        top += 1
        // get values in right
        for i := top; i < bottom; i++ {
            update(matrix[i][right - 1])
        }
        // done with right, shift left
        right -= 1
        if left >= right || top >= bottom {
            break
        }
        // got bottom in reverse
        for i := right-1; i >= left; i-- {
            update(matrix[bottom - 1][i])
        }
        // done with bottom, shift up
        bottom -= 1
        // go up 
        for i := bottom - 1; i >= top; i-- {
            update(matrix[i][left])
        }
        left += 1
    }
    
    return res
}