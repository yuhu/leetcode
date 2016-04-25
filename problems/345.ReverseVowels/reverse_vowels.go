package reversevowels

var vowels = map[rune]bool{
	'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
	'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
}

func reverseVowels(s string) string {
	b := []rune(s)
	var vowelIndices []int
	for i, r := range b {
		if vowels[r] {
			vowelIndices = append(vowelIndices, i)
		}
	}
	for i, vi := range vowelIndices {
		j := len(vowelIndices) - 1 - i
		if j <= i {
			break
		}
		vj := vowelIndices[j]
		b[vi], b[vj] = b[vj], b[vi]
	}
	return string(b)
}
