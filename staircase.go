package main

import "fmt"

/*
 * Complete the 'staircase' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func staircase(n int32) {
	for i := int32(0); i < n; i++ {
		for j := int32(0); j < n; j++ {
			if n-i-1 <= j {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

//func main() {
//	n := int32(4)
//	staircase(n)
//}
