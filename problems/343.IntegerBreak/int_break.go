package intbreak

func integerBreak(n int) int {
	switch n {
	case 2:
		return 1
	case 3:
		return 2
	case 4:
		return 4
	}
	prod := 1
	for n > 4 {
		n -= 3
		prod *= 3
	}
	return prod * n
}
