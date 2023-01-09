func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 {
        return []int{}
    }

    rows, cols := len(matrix), len(matrix[0])
    result := make([]int, 0, rows * cols)

    row, col := 0, 0
    d := 0
    dx := []int{1, 0, -1, 0}
    dy := []int{0, 1, 0, -1}

    for i := 0; i < rows * cols; i++ {
        result = append(result, matrix[row][col])
        matrix[row][col] = 0
        nrow, ncol := row + dy[d], col + dx[d]
        if nrow < 0 || nrow >= rows || ncol < 0 || ncol >= cols || matrix[nrow][ncol] == 0 {
            d = (d + 1) % 4
            nrow, ncol = row + dy[d], col + dx[d]
        }
        row, col = nrow, ncol
    }

    return result
}
