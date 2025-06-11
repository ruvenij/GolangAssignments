package main

import (
	"strings"
)

func convert(s string, numRows int) string {
	if numRows <= 1 || numRows >= len(s) {
		return s
	}

	rows := make([]strings.Builder, numRows)
	currentRow := 0
	isGoingDown := true
	for i := 0; i < len(s); i++ {
		rows[currentRow].WriteByte(s[i])

		if (currentRow == numRows-1 && isGoingDown) || (currentRow == 0 && !isGoingDown) {
			isGoingDown = !isGoingDown
		}

		if isGoingDown {
			currentRow++
		} else {
			currentRow--
		}
	}

	finalString := strings.Builder{}
	finalString.Grow(len(s))
	for rowIndex := 0; rowIndex < numRows; rowIndex++ {
		finalString.WriteString(rows[rowIndex].String())
	}

	return finalString.String()
}

//func main() {
//	fmt.Println(convert("PAYPALISHIRING", 3))
//}
