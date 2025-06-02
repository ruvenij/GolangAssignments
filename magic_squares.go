package main

import (
	"fmt"
	"math"
)

func formingMagicSquare(s [][]int32) int32 {
	magicSquares := [][]int32{
		{8, 1, 6, 3, 5, 7, 4, 9, 2},
		{6, 1, 8, 7, 5, 3, 2, 9, 4},
		{4, 9, 2, 3, 5, 7, 8, 1, 6},
		{2, 9, 4, 7, 5, 3, 6, 1, 8},
		{8, 3, 4, 1, 5, 9, 6, 7, 2},
		{4, 3, 8, 9, 5, 1, 2, 7, 6},
		{6, 7, 2, 1, 5, 9, 8, 3, 4},
		{2, 7, 6, 9, 5, 1, 4, 3, 8},
	}
	fmt.Println("MagicSquares:", magicSquares)

	flattenedInput := make([]int32, 0)
	for _, v := range s {
		flattenedInput = append(flattenedInput, v...)
	}

	minCost := int32(math.MaxInt32)

	for i := 0; i < len(magicSquares); i++ {
		currentCost := int32(0)

		for j := 0; j < len(magicSquares[i]); j++ {
			diff := abs(magicSquares[i][j] - flattenedInput[j])
			currentCost += diff
		}

		if minCost > currentCost {
			minCost = currentCost
		}
	}

	return minCost
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

//func main() {
//	s := [][]int32{
//		{4, 9, 2},
//		{3, 5, 7},
//		{8, 1, 5},
//	}
//	fmt.Println(formingMagicSquare(s)) // Output: 1
//}
