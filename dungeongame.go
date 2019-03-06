package leet

// model as a digraph: cells in MxN grid are the vertices
// edges are cell a -> cell b if b is rightward or downward from a
// cost of edge a -> b is -value of b (positive b means health so a negative cost)

// since knight moves only downward or rightward => no cycles => no negative cost cycles
// use Bellman-Ford to compute single source shortest paths from (0,0)

// there are |E| = 2 (M-1)(N-1) + (M-1) + (N-1) edges
// there are |V| = M N vertices

// => runtime: O(M^2 N^2)

// this won't work because you need to keep your health > 0 at every step
// Bellman-Ford algorithm might dip it below zero

func CalculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 {
		return 0
	}

	m := len(dungeon)
	n := len(dungeon[0])

	d := make([][]int, m)

	for i := 0; i < m;i++ {
		d[i] = make([]int, n)
		for j := 0; j < n; j++ {
			d[i][j] = 1 << 31 - 1
		}
	}
	d[0][0] = 0

	done := false

	for !done {
		done = true
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if i > 0 && d[i-1][j] - dungeon[i][j] < d[i][j] {
					done = false
					d[i][j] = d[i-1][j] - dungeon[i][j]
				}

				if j > 0 && d[i][j-1] - dungeon[i][j] < d[i][j] {
					done = false
					d[i][j] = d[i][j-1] - dungeon[i][j]
				}
			}
		}
	}

	cost := -d[m-1][n-1] + dungeon[0][0]
	if cost < 0 {
		return 0
	}
	return -cost
}


