package main

import "fmt"

func MergeSortedArray(first, second []int) []int {
	result := make([]int, 0)
	x, y := 0, 0

	for x < len(first) || y < len(second) {
		if x < len(first) && y < len(second) {
			fmt.Println("Both within range ", first[x], " = ", second[y])
			if first[x] <= second[y] {
				fmt.Println("x <= y")
				result = append(result, first[x])
				x++
			} else if first[x] > second[y] {
				fmt.Println("x > y")
				result = append(result, second[y])
				y++
			}
		} else if x < len(first) {
			fmt.Println("x is remaining ", first[x])
			result = append(result, first[x])
			x++
		} else {
			fmt.Println("y is remaining ", second[y])
			result = append(result, second[y])
			y++
		}
	}

	return result
}
