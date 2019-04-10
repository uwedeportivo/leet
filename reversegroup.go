package leet

type nodestack []*ListNode

func (s *nodestack) push(v *ListNode) {
	*s = append(*s, v)
}

func (s *nodestack) pop() *ListNode {
	n := len(*s)
	res:=(*s)[n-1]
	*s=(*s)[:n-1]
	return res
}

func (s *nodestack) peek() *ListNode {
	n := len(*s)
	return (*s)[n-1]
}

func findKGroup(head *ListNode, k int) *ListNode {
	cursor := head
	i := 1
	for i < k && cursor != nil {
		cursor = cursor.Next
		i++
	}

	if i == k {
		return cursor
	}
	return nil
}

func reverseGroup(head, tail *ListNode) {
	if head == nil || tail == nil {
		return
	}
	tailsNext := tail.Next
	nss := make(nodestack, 0, 64)

	for cursor := head; cursor != tail && cursor != nil; cursor = cursor.Next {
		nss.push(cursor)
	}

	if len(nss) == 0 {
		return
	}

	prev, cursor := tail, nss.pop()
	prev.Next = cursor

	for len(nss) > 0 {
		prev, cursor = cursor, nss.pop()
		prev.Next = cursor
	}

	head.Next = tailsNext
}

func ReverseKGroup(head *ListNode, k int) *ListNode {
	cursor := head
	tail := findKGroup(cursor, k)
	res := tail
	if tail == nil {
		res = head
	}

	var prevGroupTail *ListNode

	for tail != nil {
		reverseGroup(cursor, tail)
		if prevGroupTail != nil {
			prevGroupTail.Next = tail
		}
		prevGroupTail = cursor
		cursor = cursor.Next
		tail = findKGroup(cursor, k)
	}

	return res
}
