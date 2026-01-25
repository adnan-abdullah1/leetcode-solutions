func myPow(x float64, n int) float64 {
	ans := 1.0

	nn := n
	if nn < 0 {
		nn = -nn
	}

	for nn > 0 {
		if nn%2 == 0 {
			x = x * x
			nn = nn / 2
		} else {
			ans = ans * x
			nn = nn - 1
		}
	}

	if n < 0 {
		return 1 / ans
	}
	return ans
}
