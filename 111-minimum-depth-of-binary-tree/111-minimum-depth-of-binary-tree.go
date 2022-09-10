/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth, queue := 1, []*TreeNode{root}

	for len(queue) > 0 {
        tmp := []*TreeNode{}
		for _, curr_node := range queue {
			// are we at a leaf node?
			if curr_node.Left == nil && curr_node.Right == nil {
				return depth
			}
			if curr_node.Left != nil {
				tmp = append(tmp, curr_node.Left)
			}
			if curr_node.Right != nil {
				tmp = append(tmp, curr_node.Right)
			}
		}
		// if we got here means we are still going
		depth++
		queue = tmp
	}
	return depth
}