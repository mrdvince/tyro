/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
    result, queue := make([]float64, 0), []*TreeNode{root}
	for len(queue) > 0 {
		size, sum := len(queue), 0
		for i := 0; i < size; i++ {
			temp := queue[0]
			sum += temp.Val
			if temp.Left != nil {
				queue = append(queue, temp.Left)
			}
			if temp.Right != nil {
				queue = append(queue, temp.Right)
			}
			queue = queue[1:]
		}
		result = append(result, float64(sum)/float64(size))
	}
	return result
}