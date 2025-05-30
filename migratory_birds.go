package main

/*
 * Complete the 'migratoryBirds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

// returns max bird type seen
func migratoryBirds(arr []int32) int32 {
	frequencyMap := make(map[int32]int32)
	maxSight := arr[0]

	for _, val := range arr {
		frequencyMap[val]++

		if frequencyMap[val] > maxSight {
			maxSight = frequencyMap[val]
		}
	}

	for i := int32(1); i <= 5; i++ {
		if frequencyMap[i] == maxSight {
			return i
		}
	}

	return -1
}

//func main() {
//	arr := []int32{1, 1, 2, 2, 3}
//	result := migratoryBirds(arr)
//	fmt.Printf("%d\n", result)
//}
