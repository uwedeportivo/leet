package leet

// https://www.youtube.com/watch?v=oDhu5uGq_ic

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

	profits := make([][]int, len(xs))
	for i := 0; i < len(profits); i++ {
		profits[i] = make([]int, maxNumTransactions+1)
	}

	for numT := 1; numT <= maxNumTransactions; numT++ {
		maxDiff := -xs[0]

		for day := 1; day < len(xs); day++ {
			ifIDoNothing := profits[day-1][numT]
			ifISell := xs[day] + maxDiff
			maxDiff = max(maxDiff, profits[day][numT-1] - xs[day])

			profits[day][numT] = max(ifIDoNothing, ifISell)
		}
	}

	return profits[len(xs)-1][maxNumTransactions]
}
