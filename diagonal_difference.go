package main

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func diagonalDifference(arr [][]int32) int32 {
	rowCount := len(arr)
	firstDiagonal := int32(0)
	secondDiagonal := int32(0)
	j := rowCount - 1
	for i := 0; i < rowCount; i++ {
		firstDiagonal += arr[i][i]
		secondDiagonal += arr[i][j]
		j--
	}

	diff := firstDiagonal - secondDiagonal
	if diff < 0 {
		diff = -diff
	}

	return diff
}

//func main() {
//
//	arr := [][]int32{
//		{1, 2, 3},
//		{4, 5, 6},
//		{9, 8, 9},
//	}
//
//	result := diagonalDifference(arr)
//
//	fmt.Printf("%d\n", result)
//
//}
