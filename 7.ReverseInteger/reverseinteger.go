package ReverseInteger

func reverseInteger(x int) int {
	revNum := 0
	for x != 0 {
		revNum = revNum*10 + x%10
		x /= 10
	}
	if revNum > 1<<31-1 || revNum < -1<<31-1 {
		return 0
	}
	return revNum
}
