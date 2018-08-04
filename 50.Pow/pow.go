package Pow

func myPow(x float64, n int) float64 {
	if x < -100.0 || x > 100.0 {
		return -1
	}
	if n == 0 || x == 1.0 {
		return 1
	}
	if x == -1 {
		return -1
	}
	res := x
	isNeg := false
	if n < 0 {
		n = -n
		isNeg = true
	}
	for i := 1; i < n; i++ {
		res *= x
	}
	if isNeg {
		res = 1 / res
	}

	return res
}
