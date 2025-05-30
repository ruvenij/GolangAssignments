package main

/*
 * Complete the 'birthday' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY s
 *  2. INTEGER d
 *  3. INTEGER m
 */

func birthday(s []int32, d int32, m int32) int32 {
	blockCount := int32(0)
	for i := 0; i <= len(s)-int(m); i++ {
		blockSum := int32(0)
		for j := 0; j < int(m); j++ {
			blockSum += s[i+j]
		}

		if blockSum == d {
			blockCount++
		}
	}

	return blockCount
}

//func main() {
//	s := []int32{2, 2, 1, 3, 2}
//	d := int32(4)
//	m := int32(2)
//	result := birthday(s, d, m)
//	fmt.Printf("%d\n", result)
//}
