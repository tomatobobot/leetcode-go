package ImplementstrStr

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	haystacklen := len(haystack)
	needlelen := len(needle)
	for i := 0; i < haystacklen; i++ {
		if i+needlelen > haystacklen {
			break
		}
		if haystack[i:i+needlelen] == needle {
			return i
		}
	}
	return -1
}
