package main

import (
	"fmt"
)

/*
 * Complete the 'dayOfProgrammer' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts INTEGER year as parameter.
 */

func dayOfProgrammer(year int32) string {
	isLeapYear := false
	if year < 1918 {
		// julian calendar
		if year%4 == 0 {
			isLeapYear = true
		}
	} else {
		// gregorian calendar
		if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
			isLeapYear = true
		}
	}

	febDays := 28
	if isLeapYear {
		febDays = 29
	}

	remainingDays := 41 - febDays
	if year == 1918 {
		remainingDays += 13
	}

	return fmt.Sprintf("%d.09.%d", remainingDays, year)

}

//func main() {
//	result := dayOfProgrammer(2017)
//	fmt.Printf("%s\n", result)
//}
