package leet

import (
	"reflect"
	"testing"
)

func TestFullJustify(t *testing.T) {
	words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	maxWidth := 16
	expected := []string{
		"This    is    an",
		"example  of text",
		"justification.  ",
	}

	res := FullJustify(words, maxWidth)

	if ! reflect.DeepEqual(expected, res) {
		t.Errorf("expected %+v, got %+v", expected, res)
	}

	words = []string{"What","must","be","acknowledgment","shall","be"}
	maxWidth = 16
	expected = []string{
		"What   must   be",
		"acknowledgment  ",
		"shall be        ",
	}

	res = FullJustify(words, maxWidth)

	if ! reflect.DeepEqual(expected, res) {
		t.Errorf("expected %+v, got %+v", expected, res)
	}

	words = []string{"Science","is","what","we","understand","well","enough","to","explain",
		"to","a","computer.","Art","is","everything","else","we","do"}
	maxWidth = 20
	expected = []string{
		"Science  is  what we",
		"understand      well",
		"enough to explain to",
		"a  computer.  Art is",
		"everything  else  we",
		"do                  ",
	}

	res = FullJustify(words, maxWidth)

	if ! reflect.DeepEqual(expected, res) {
		t.Errorf("expected %+v, got %+v", expected, res)
	}
}
