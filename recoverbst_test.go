package leet

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"text/scanner"
)

type bstParser struct {
	scn scanner.Scanner
	firstErr error
	peeked bool
}

func (bp *bstParser) peek() string {
	if bp.firstErr != nil {
		return ""
	}
	tok := bp.nextToken()
	bp.peeked = true
	return tok
}

func (bp *bstParser) nextToken() string {
	if bp.firstErr != nil {
		return ""
	}
	if bp.peeked {
		bp.peeked = false
	} else {
		tok := bp.scn.Scan()

		if tok == scanner.EOF {
			bp.firstErr = fmt.Errorf("parse error: unexpected EOF at %v", bp.scn.Pos())
			return ""
		}
	}
	txt := bp.scn.TokenText()

	if txt == "-" {
		valStr := bp.nextToken()
		return "-" + valStr
	}
	return txt
}

func (bp *bstParser) consumeInt() int64 {
	if bp.firstErr != nil {
		return 0
	}
	tokenStr:= bp.nextToken()
	val, err := strconv.ParseInt(tokenStr, 10, 0)
	if err != nil {
		bp.firstErr = err
		return 0
	}

	return val
}

func (bp *bstParser) consume(expected string) {
	if bp.firstErr != nil {
		return
	}
	tokenStr := bp.nextToken()

	if tokenStr != expected {
		bp.firstErr = fmt.Errorf("parse error: expected %s at %v", expected, bp.scn.Pos())
	}
}

func (bp *bstParser) parseNode() *TreeNode {
	if bp.firstErr != nil {
		return nil
	}

	if bp.peek() != "(" {
		val := bp.consumeInt()
		return &TreeNode{
			Val:val,
		}
	}

	bp.consume("(")
	defer bp.consume(")")

	if bp.peek() == ")" {
		return nil
	}

	val := bp.consumeInt()
	leftChild := bp.parseNode()
	rightChild := bp.parseNode()

	return &TreeNode{
		Val:val,
		Left:leftChild,
		Right:rightChild,
	}
}

// (5 (4 () 3) ())
func parseBST(str string) (*TreeNode, error) {
	var s scanner.Scanner
	s.Init(strings.NewReader(str))
	s.Mode |= scanner.ScanInts

	bp := bstParser{
		scn:s,
	}

	root := bp.parseNode()

	return root, bp.firstErr
}

func (node *TreeNode) toString(strbldr *strings.Builder) {
	if node != nil && node.Left == nil && node.Right == nil {
		strbldr.WriteString(fmt.Sprintf("%d", node.Val))
		return
	}
	strbldr.WriteString("(")
	if node != nil {
		strbldr.WriteString(fmt.Sprintf("%d ", node.Val))
		node.Left.toString(strbldr)
		strbldr.WriteString(" ")
		node.Right.toString(strbldr)
	}
	strbldr.WriteString(")")
}

func (node *TreeNode) String() string {
	var strBldr strings.Builder

	node.toString(&strBldr)
	return strBldr.String()
}

func TestRecoverTree(t *testing.T) {
	root, err := parseBST("(3 1 (4 2 ()))")
	if err != nil {
		t.Error(err)
	}

	RecoverTree(root)
	t.Log(root)

	root, err = parseBST("(1 (3 () 2) ())")
	if err != nil {
		t.Error(err)
	}

	RecoverTree(root)
	t.Log(root)

	root, err = parseBST("(3 () (2 () 1))")
	if err != nil {
		t.Error(err)
	}

	RecoverTree(root)
	t.Log(root)
}
