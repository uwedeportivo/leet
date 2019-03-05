package leet

type jumpGraph []int

func (jg jumpGraph) isEdge(u, v int) bool {
	if v <= u || u < 0 || v < 0 || u >= len(jg) || v >= len(jg) {
		return false
	}

	return u + jg[u] >= v
}

type fifo []int

func (s *fifo) push(v int) {
	*s = append(*s, v)
}

func (s *fifo) pop() int {
	res:=(*s)[0]
	*s=(*s)[1:]
	return res
}

func JumpGame(nums []int) int {
	jg := jumpGraph(nums)
	bfs := make(fifo, 0)

	ds := make([]int, len(nums))
	for i, _ := range ds {
		ds[i] = -1
	}
	ds[0] = 0
	bfs.push(0)

	for len(bfs) > 0 {
		u := bfs.pop()

		for v := u + 1; jg.isEdge(u, v); v++ {
			if ds[v] == -1 {
				ds[v] = ds[u] + 1
				bfs.push(v)
			}
		}
	}

	return ds[len(nums) - 1]
}
