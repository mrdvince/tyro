func spiralOrder(matrix [][]int) []int {
    res := []int{}
    left, right, top, bottom := 0, len(matrix[0]), 0, len(matrix)

    for left < right && top < bottom {
        // left to right
        for i := left; i < right; i++ {
            res = append(res, matrix[top][i])
        }
        // done with top, shift down
        top += 1
        // get values in right
        for i := top; i < bottom; i++ {
            res = append(res, matrix[i][right - 1])
        }
        // done with right, shift left
        right -= 1
        if left >= right || top >= bottom {
            break
        }
        // got bottom in reverse
        for i := right-1; i >= left; i-- {
            res = append(res, matrix[bottom - 1][i])
        }
        // done with bottom, shift up
        bottom -= 1
        // go up 
        for i := bottom - 1; i >= top; i-- {
            res = append(res, matrix[i][left])
        }
        left += 1
    }
    
    return res
}