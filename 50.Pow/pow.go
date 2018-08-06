package Pow

func myPow(x float64, n int) float64 {
	if x < -100.0 || x > 100.0 || (x == -1 && n%2 == 1) {
		return -1
	}
	if n == 0 || x == 1.0 || (x == -1 && n%2 == 0) {
		return 1
	}
	res := x
	isNeg := false
	if n < 0 {
		isNeg = true
		if n > (-1 << 31) {
			n = -n
		} else {
			n = 1<<31 - 1
		}
	}
	for i := 1; i < n; i++ {
		if (res * x) > 1<<31-1 {
			return 0
		}
		res *= x
	}
	if isNeg {
		res = 1 / res
	}

	return res
}
