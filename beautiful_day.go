package main

/*
 * Complete the 'beautifulDays' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER i
 *  2. INTEGER j
 *  3. INTEGER k
 */

func beautifulDays(i int32, j int32, k int32) int32 {
	dayCount := int32(0)
	for x := i; x <= j; x++ {
		diff := abs(x - getReverse(x))
		if diff%k == 0 {
			dayCount++
		}
	}
	return dayCount
}

func getReverse(i int32) int32 {
	rev := int32(0)
	for i > 0 {
		digit := i % 10
		rev = (10 * rev) + digit
		i /= 10
	}

	return rev
}

//func abs(i int32) int32 {
//	if i < 0 {
//		return -i
//	}
//
//	return i
//}

//func main() {
//	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
//
//	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
//	checkError(err)
//
//	defer stdout.Close()
//
//	writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)
//
//	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")
//
//	iTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
//	checkError(err)
//	i := int32(iTemp)
//
//	jTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
//	checkError(err)
//	j := int32(jTemp)
//
//	kTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
//	checkError(err)
//	k := int32(kTemp)
//
//	result := beautifulDays(i, j, k)
//
//	fmt.Fprintf(writer, "%d\n", result)
//
//	writer.Flush()
//}
