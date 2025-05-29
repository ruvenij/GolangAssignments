package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	timeWithoutHour := s[2:8]
	hour, _ := strconv.Atoi(s[0:2])
	if strings.HasSuffix(s, "PM") && hour < 12 {
		militaryHour := 12 + hour
		formattedHour := fmt.Sprintf("%02d", militaryHour)
		return formattedHour + timeWithoutHour
	} else if strings.HasSuffix(s, "AM") && hour == 12 {
		return "00" + timeWithoutHour
	}

	return s[0:8]
}

//func main() {
//	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
//
//	stdout, _ := os.Create(os.Getenv("OUTPUT_PATH"))
//	//checkError(err)
//
//	defer stdout.Close()
//
//	writer := bufio.NewWriterSize(stdout, 16*1024*1024)
//
//	s := readLine(reader)
//
//	result := timeConversion(s)
//
//	fmt.Fprintf(writer, "%s\n", result)
//
//	writer.Flush()
//}
