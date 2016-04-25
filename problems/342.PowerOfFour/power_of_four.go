package poweroffour

func isPowerOfFour(num int) bool {
	// num's binary form looks like: "1(00)*". In other words, it has exactly
	// one "1", which can occur in any of position 1, 3, 5, ...
	if num == 0 {
		return false
	}
	if num&(num-1) != 0 {
		return false
	}
	if num&0x55555555 == 0 {
		return false
	}
	return true
}
