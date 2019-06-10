package leet

import (
	"reflect"
	"testing"
)

func TestInsertInterval(t *testing.T) {
	intervals := [][]int{{1,3}, {6,9}}
	newInterval := []int{2,5}

	expected := [][]int{{1,5},{6,9}}

	res := InsertInterval(intervals, newInterval)

	if ! reflect.DeepEqual(expected, res) {
		t.Errorf("expected %+v, got %+v", expected, res)
	}

	intervals = [][]int{{1,2},{3,5},{6,7},{8,10},{12,16}}
	newInterval = []int{4,8}

	expected = [][]int{{1,2},{3,10},{12,16}}

	res = InsertInterval(intervals, newInterval)

	if ! reflect.DeepEqual(expected, res) {
		t.Errorf("expected %+v, got %+v", expected, res)
	}

	intervals = [][]int{{1,5}}
	newInterval = []int{2,3}

	expected = [][]int{{1,5}}

	res = InsertInterval(intervals, newInterval)

	if ! reflect.DeepEqual(expected, res) {
		t.Errorf("expected %+v, got %+v", expected, res)
	}

	intervals = [][]int{{1,5}}
	newInterval = []int{6,8}

	expected = [][]int{{1,5}, {6, 8}}

	res = InsertInterval(intervals, newInterval)

	if ! reflect.DeepEqual(expected, res) {
		t.Errorf("expected %+v, got %+v", expected, res)
	}
}
