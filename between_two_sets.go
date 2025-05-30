package main

import (
	"slices"
)

/*
 * Complete the 'getTotalX' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func getTotalX(a []int32, b []int32) int32 {
	applicableNumCount := 0
	matchedForA := true
	matchedForB := true

	minNum := slices.Min(a)
	maxNum := slices.Max(b)

	for num := minNum; num <= maxNum; num++ {
		matchedForA = true
		matchedForB = true

		for i := 0; i < len(a); i++ {
			if num%a[i] != 0 {
				matchedForA = false
				break
			}
		}

		if matchedForA {
			for j := 0; j < len(b); j++ {
				if b[j]%num != 0 {
					matchedForB = false
					break
				}
			}
		}

		if matchedForA && matchedForB {
			applicableNumCount++
		}
	}

	return int32(applicableNumCount)
}

//func main() {
//	arr := []int32{2, 4}
//	brr := []int32{16, 32, 96}
//	total := getTotalX(arr, brr)
//	fmt.Printf("%d\n", total)
//}
