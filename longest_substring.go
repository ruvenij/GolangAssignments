package main

//func lengthOfLongestSubstring(s string) int {
//	runes := []rune(s)
//	maxLenOfSubstring := 0
//	for i := 0; i < len(runes); i++ {
//		currentSubStringLen := 0
//		handledRunes := make(map[rune]struct{})
//		for j := i; j < len(runes); j++ {
//			_, ok := handledRunes[runes[j]]
//			if ok {
//				break
//			}
//
//			handledRunes[runes[j]] = struct{}{}
//			currentSubStringLen++
//		}
//
//		if maxLenOfSubstring < currentSubStringLen {
//			maxLenOfSubstring = currentSubStringLen
//		}
//	}
//	return maxLenOfSubstring
//}

// better approach : sliding window
func lengthOfLongestSubstring(s string) int {
	runes := []rune(s)
	maxLenOfSubstring := 0
	start := 0
	charMap := make(map[rune]int)
	for i := 0; i < len(runes); i++ {
		if lastSeen, ok := charMap[runes[i]]; ok && lastSeen >= start {
			start = lastSeen + 1
		}

		charMap[runes[i]] = i
		currentSubStringLen := i - start + 1
		if maxLenOfSubstring < currentSubStringLen {
			maxLenOfSubstring = currentSubStringLen
		}
	}
	return maxLenOfSubstring
}

//func main() {
//	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
//	fmt.Println(lengthOfLongestSubstring("bbbbb"))
//	fmt.Println(lengthOfLongestSubstring("pwwkew"))
//}
