package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	var pos, neg, zero int
	for _, val := range arr {
		switch {
		case val > 0:
			pos++
		case val < 0:
			neg++
		default:
			zero++
		}
	}

	fmt.Printf("%.6f\n", float64(pos)/float64(len(arr)))
	fmt.Printf("%.6f\n", float64(neg)/float64(len(arr)))
	fmt.Printf("%.6f\n", float64(zero)/float64(len(arr)))
}

//func main() {
//	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
//
//	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//	checkError(err)
//	n := int32(nTemp)
//
//	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
//
//	var arr []int32
//
//	for i := 0; i < int(n); i++ {
//		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
//		checkError(err)
//		arrItem := int32(arrItemTemp)
//		arr = append(arr, arrItem)
//	}
//
//	plusMinus(arr)
//}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
