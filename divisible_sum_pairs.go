package main

/*
 * Complete the 'divisibleSumPairs' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER k
 *  3. INTEGER_ARRAY ar
 */

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	matchCount := int32(0)
	for i := 0; i < len(ar)-1; i++ {
		for j := i + 1; j < len(ar); j++ {
			if (ar[i]+ar[j])%k == 0 {
				matchCount++
			}
		}
	}

	return matchCount
}

//func main() {
//	n := int32(6)
//	k := int32(3)
//	ar := []int32{1, 3, 2, 6, 1, 2}
//
//	result := divisibleSumPairs(n, k, ar)
//
//	fmt.Printf("%d\n", result)
//}
