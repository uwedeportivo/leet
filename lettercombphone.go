package leet

type wheel struct {
	val byte
	max byte
}

func digit2Rune(d int, offset wheel) byte {
	if d >= 2 && d <= 6 {
		return byte(97 + (d - 2) * 3 + int(offset.val))
	} else if d == 7 {
		return byte(112 + offset.val)
	} else if d == 8 {
		return byte(116 + offset.val)
	}
	return byte(119 + offset.val)
}

func incr(offsets []wheel) bool {
	n := len(offsets)
	i := n - 1
	for i >= 0 {
		offsets[i].val = offsets[i].val + 1
		if offsets[i].val == offsets[i].max {
			for j := n - 1; j >= i; j-- {
				offsets[j].val = 0
			}
			i--
		} else {
			return false
		}
	}
	return i < 0
}


func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	n := len(digits)

	nums := make([]int, n)

	total := 1
	for i, r := range digits {
		nums[i] = int(r) - 50 + 2
		total = total * 3
	}

	offsets := make([]wheel, n)

	for i := 0; i < n; i++ {
		if nums[i] == 7 || nums[i] == 9 {
			offsets[i].max = 4
		} else {
			offsets[i].max = 3
		}
	}

	res := make([]string, 0, total)

	for {
		ss := make([]byte, n)

		for i := 0; i < n; i++ {
			ss[i] = digit2Rune(nums[i], offsets[i])
		}

		res = append(res, string(ss))

		if incr(offsets) {
			break
		}
	}

	return res
}

