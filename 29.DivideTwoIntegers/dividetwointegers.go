package DivideTwoIntegers

func divide(dividend int, divisor int) int {
	if divisor == 0 || (dividend == -1<<31 && divisor == -1) {
		return 1<<31 - 1
	}
	n := dividend % divisor
	res := (dividend - n) / divisor
	return res
}
