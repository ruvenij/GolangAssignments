package main

import (
	"fmt"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
	//var minSum int64
	//var maxSum int64
	//slices.Sort(arr)
	//
	//minSum = int64(arr[0]) + int64(arr[1]) + int64(arr[2]) + int64(arr[3])
	//maxSum = int64(arr[1]) + int64(arr[2]) + int64(arr[3]) + int64(arr[4])
	//
	//fmt.Println(minSum, maxSum)

	sum := int64(0)
	minVal, maxVal := arr[0], arr[0]
	for _, v := range arr {
		sum += int64(v)

		if v > maxVal {
			maxVal = v
		}

		if v < minVal {
			minVal = v
		}
	}

	minSum := sum - int64(maxVal)
	maxSum := sum - int64(minVal)

	fmt.Println(minSum, maxSum)
}

//func main() {
//	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
//
//	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
//
//	var arr []int32
//
//	for i := 0; i < 5; i++ {
//		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
//		checkError(err)
//		arrItem := int32(arrItemTemp)
//		arr = append(arr, arrItem)
//	}
//
//	miniMaxSum(arr)
//}
