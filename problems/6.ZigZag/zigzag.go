package zigzag

import "bytes"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([][]byte, numRows)
	curr := 0 // current row
	var dir int
	for i := 0; i < len(s); i++ {
		rows[curr] = append(rows[curr], s[i])
		if curr == 0 {
			dir = 1 // going down
		} else if curr == numRows-1 {
			dir = -1 // going up
		}
		curr += dir
	}
	var out bytes.Buffer
	for _, row := range rows {
		out.Write(row)
	}
	return out.String()
}
