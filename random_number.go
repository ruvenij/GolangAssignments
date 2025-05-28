package main

import "math/rand"

func GetRandomNumber(start, end int) int {
	randNum := rand.Intn(end - start + 1)
	return start + randNum
}
