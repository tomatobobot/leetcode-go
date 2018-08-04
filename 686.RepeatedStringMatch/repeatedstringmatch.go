package RepeatedStringMatch

import (
	"bytes"
	"strings"
)

func repeatedStringMatch(A string, B string) int {
	count := 0
	var buf bytes.Buffer
	for {
		buf.WriteString(A)
		count++
		// if findSubString(buf.String(), B) {
		// 	return count
		// }
		if strings.Contains(buf.String(), B) {
			return count
		}
		if buf.Len() > 10000 {
			return -1
		}
	}
}

func findSubString(B string, A string) bool {
	idx, alen, blen := 0, len(A), len(B)
	if blen >= alen {
		for idx < blen {
			if idx+alen <= blen && A == B[idx:idx+alen] {
				return true
			}
			idx++
		}

	}
	return false
}
