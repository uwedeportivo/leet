package leet

func IsInterleave(s1, s2, s3 string) bool {
	n, m := len(s1), len(s2)

	if n + m != len(s3) {
		return false
	}

	if n == 0 && m == 0 {
		return true
	}

	if n == 0 {
		return s2 == s3
	}

	if m == 0 {
		return s1 == s3
	}

	grid := make([][]bool, n+1)

	for i := 0; i <= n; i++ {
		grid[i] = make([]bool, m+1)
	}

	grid[0][0] = s1[0] == s3[0] || s2[0] == s3[0]

	if !grid[0][0] {
		return false
	}

	for i := 1; i <= n; i++ {
		grid[i][0] = s1[i-1] == s3[i-1] && grid[i-1][0]
	}

	for j := 1; j <= m; j++ {
		grid[0][j] = s2[j-1] == s3[j-1] && grid[0][j-1]
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			lb := s1[i-1] == s3[i+j-1] && grid[i-1][j]
			rb := s2[j-1] == s3[i+j-1] && grid[i][j-1]
			grid[i][j] = lb || rb
		}
	}

	return grid[n][m]
}
