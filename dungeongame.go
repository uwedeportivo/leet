package leet

import "fmt"

// doesn't work either because final solution is not composition
// of sub solutions (ie path to previous neighbors is suboptimal
// but leads to optimal final)

const (
	infinity = 1 << 31 - 1
)

func from(hp, ap, cell int) (int, int) {
	if cell >= 0 {
		return hp, ap + cell
	}

	delta := ap + cell
	if delta >= 0 {
		return hp, delta
	}
	return hp + delta, 0
}

func minHP(dungeon, hp, ap [][]int, i, j int) (int, int) {
	if hp[i][j] > -infinity {
		return hp[i][j], ap[i][j]
	}

	nhp := -infinity
	nap := 0
	if i > 0 {
		nhp, nap = minHP(dungeon, hp, ap, i-1, j)
		nhp, nap = from(nhp, nap, dungeon[i][j])
	}
	whp := -infinity
	wap := 0
	if j > 0 {
		whp, wap = minHP(dungeon, hp, ap, i, j-1)
		whp, wap = from(whp, wap, dungeon[i][j])
	}

	if i == 2 && j == 2 {
		fmt.Println("boo")
	}

	if nhp > whp {
		hp[i][j] = nhp
		ap[i][j] = nap
	} else {
		hp[i][j] = whp
		ap[i][j] = wap
	}
	return hp[i][j], ap[i][j]
}

func CalculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 {
		return 0
	}

	m := len(dungeon)
	n := len(dungeon[0])

	hp := make([][]int, m)
	ap := make([][]int, m)

	for i := 0; i < m;i++ {
		hp[i] = make([]int, n)
		ap[i] = make([]int, n)
		for j := 0; j < n; j++ {
			hp[i][j] = -infinity
		}
	}

	if dungeon[0][0] >= 0 {
		hp[0][0] = 0
		ap[0][0] = dungeon[0][0]
	} else {
		hp[0][0] = dungeon[0][0]
		ap[0][0] = 0
	}

	h, _ :=  minHP(dungeon, hp, ap, m -1, n-1)

	if h > 0 {
		return 1
	}
	return -h + 1
}


