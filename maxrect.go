package leet

type stack []int

func (s *stack) push(v int) {
	*s = append(*s, v)
}

func (s *stack) pop() int {
	n := len(*s)
	res:=(*s)[n-1]
	*s=(*s)[:n-1]
	return res
}

func (s *stack) peek() int {
	n := len(*s)
	return (*s)[n-1]
}

// for each hg[j] assume we know indices i and k such that
// i < j < k && hg[i] < hg[j] && hg[k] < hg[j]
// && (forall u: i < u < j:  hg[u] >= hg[j])
// && (forall u: j < u < k:  hg[u] >= hg[j])

// then the max rect we can achieve with height hg[j] has
// width [i+1, k), so k - i - 1 and the area is
// hg[j] * (k - i - 1)

// how can we get (i, k) for a given j in O(1)
//
// the idea is to keep a stack of indices for which hg values
// are strictly ascending (*)
// before pushing a new index, all indices that would violate
// (*) are popped
// for an index j that is popped the incoming index popping it is its k
// and the index under it in the stack is its i
//
// each index with distinct hg will be put on the stack and popped so each hg value will be
// considered
// for each index j getting its i and k is O(1) courtesy of the stack
// overall runtime is O(n) in length of histogram
func maxRectUnderHistogram(hg []int) int {
	res := 0

	s := make(stack, 0, len(hg))

	for idx, h := range hg {
		if len(s) == 0 {
			s.push(idx)
			continue
		}

		for len(s) > 0 && h < hg[s.peek()] {
			j := s.pop()
			var i int
			if len(s) > 0 {
				i = s.peek()
			} else {
				i = -1
			}
			k := idx

			a := hg[j] * (k - i - 1)
			if a > res {
				res = a
			}
		}
		if len(s) == 0 || h > hg[s.peek()] {
			s.push(idx)
		} else if len(s) > 0 && h == hg[s.peek()] {
			s.pop()
			s.push(idx)
		}
	}

	for len(s) > 0 {
		j := s.pop()
		var i int
		if len(s) > 0 {
			i = s.peek()
		} else {
			i = -1
		}
		k := j + 1

		a := hg[j] * (k - i - 1)
		if a > res {
			res = a
		}
	}
	return res
}

// the idea is to use maxRectUnderHistogram to get max area for rects
// with the bottom horizontal line on the current row
// the histogram from one row to the next can computed in O(numColumns)
// so overall algorithm has runtime O(numRows x numColumns)
func MaximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	n := len(matrix[0])
	hg := make([]int, n)
	res := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				hg[j] = hg[j] + 1
			} else {
				hg[j] = 0
			}
		}
		a := maxRectUnderHistogram(hg)
		if res < a {
			res = a
		}
	}
	return res
}
