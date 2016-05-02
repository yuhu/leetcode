package palindromenum

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var digits []int
	for x > 0 {
		digits = append(digits, x%10)
		x /= 10
	}
	for i, d := range digits {
		j := len(digits) - 1 - i
		if d != digits[j] {
			return false
		}
		if j <= i {
			break
		}
	}
	return true
}
