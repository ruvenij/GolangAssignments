package main

/*
 * Complete the 'angryProfessor' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY a
 */

func angryProfessor(k int32, a []int32) string {
	studentCountOnTime := int32(0)
	for _, val := range a {
		if val <= 0 {
			studentCountOnTime++
		}
	}

	if k > studentCountOnTime {
		return "YES"
	}

	return "NO"
}

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
//	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//	checkError(err)
//	t := int32(tTemp)
//
//	for tItr := 0; tItr < int(t); tItr++ {
//		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")
//
//		nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
//		checkError(err)
//		n := int32(nTemp)
//
//		kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
//		checkError(err)
//		k := int32(kTemp)
//
//		aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
//
//		var a []int32
//
//		for i := 0; i < int(n); i++ {
//			aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
//			checkError(err)
//			aItem := int32(aItemTemp)
//			a = append(a, aItem)
//		}
//
//		result := angryProfessor(k, a)
//
//		fmt.Fprintf(writer, "%s\n", result)
//	}
//
//	writer.Flush()
//}
