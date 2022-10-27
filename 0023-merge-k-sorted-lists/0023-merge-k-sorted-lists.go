/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
   if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}
    
    prev := lists[0]
	var res *ListNode
	for i := 1; i < len(lists); i++ {
		res = mergeTwoLists(prev, lists[i])
        prev= res      
	}

	return res
    
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := ListNode{}
    dummy.Next = nil
    tail := &dummy
    
    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            tail.Next = list1
            list1 = list1.Next
        } else {
            tail.Next = list2
            list2 = list2.Next
        }
        tail = tail.Next
    }
    
    if list1 != nil {
        tail.Next = list1
    }
    if list2 != nil {
        tail.Next = list2
    }
    return dummy.Next
}