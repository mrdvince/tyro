/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    fast := head
    for fast != nil && fast.Next !=nil {
        head = head.Next
        fast = fast.Next.Next
    }
    return head
}