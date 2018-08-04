package ExcelSheetColumnTitle

func convertToTitle(n int) string {
	res := ""
	for n != 0 {
		res = string((n-1)%26+65) + res
		n = (n - 1) / 26
	}
	return res
}
