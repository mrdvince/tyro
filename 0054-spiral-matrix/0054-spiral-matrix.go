func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 {
        return []int{}
    }
    // rows and cols represent the number of rows and columns in the matrix
    rows, cols := len(matrix), len(matrix[0])
    result := make([]int, rows * cols)
    // row and col represent the current row and column position
    row, col := 0, 0
    // d represents the current direction of movement (0 = right, 1 = down, 2 = left, 3 = up)
    d := 0
    // dx and dy are slices that specify the change in the column (horizontal) and row (vertical) position
    dy := []int{0, 1, 0, -1}
    dx := []int{1, 0, -1, 0}

    for i := 0; i < rows * cols; i++ {
        result[i] = matrix[row][col]
        matrix[row][col] = 0
        nrow, ncol := row + dy[d], col + dx[d]
        // if next position is out of bounds or has already been visited, 
        // change the direction of movement and recalculate the next position
        if nrow < 0 || nrow >= rows || ncol < 0 || ncol >= cols || matrix[nrow][ncol] == 0 {
            d = (d + 1) % 4
            nrow, ncol = row + dy[d], col + dx[d]
        }
        row, col = nrow, ncol
    }
    return result
}
