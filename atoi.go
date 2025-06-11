package main

import (
	"math"
	"unicode"
)

func myAtoi(s string) int {
	startIndex := 0
	if s == "" {
		return 0
	}

	isNegative := false
	for s[startIndex] == ' ' {
		startIndex++
	}

	if startIndex >= len(s) {
		return 0
	}

	if s[startIndex] == '+' {
		startIndex++
	} else if s[startIndex] == '-' {
		startIndex++
		isNegative = true
	}

	for startIndex < len(s) && s[startIndex] == '0' {
		startIndex++
	}

	if startIndex >= len(s) {
		return 0
	}

	if !unicode.IsDigit(rune(s[startIndex])) {
		return 0
	}

	res := 0
	for startIndex < len(s) && unicode.IsDigit(rune(s[startIndex])) {
		digit := int(s[startIndex] - '0')
		if res > (math.MaxInt32-digit)/10 {
			if isNegative {
				return math.MinInt32

			}

			return math.MaxInt32
		}

		res = res*10 + int(s[startIndex]-'0')
		startIndex++
	}

	if isNegative {
		res = -res
	}

	return res
}

//func main() {
//	fmt.Println(myAtoi(" -00042"))
//	fmt.Println(myAtoi("-91283472332"))
//	fmt.Println(myAtoi("-"))
//	fmt.Println(myAtoi("2147483646"))
//}
