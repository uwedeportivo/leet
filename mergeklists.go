package leet

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func (ll *ListNode) String() string {
	sb := strings.Builder{}

	sb.WriteString("[")
	first := true

	for cl := ll; cl != nil; cl = cl.Next {
		if first {
			first = false
		} else {
			sb.WriteString(", ")
		}

		sb.WriteString(fmt.Sprintf("%d", cl.Val))
	}
	sb.WriteString("]")
	return sb.String()
}

type posval struct {
	val int
	pos int
}

type intheap []posval

func (ih *intheap) push(pv posval) {
	*ih = append(*ih, pv)

	xs := *ih
	i := len(xs)

	for i > 1 {
		pi := i / 2

		if xs[pi-1].val > xs[i-1].val {
			xs[pi-1], xs[i-1] = xs[i-1], xs[pi-1]
		}
		i = pi
	}
}

func (ih *intheap) pop() posval {
	xs := *ih
	res := xs[0]
	n := len(xs)
	xs[0] = xs[n-1]
	i := 1
	*ih=(*ih)[:n-1]
	xs = *ih

	for {
		li := 2 * i
		ri := 2 * i + 1

		if li <= len(xs) && ri <= len(xs) {
			if xs[i-1].val <= xs[li-1].val && xs[i-1].val <= xs[ri-1].val {
				break
			}

			if xs[li-1].val < xs[ri-1].val {
				xs[i-1], xs[li-1] = xs[li-1], xs[i-1]
				i = li
			} else {
				xs[i-1], xs[ri-1] = xs[ri-1], xs[i-1]
				i = ri
			}
		} else if li <= len(xs) && xs[i-1].val > xs[li-1].val {
			xs[i-1], xs[li-1] = xs[li-1], xs[i-1]
			i = li
		} else {
			break
		}
	}

	return res
}

func MergeKLists(lists []*ListNode) *ListNode {
	ih := make(intheap, 0)
	var res *ListNode
	var cur *ListNode

	for i, ll := range lists {
		if ll != nil {
			ih.push(posval{ll.Val, i})
			lists[i] = ll.Next
		}
	}

	for len(ih) > 0 {
		pv := ih.pop()

		nn := &ListNode{
			Val:pv.val,
		}

		if cur == nil {
			res = nn
		} else {
			cur.Next = nn
		}
		cur = nn

		if lists[pv.pos] != nil {
			ll := lists[pv.pos]
			ih.push(posval{ll.Val, pv.pos})
			lists[pv.pos] = ll.Next
		}
	}

	return res
}
