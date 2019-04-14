package leet

import (
	"sort"
	"strings"
)

func indices(str, substr string) []int {
	var res []int

	n := len(substr)
	offset := 0
	idx := strings.Index(str, substr)

	for idx != -1 && len(str) >= n {
		res = append(res, idx + offset)

		offset += idx + 1
		str = str[idx + 1:]

		idx = strings.Index(str, substr)
	}
	return res
}

type wordNode struct {
	value string
	index int
	tag int
}

type byIndex []*wordNode

func (a byIndex) Len() int           { return len(a) }
func (a byIndex) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byIndex) Less(i, j int) bool { return a[i].index < a[j].index }

func tagMap(ws []string) (map[string]int, []int) {
	tm := make(map[string]int)
	ps := make([]int, 0, len(ws))
	for i, w := range ws {
		t, ok := tm[w]
		if !ok {
			tm[w] = i
			t = i
		}
		ps = append(ps, t)
	}
	sort.Ints(ps)
	return tm, ps
}

func nextBlock(chain []*wordNode, wl int) int {
	if len(chain) == 0 {
		return 0
	}

	k := 1
	for k < len(chain) && chain[k].index - chain[k-1].index == wl {
		k++
	}
	return k
}

func setEqual(xs, ys []int) bool {

	if len(xs) != len(ys) {
		return false
	}

	for i := 0; i < len(xs); i++ {
		if xs[i] != ys[i] {
			return false
		}
	}
	return true
}

func processBlock(block []*wordNode, ps []int) []int {
	var res []int

	if len(block) < len(ps) {
		return nil
	}

	cs := make([]int, len(ps))

	l, r := 0, len(ps)

	for r <= len(block) {
		for i := l; i < r; i++ {
			cs[i-l] = block[i].tag
		}
		sort.Ints(cs)

		if setEqual(ps, cs) {
			res = append(res, block[l].index)
		}

		l, r = l + 1, r + 1
	}

	return res
}

func processChain(chain []*wordNode, ps []int, wl int) []int {
	if len(chain) < len(ps) {
		return nil
	}

	sort.Sort(byIndex(chain))

	var res []int

	be := nextBlock(chain, wl)
	for be > 0 {
		idxs := processBlock(chain[:be], ps)
		res = append(res, idxs...)
		chain = chain[be:]
		be = nextBlock(chain, wl)
	}

	return res
}

func FindSubstring(str string, ws []string) []int {
	if len(ws) == 0 {
		return nil
	}

	var res []int

	wl := len(ws[0])

	tags, ps := tagMap(ws)
	processed := make(map[string]bool)
	chains := make([][]*wordNode, wl)

	for _, w := range ws {
		if !processed[w] {
			tag := tags[w]
			processed[w] = true
			idxs := indices(str, w)

			for _, idx := range idxs {
				wn := &wordNode{
					value: w,
					index: idx,
					tag: tag,
				}
				c := idx % wl
				chains[c] = append(chains[c], wn)
			}
		}
	}

	for _, chain := range chains {
		idxs := processChain(chain, ps, wl)
		res = append(res, idxs...)
	}

	return res
}
