package LongestCommonPrefix

func longestCommonPrefix(strs []string) string {
	res := ""
	if len(strs) > 0 {
		first := strs[0]
		for i := range first {
			for _, word := range strs[1:] {
				if i > len(word)-1 || first[i:i+1] != word[i:i+1] {
					goto LOOP
				}
			}
			res += first[i : i+1]
		}
	LOOP:
	}
	return res
}
