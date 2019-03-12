package leet

// Removes adjacent duplicates
func delAdjacentDuplcates(xs []int) []int {
	if len(xs) == 0 {
		return xs
	}
	ys := make([]int, 0, len(xs))
	prev := xs[0]
	ys = append(ys, prev)

	for i := 1; i < len(xs); i++ {
		if xs[i] != prev {
			ys = append(ys, xs[i])
			prev = xs[i]
		}
	}
	return ys
}

// Returns end (excluding) index of monotone sequence
// starting at start index.
// Assumes adjacent duplicates have been removed.
func monotoneEnd(xs []int, start int) int {
	if start >= len(xs) - 1 {
		return len(xs)
	}

	initialDirection := xs[start] < xs[start+1]

	for i := start + 2; i < len(xs); i++ {
		direction := xs[i-1] < xs[i]

		if direction != initialDirection {
			return i
		}
	}
	return len(xs)
}

// Collapses descends and ascends into just peaks and valleys.
// Assumes adjacent duplicates have been removed.
func collapseHills(xs []int) []int {
	if len(xs) == 0 || len(xs) == 1 {
		return xs
	}

	ys := make([]int, 0, len(xs))
	ys = append(ys, xs[0])
	start := 0
	for {
		end := monotoneEnd(xs, start)
		if start < end-1 {
			ys = append(ys, xs[end-1])
		}
		start = end-1
		if end == len(xs) {
			break
		}
	}
	return ys
}

// Deletes decreasing prefix.
// Assumes adjacent duplicates have been removed.
func trimDecreasingPrefix(xs []int) []int {
	if len(xs) == 0 || len(xs) == 1 {
		return xs
	}

	ys := xs

	if ys[0] > ys[1] {
		ys = ys[1:]
	}

	return ys
}

// Deletes decreasing suffix.
// Assumes adjacent duplicates have been removed.
func trimDecreasingSuffix(xs []int) []int {
	if len(xs) == 0 || len(xs) == 1 {
		return xs
	}

	ys := xs
	n := len(ys)

	if ys[n-2] > ys[n-1] {
		ys = ys[:n-1]
	}

	return ys
}

// Returns profit from the best single transaction
// within start and end by visiting all possible
// buys and sells pairings.
func bestTransaction(xs []int, start, end int) int {
	if start == end || start == end - 1 {
		return 0
	}

	mp := -1
	for i := start; i < end; i+=2 {
		for j := i+1; j < end; j+=2 {
			if xs[i] < xs[j] {
				p := xs[j] - xs[i]
				if p > mp {
					mp = p
				}
			}
		}
	}
	return mp
}

func maxProfitImpl(k int, xs []int, start, end int, profits [][][]int) int {
	if k == 0 {
		profits[start][end][k] = 0
		return 0
	}
	if start == end {
		profits[start][end][k] = 0
		return 0
	}
	if end - start == 2 {
		profits[start][end][k] = xs[end-1]-xs[start]
		return xs[end-1]-xs[start]
	}

	if profits[start][end][k] >= 0 {
		return profits[start][end][k]
	}

	if k >= (end - start) / 2 {
		sum := 0
		for i := start; i < end; i+=2 {
			sum += xs[i+1] - xs[i]
		}
		profits[start][end][k] = sum
		return sum
	}

	bt := bestTransaction(xs, start, end)
	if k == 1 {
		profits[start][end][k] = bt
		return bt
	}

	mp := -1

	for i := start+2; i < end; i+=2 {
		pl := xs[i-1] - xs[start]
		pr := maxProfitImpl(k-1, xs, i, end, profits)

		if mp < pl + pr {
			mp = pl + pr
		}
	}

	pp := maxProfitImpl(k, xs, start+2, end, profits)
	if mp < pp {
		mp = pp
	}

	pp = maxProfitImpl(k-1, xs, start, end, profits)
	if mp < pp {
		mp = pp
	}

	profits[start][end][k] = mp
	return mp
}

func MaxProfit(maxNumTransactions int, prices []int) int {
	if maxNumTransactions == 0 {
		return 0
	}

	xs := prices

	xs = delAdjacentDuplcates(xs)
	xs = collapseHills(xs)
	xs = trimDecreasingPrefix(xs)
	xs = trimDecreasingSuffix(xs)

	if len(xs) == 0 || len(xs) == 1 {
		return 0
	}

	if maxNumTransactions >= len(xs) / 2 {
		sum := 0
		for i := 0; i < len(xs); i+=2 {
			sum += xs[i+1] - xs[i]
		}
		return sum
	}

	profits := make([][][]int, len(xs))
	for i := 0; i < len(xs); i++ {
		profits[i] = make([][]int, len(xs) + 1)

		for j := 0; j < len(profits[i]); j++ {
			profits[i][j] = make([]int, maxNumTransactions + 1)
			for k := 0; k < len(profits[i][j]); k++ {
				profits[i][j][k] = -1
			}
		}
	}

	return maxProfitImpl(maxNumTransactions, xs, 0, len(xs), profits)
}
