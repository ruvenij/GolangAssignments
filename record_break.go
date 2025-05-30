package main

/*
 * Complete the 'breakingRecords' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY scores as parameter.
 */

func breakingRecords(scores []int32) []int32 {
	minScore := scores[0]
	maxScore := scores[0]
	minBreaks := 0
	maxBreaks := 0

	for _, val := range scores {
		if val < minScore {
			minScore = val
			minBreaks++
		}

		if val > maxScore {
			maxScore = val
			maxBreaks++
		}
	}

	return []int32{int32(maxBreaks), int32(minBreaks)}

}

//func main() {
//	scores := []int32{10, 5, 20, 20, 4, 5, 2, 25, 1}
//	result := breakingRecords(scores)
//
//	for i, resultItem := range result {
//		fmt.Printf("%d", resultItem)
//
//		if i != len(result)-1 {
//			fmt.Printf(" ")
//		}
//	}
//
//	fmt.Printf("\n")
//}
