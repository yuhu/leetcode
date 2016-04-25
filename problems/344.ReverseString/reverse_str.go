package reversestr

func reverseString(s string) string {
	r := []rune(s)
	for i := range r {
		j := len(r) - 1 - i
		if j <= i {
			break
		}
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
