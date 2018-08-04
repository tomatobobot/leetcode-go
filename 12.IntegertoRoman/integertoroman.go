package IntegertoRoman

func intToRoman(num int) string {
	res := ""
	for num > 0 {
		switch {
		case num >= 1000:
			res += "M"
			num -= 1000
			break
		case num > 899 && num <= 999:
			res += "CM"
			num -= 900
			break
		case num > 499 && num <= 899:
			res += "D"
			num -= 500
			break
		case num > 399 && num <= 499:
			res += "CD"
			num -= 400
			break
		case num > 99 && num <= 399:
			res += "C"
			num -= 100
			break
		case num > 89 && num <= 99:
			res += "XC"
			num -= 90
			break
		case num > 49 && num < 90:
			res += "L"
			num -= 50
			break
		case num > 39 && num <= 49:
			res += "XL"
			num -= 40
			break
		case num > 9 && num <= 39:
			res += "X"
			num -= 10
			break
		case num == 9:
			res += "IX"
			num = 0
			break
		case num > 4 && num < 9:
			res += "V"
			num -= 5
			break
		case num == 4:
			res += "IV"
			num = 0
			break
		case num <= 3:
			for i := 0; i < num; i++ {
				res += "I"
			}
			num = 0
			break
		}
	}
	return res
}
