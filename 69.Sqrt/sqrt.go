package Sqrt

import (
	"math"
)

func mySqrt(x int) int {
	k := 1.0
	y := 0.5
	for math.Abs(y*(k+float64(x)/k)-k) >= 0.00001 {
		k = y * (k + float64(x)/k)
	}
	return int(k)
}
