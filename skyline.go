package leet

type rect struct {
	left int
	height int
	right int
}

type skypart struct {
	start int
	height int
}

func appendSkypart(sps []skypart, sp skypart) []skypart {
	n := len(sps)
	if n == 0 {
		return append(sps, sp)
	}

	if sps[n-1].height == sp.height {
		return sps
	}

	if sps[n-1].start == sp.start {
		sps[n-1].height = max(sps[n-1].height, sp.height)
		return sps
	}

	return append(sps, sp)
}

func mergeSkylines(lxs, rxs []skypart) []skypart {
	leftHeight, rightHeight := 0, 0
	li, ri := 0, 0
	var sps []skypart

	for li < len(lxs) && ri < len(rxs) {
		if lxs[li].start < rxs[ri].start {
			leftHeight = lxs[li].height
			height := max(leftHeight, rightHeight)
			sps = appendSkypart(sps, skypart{lxs[li].start, height})
			li++
		} else if lxs[li].start == rxs[ri].start {
			leftHeight = lxs[li].height
			rightHeight = rxs[ri].height
			height := max(leftHeight, rightHeight)
			sps = appendSkypart(sps, skypart{rxs[ri].start, height})
			li++
			ri++
		} else {
			rightHeight = rxs[ri].height
			height := max(leftHeight, rightHeight)
			sps = appendSkypart(sps, skypart{rxs[ri].start, height})
			ri++
		}
	}

	for li < len(lxs) {
		sps = appendSkypart(sps, lxs[li])
		li++
	}

	for ri < len(rxs) {
		sps = appendSkypart(sps, rxs[ri])
		ri++
	}

	return sps
}

func computeSkyline(xs []rect) []skypart {
	if len(xs) == 1 {
		return []skypart{{xs[0].left, xs[0].height}, {xs[0].right, 0}}
	}

	m := len(xs) / 2
	lxs := computeSkyline(xs[:m])
	rxs := computeSkyline(xs[m:])

	return mergeSkylines(lxs, rxs)
}

func GetSkyline(buildings [][]int) [][]int {
	if len(buildings) == 0 {
		return nil
	}
	
	xs := make([]rect, len(buildings))

	for i := 0; i < len(buildings); i++ {
		xs[i].left = buildings[i][0]
		xs[i].right = buildings[i][1]
		xs[i].height = buildings[i][2]
	}

	sps := computeSkyline(xs)

	res := make([][]int, len(sps))

	for i := 0; i < len(sps); i++ {
		res[i] = make([]int, 2)
		res[i][0] = sps[i].start
		res[i][1] = sps[i].height
	}
	return res
}
