/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    var nums []int
    for head != nil {
        nums = append(nums, head.Val)
        head = head.Next
    }
    left_ptr, right_ptr := 0, len(nums) - 1
    for left_ptr <= right_ptr {
        if nums[left_ptr] != nums[right_ptr] {
            return false
        }
        left_ptr += 1
        right_ptr -= 1
    }
    return true
}