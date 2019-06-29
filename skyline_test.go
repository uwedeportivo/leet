package leet

import (
	"reflect"
	"testing"
)

func TestGetSkyline(t *testing.T) {
	buildings := [][]int{{1,2,1},{1,2,2},{1,2,3}}
	expected := [][]int{{1,3},{2,0}}

	res := GetSkyline(buildings)

	if ! reflect.DeepEqual(expected, res) {
		t.Errorf("expected %+v, got %+v", expected, res)
	}
}


