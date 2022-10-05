/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    nodeList := []*TreeNode{root}
    values := [][]int{}
    idx := 0
    
    for len(nodeList) > 0 {
        values = append(values, []int{})
        for _, node := range nodeList {
            values[idx] = append(values[idx], node.Val)
            nodeList = nodeList[1:]
            if node.Left != nil {
                nodeList = append(nodeList, node.Left)
            }
            if node.Right != nil {
                nodeList = append(nodeList, node.Right)
            }
        } 
        idx ++
    }
    return values
}