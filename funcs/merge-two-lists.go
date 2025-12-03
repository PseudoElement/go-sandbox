package funcs

func MergeTwoListsToSlice(list1 *ListNode, list2 *ListNode) []int {
	return toSlice(mergeTwoLists(list1, list2))
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	tail := result

	next1 := list1
	next2 := list2

	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	for next1 != nil && next2 != nil {
		if next1.Val < next2.Val {
			tail.Val = next1.Val
			next1 = next1.Next
		} else {
			tail.Val = next2.Val
			next2 = next2.Next
		}

		tail.Next = &ListNode{}
		tail = tail.Next
	}

	if next1 != nil {
		for next1 != nil {
			tail.Val = next1.Val
			next1 = next1.Next
			if next1 != nil {
				tail.Next = &ListNode{}
				tail = tail.Next
			}
		}
		return result
	}
	if next2 != nil {
		for next2 != nil {
			tail.Val = next2.Val
			next2 = next2.Next
			if next2 != nil {
				tail.Next = &ListNode{}
				tail = tail.Next
			}
		}
		return result
	}

	return result
}

func toSlice(list *ListNode) []int {
	s := make([]int, 0)

	next := list
	for next != nil {
		s = append(s, next.Val)
		next = next.Next
	}

	return s
}
