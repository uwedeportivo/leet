package leet


func GenerateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}

	var res []string
	for k := 1; k <=n; k++ {
		left := GenerateParenthesis(k - 1)
		right := GenerateParenthesis(n - k)

		for _, ls := range left {
			for _, rs := range right {
				res = append(res, "(" + ls + ")" + rs)
			}
		}
	}
	return res
}

