package main

func intToRoman(num int) string {
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbol := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""
	for i := 0; i < len(val); i++ {
		for num >= val[i] {
			result += symbol[i]
			num -= val[i]
		}
	}

	return result
}

func romanToInt(s string) int {
	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := romanMap[rune(s[len(s)-1])]
	oldVal := romanMap[rune(s[len(s)-1])]
	for i := len(s) - 2; i >= 0; i-- {
		currentVal := romanMap[rune(s[i])]
		if currentVal >= oldVal {
			total += currentVal
		} else {
			total -= currentVal
		}

		oldVal = currentVal
	}

	return total
}

//func main() {
//	fmt.Println(intToRoman(3749))
//	fmt.Println(romanToInt("LVIII"))
//	fmt.Println(romanToInt("MMMDCCXLIX"))
//}
