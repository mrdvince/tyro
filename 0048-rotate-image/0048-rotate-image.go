func rotate(matrix [][]int) {
	left, right := 0, len(matrix)-1

	for left < right {
		for idx := 0; idx < (right - left); idx++ {
			top, bottom := left, right
			topLeft := matrix[top][left+idx]
			matrix[top][left+idx] = matrix[bottom-idx][left]
			matrix[bottom-idx][left] = matrix[bottom][right-idx]
			matrix[bottom][right-idx] = matrix[top+idx][right]
			matrix[top+idx][right] = topLeft
		}
		left += 1
		right -= 1
	}
}