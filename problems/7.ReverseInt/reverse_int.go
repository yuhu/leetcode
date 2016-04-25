package reverseint

// reverse reverses the digits of x, and it returns 0 when the reversion
// overflows an int32.
func reverse(x int) int {
	neg := false
	if x < 0 {
		neg = true
		x = -x
	}
	r := 0
	for x != 0 {
		d := x % 10
		r = r*10 + d
		x /= 10
	}
	if neg {
		r = -r
	}
	if r >= (1<<31-1) || r < -(1<<31) {
		return 0
	}
	return r
}
