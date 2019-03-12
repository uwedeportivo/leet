package leet

import (
	"bytes"
	"io/ioutil"
	"strconv"
	"testing"
	"text/scanner"
)

func TestMonotoneEnd(t *testing.T) {
	nums := []int{1,2,3,4,5,3,2,1,2,1,2,3,4,3}

	for s := 0; s < len(nums); s++ {
		t.Logf("monotone sequence starting at %d ends at %d", s, monotoneEnd(nums, s))
	}
}

func TestDelAdjacentDuplcates(t *testing.T) {
	nums := []int{1,1,2,2,3,4,5,5}

	t.Logf("adjancent duplicates of %v removed %v", nums, delAdjacentDuplcates(nums))
}

func TestCollapseHills(t *testing.T) {
	nums := []int{3,2,6,5,0,3}

	t.Logf("collapsed hills of %v are %v", nums, collapseHills(nums))
}


func readBuySellTestData(filename string) ([]int, error) {
	buysellData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var s scanner.Scanner
	s.Init(bytes.NewReader(buysellData))
	s.Filename = filename
	s.Mode |= scanner.ScanInts

	var nums []int
	negate := false
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		txt := s.TokenText()

		if txt == "[" || txt == "," || txt == "]" {
		} else if txt == "-" {
			negate = true
		} else {
			val, err := strconv.ParseInt(txt, 10, 0)
			if err != nil {
				return nil, err
			}
			if negate {
				val = -val
				negate = false
			}
			nums = append(nums, int(val))
		}
	}
	return nums, nil
}

func TestMaxProfit(t *testing.T) {
	nums := []int{2,4,1}
	k := 2
	expected := 2

	res := MaxProfit(k, nums)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}

	nums = []int{3,2,6,5,0,3}
	k = 2
	expected = 7

	res = MaxProfit(k, nums)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}

	nums = []int{1,2,4,2,5,7,2,4,9,0}
	k = 2
	expected = 13

	res = MaxProfit(k, nums)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}

	nums = []int{1,4,7,5,6,2,5,1,9,7,9,7,0,6,8}
	k = 5
	expected = 27

	res = MaxProfit(k, nums)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}

	nums, err := readBuySellTestData("testdata/buysell.txt")
	if err != nil {
		t.Errorf("error reading testdata/buysell.txt: %v", err)
	}

	k = 29
	expected = 2818

	res = MaxProfit(k, nums)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}

	nums, err = readBuySellTestData("testdata/buysell2.txt")
	if err != nil {
		t.Errorf("error reading testdata/buysell2.txt: %v", err)
	}

	k = 1000000000
	expected = 1648961

	res = MaxProfit(k, nums)

	if expected != res {
		t.Errorf("expected %d, got %d", expected, res)
	}
}
