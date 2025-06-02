package main

/*
 * Complete the 'pickingNumbers' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func pickingNumbers(a []int32) int32 {
	maxLength := int32(0)
	frequencyMap := make(map[int32]int32)
	distinctNumbers := make([]int32, 0)

	for i := 0; i < len(a); i++ {
		if _, ok := frequencyMap[a[i]]; !ok {
			distinctNumbers = append(distinctNumbers, a[i])
		}

		frequencyMap[a[i]]++
	}

	for j := 0; j < len(distinctNumbers); j++ {
		currentVal := distinctNumbers[j]
		currentFreq := frequencyMap[currentVal]
		if val, ok := frequencyMap[currentVal+1]; ok {
			currentFreq += val
		}

		if maxLength < currentFreq {
			maxLength = currentFreq
		}
	}

	return maxLength
}

//func main() {
//	a := []int32{1, 1, 2, 2, 4, 4, 5, 5, 5}
//	result := pickingNumbers(a)
//	fmt.Printf("%d\n", result)
//
//	a = []int32{4, 6, 5, 3, 3, 1}
//	result = pickingNumbers(a)
//	fmt.Printf("%d\n", result)
//
//	a = []int32{1, 2, 2, 3, 1, 2}
//	result = pickingNumbers(a)
//	fmt.Printf("%d\n", result)
//}
