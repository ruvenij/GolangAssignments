package main

import (
	"math"
)

/*
 * Complete the 'kangaroo' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER x1
 *  2. INTEGER v1
 *  3. INTEGER x2
 *  4. INTEGER v2
 */

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if x1 > x2 && v1 > v2 {
		return "NO"
	}

	if x1 < x2 && v1 < v2 {
		return "NO"
	}

	if v1 == 0 || v2 == 0 {
		return "NO"
	}

	if x1 != x2 && v1 == v2 {
		return "NO"
	}

	if x1 == x2 && v1 != v2 {
		return "NO"
	}

	if x1 == x2 && v1 == v2 {
		return "YES"
	}

	currentFirstLocation := x1
	currentSecondLocation := x2
	for {
		currentFirstLocation += v1
		currentSecondLocation += v2

		//fmt.Println("LOC ", currentFirstLocation, currentSecondLocation)
		if currentFirstLocation == currentSecondLocation {
			return "YES"
		}

		if currentFirstLocation < 0 || currentSecondLocation < 0 {
			return "NO"
		}

		if currentFirstLocation > math.MaxInt32 || currentSecondLocation > math.MaxInt32 {
			return "NO"
		}
	}

	return "NO"
}

//func main() {
//	result := kangaroo(0, 3, 4, 2)
//	//result := kangaroo(21, 6, 47, 3)
//	fmt.Printf("%s\n", result)
//}
