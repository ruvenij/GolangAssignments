package main

import "strings"

func GetWordCount(input string) int {
	words := strings.Fields(input)
	return len(words)
}
