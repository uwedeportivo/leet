package leet

import (
	"fmt"
	"strings"
	"testing"
)

func toLinkedList(xs []int) *ListNode {
	if len(xs) == 0 {
		return nil
	}

	head := &ListNode{
		Val:xs[0],
	}

	prev := head
	for i := 1; i < len(xs); i++ {
		n := &ListNode{
			Val:xs[i],
		}
		prev.Next = n
		prev = n
	}

	return head
}

func toString(head *ListNode) string {
	var sb strings.Builder

	for cursor := head; cursor != nil; cursor = cursor.Next {
		sb.WriteString(fmt.Sprintf("%d", cursor.Val))
		if cursor.Next != nil {
			sb.WriteString(" -> ")
		}
	}
	return sb.String()
}

func TestReverseKGroup(t *testing.T) {
	xs := []int{1,2,3,4,5}


	t.Log(ReverseKGroup(toLinkedList(xs), 2))
	t.Log(ReverseKGroup(toLinkedList(xs), 3))
}
