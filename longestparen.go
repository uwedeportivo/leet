package leet

func LongestValidParen(s string) int {
	l := 0 // start of current valid
	h := 0 // height of current valid
	max := 0 // longest valid so far
	n := len(s)
	z := -1 // last zero in current valid

	for k := 0; k < n; k++ {
		if s[k] == '(' {
			h++
		} else {
			h--
		}
		if h == 0 && s[k] == ')' {
			z = k
			if max < z + 1 - l {
				max = z + 1 - l
			}
		} else if h < 0 {
			if max < k - l {
				max = k - l
			}
			l = k + 1
			h = 0
		}
	}

	if h > 0 {
		// we have terminated before going back to zero
		// inside might be some valid string candidates

		if max < z + 1 - l {
			max = z + 1 - l
		}

		insideMax := LongestValidParen(s[z+2:])

		if max < insideMax {
			max = insideMax
		}
	}
	return max
}
