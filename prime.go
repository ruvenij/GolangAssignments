package main

import "math"

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}

	sqrVal := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrVal; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
