/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    newNode := ListNode{Next:head}
    dummy := &newNode
    prev, curr := dummy, head
    for curr != nil {
        next := curr.Next
        if curr.Val == val {
            prev.Next = next
        } else {
            prev = curr
        }
        curr = next
    }
    return dummy.Next
}