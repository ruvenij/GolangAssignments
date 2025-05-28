package main

func FindSum(start, end int) int {
	total := 0

	for i := start; i <= end; i++ {
		if i%2 == 0 {
			total += i
		}
	}

	return total
}
