package leet

import "testing"

func TestIntHeap(t *testing.T) {
	ih := make(intheap, 0)

	ih.push(posval{1,1})
	ih.push(posval{6,1})
	ih.push(posval{3,1})
	ih.push(posval{4,1})
	ih.push(posval{2,1})
	ih.push(posval{1,1})
	ih.push(posval{6,1})
	ih.push(posval{3,1})

	t.Log(ih)

	for len(ih) > 0 {
		t.Log(ih.pop())
	}
}

func convertSlice2ListNode(xs []int) *ListNode {
	if len(xs) == 0 {
		return nil
	}

	cl := &ListNode{
		Val: xs[0],
	}

	res := cl

	for i := 1; i < len(xs); i++ {
		nn := &ListNode{
			Val: xs[i],
		}

		cl.Next = nn
		cl = nn
	}
	return res
}

func convertSlices2ListNodes(xys [][]int) []*ListNode {
	res := make([]*ListNode, len(xys))

	for i, xs := range xys {
		res[i] = convertSlice2ListNode(xs)
	}
	return res
}

func isEqual(xs []int, ll *ListNode) bool {
	i, cl := 0, ll
	for i < len(xs) && cl != nil {
		if xs[i] != cl.Val {
			return false
		}
		i, cl = i+1, cl.Next
	}
	return i == len(xs) && cl == nil
}

func TestMergeKLists(t *testing.T) {
	input := convertSlices2ListNodes([][]int{{-8,-7,-7,-5,1,1,3,4},{-2},{-10,-10,-7,0,1,3},{2}})

	expected := []int{-10,-10,-8,-7,-7,-7,-5,-2,0,1,1,1,2,3,3,4}

	res := MergeKLists(input)

	if !isEqual(expected, res) {
		t.Errorf("expected %v, got %s", expected, res.String())
	}
}
