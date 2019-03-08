package leet

// https://twitter.com/shachaf clued me in
// here's his more elegant code:
//func calculateMinimumHP(dungeon [][]int) int {
//	rows := len(dungeon)
//	cols := len(dungeon[0])
//
//	// Could use dungeon in place; could store only one row/column at a time or something.
//	costs := make([][]int, rows)
//	for r := 0; r < rows; r++ {
//		costs[r] = make([]int, cols)
//	}
//
//	for r := rows - 1; r >= 0; r-- {
//		for c := cols - 1; c >= 0; c-- {
//			minToExit := math.MaxInt32
//			if r == rows-1 && c == cols-1 {
//				minToExit = 1
//			}
//			if r < rows-1 && costs[r+1][c] < minToExit {
//				minToExit = costs[r+1][c]
//			}
//			if c < cols-1 && costs[r][c+1] < minToExit {
//				minToExit = costs[r][c+1]
//			}
//
//			minToEnter := minToExit - dungeon[r][c]
//			if minToEnter < 1 {
//				minToEnter = 1
//			}
//
//			costs[r][c] = minToEnter
//		}
//	}
//
//	return costs[0][0]
//}

const (
	infinity = 1 << 31 -1
)
type board [][]int

func (bd board) at(x, y int) int {
	m := len(bd)
	n := len(bd[0])

	return bd[m - 1 - x][n - 1 - y]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calcMinImpl(bd board, hp [][]int, x, y int) int {
	if hp[x][y] < infinity {
		return hp[x][y]
	}

	toSouth := infinity
	if x > 0 {
		toSouth = calcMinImpl(bd, hp, x-1, y)
	}

	toEast := infinity
	if y > 0 {
		toEast = calcMinImpl(bd, hp, x, y-1)
	}

	finishing := min(toSouth, toEast)
	if bd.at(x, y) > 0 {
		finishing = max(0, finishing - bd.at(x, y))
	}

	cost := finishing
	if bd.at(x, y) < 0 {
		cost += -bd.at(x, y)
	} else if finishing == 0 {
		cost = 1
	}

	hp[x][y] = cost

	return hp[x][y]
}

func CalculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 {
		return 0
	}

	m := len(dungeon)
	n := len(dungeon[0])

	hp := make([][]int, m)
	for i := 0; i < m; i++ {
		hp[i] = make([]int, n)
		for j := 0; j < n;j++{
			hp[i][j] = infinity
		}
	}

	bd := board(dungeon)

	hp[0][0] = 1
	if bd.at(0, 0) < 0 {
		hp[0][0] = -bd.at(0, 0) + 1
	}

	return calcMinImpl(bd, hp, m-1, n-1)
}


