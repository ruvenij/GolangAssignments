package main

/*
 * Complete the 'hurdleRace' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY height
 */

func hurdleRace(k int32, height []int32) int32 {
	maxHeight := height[0]
	for _, h := range height {
		if h > maxHeight {
			maxHeight = h
		}
	}

	val := maxHeight - k
	if val < 0 {
		return 0
	}

	return val
}

//func main() {
//	k := int32(4)
//	height := []int32{1, 6, 3, 5, 2}
//	result := hurdleRace(k, height)
//	fmt.Printf("%d\n", result)
//}
