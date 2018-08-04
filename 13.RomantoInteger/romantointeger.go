package RomantoInteger

func romanToInt(s string) int {
	rule := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	idx, num := 0, 0
	for idx < len(s) {
		if idx+1 < len(s) && (s[idx] == 'I' || s[idx] == 'X' || s[idx] == 'C') {
			sub := s[idx : idx+2]
			if v, ok := rule[sub]; ok {
				num += v
				idx += 2
				continue
			}
		}
		sub := s[idx : idx+1]
		num += rule[sub]
		idx++
	}
	return num
}
