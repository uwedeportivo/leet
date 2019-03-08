package leet

import (
	"bytes"
	"io/ioutil"
	"strconv"
	"testing"
	"text/scanner"
)

func TestCalculateMinimumHP(t *testing.T) {
	nums := [][]int{{-2,-3,3},{-5,-10,1},{10,30,-5}}
	expected := 7

	res := CalculateMinimumHP(nums)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}

	nums = [][]int{{1,-3,3},{0,-2,0},{-3,-3,-3}}
	expected = 3

	res = CalculateMinimumHP(nums)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}

	dungeonData, err := ioutil.ReadFile("testdata/dungeon.txt")
	if err != nil {
		t.Errorf("error reading testdata/dungeon.txt: %v", err)
	}

	var s scanner.Scanner
	s.Init(bytes.NewReader(dungeonData))
	s.Filename = "testdata/dungeon.txt"
	s.Mode |= scanner.ScanInts

	nums = nil
	var row []int
	negate := false
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		txt := s.TokenText()

		if txt == "[" || txt == "," {
		} else if txt == "]" {
			if row != nil {
				nums = append(nums, row)
				row = nil
			}
		} else if txt == "-" {
			negate = true
		} else {
			val, err := strconv.ParseInt(txt, 10, 0)
			if err != nil {
				t.Errorf("error reading testdata/dungeon.txt: %v", err)
			}
			if negate {
				val = -val
				negate = false
			}
			row = append(row, int(val))
		}
	}

	expected = 85

	res = CalculateMinimumHP(nums)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
