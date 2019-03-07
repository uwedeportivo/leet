package leet

import "fmt"

type subsetIterator struct {
	n int
	r int
	cur []int
}

func createSubsetIterator(n, r int) *subsetIterator {
	first := make([]int, r)

	for i:=0; i < r; i++ {
		first[i] = i+1
	}
	return &subsetIterator{
		n:n,
		r:r,
		cur:first,
	}
}

func (si *subsetIterator) current(buf []int) {
	copy(buf, si.cur)
}

func (si *subsetIterator) next() bool {
	l := -1
	for k := 0; k < si.r; k++ {
		if si.cur[k] < si.n - si.r + k + 1 {
			l = k
		}
	}
	if l == -1 {
		return false
	}

	c := si.cur[l]
	for i := l; i < si.r; i++ {
		si.cur[i] = c + i - l + 1
	}
	return true
}


type pathIterator struct {
	m int
	n int
	si *subsetIterator
	subsetBuf []int
}

func createPathIterator(m, n int) *pathIterator {
	return &pathIterator{
		m:m,
		n:n,
		si:createSubsetIterator(m+n, m),
		subsetBuf:make([]int, m),
	}
}

func (pi *pathIterator) current(buf []byte) {
	pi.si.current(pi.subsetBuf)

	for i := 0; i < len(buf); i++ {
		buf[i] = 0
	}

	for j := 0; j < pi.m; j++ {
		buf[pi.subsetBuf[j]-1] = 1
	}
}

func (pi *pathIterator) next() bool {
	return pi.si.next()
}

func move(hp, ap, cell int) (int, int) {
	if cell >= 0 {
		return hp, ap + cell
	}

	delta := ap + cell
	if delta >= 0 {
		return hp, delta
	}
	return hp + delta, 0
}

func needed(h int) int {
	if h > 0 {
		return 1
	}
	return -h + 1
}

func CalculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 {
		return 0
	}

	m := len(dungeon)
	n := len(dungeon[0])

	pi := createPathIterator(m-1, n-1)

	buf := make([]byte, m+n-2)

	minNeeds := 1 << 31 -1

	for {
		x, y := 0, 0
		hp, ap := move(0, 0, dungeon[x][y])
		pi.current(buf)

		for _, b := range buf {
			if b == 1 {
				x = x + 1
			} else {
				y = y + 1
			}

			if x >= len(dungeon) || y >= len(dungeon[x]) {
				fmt.Println("boom")
			}
			hp, ap = move(hp, ap, dungeon[x][y])
		}

		needs := needed(hp)
		if needs < minNeeds {
			minNeeds = needs
		}

		if !pi.next() {
			break
		}
	}

	return minNeeds
}


