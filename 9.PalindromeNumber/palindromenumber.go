package PalindromeNumber

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	b := 1
	for x/b >= 10 {
		b *= 10
	}
	for b >= 10 {
		if x/b != x%10 {
			return false
		}
		x, b = (x%b)/10, b/100
	}
	return true
}
