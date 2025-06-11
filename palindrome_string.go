package main

func isPalindrome(s string) bool {
	start, end := 0, len(s)-1
	for start <= end {
		if s[start] != s[end] {
			return false
		}

		start++
		end--
	}

	return true
}

func isPalindromeFromInteger(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	reversed := 0
	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}

	if x == reversed || x == reversed/10 {
		return true
	}

	return false
}

//func main() {
//	fmt.Println(isPalindrome("abcba"))
//	fmt.Println(isPalindrome("abba"))
//	fmt.Println(isPalindrome("abbjjda"))
//
//	fmt.Println(isPalindromeFromInteger(121))
//	fmt.Println(isPalindromeFromInteger(10))
//}
