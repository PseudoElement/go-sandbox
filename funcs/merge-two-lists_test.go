package funcs

import (
	"slices"
	"testing"
)

type TestCase struct {
	list1  *ListNode
	list2  *ListNode
	expect []int
}

var cases_TestMergeTwoLists = []TestCase{
	{
		list1: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
		list2: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
		expect: []int{1, 1, 2, 3, 4, 4},
	},
	{
		list1:  &ListNode{},
		list2:  &ListNode{},
		expect: []int{0, 0},
	},
	{
		list1: &ListNode{
			Val: (-9),
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
		list2: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val:  7,
				Next: nil,
			},
		},
		expect: []int{-9, 3, 5, 7},
	},
}

func TestMergeTwoLists(t *testing.T) {
	for idx, testCase := range cases_TestMergeTwoLists {
		res := MergeTwoListsToSlice(testCase.list1, testCase.list2)
		if slices.Compare(res, testCase.expect) != 0 {
			t.Errorf(`%d TestMergeTwoLists error: got %+v, expected %+v.`, idx, res, testCase.expect)
		}
	}
}
