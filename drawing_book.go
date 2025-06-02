package main

/*
 * Complete the 'pageCount' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER p
 */

func pageCount(n int32, p int32) int32 {
	isEndPageOdd := true
	if n%2 == 0 {
		isEndPageOdd = false
	}

	isTargetOdd := true
	if p%2 == 0 {
		isTargetOdd = false
	}

	turnsFromBeginning := int32(0)
	turnsFromEnd := int32(0)

	if isTargetOdd {
		turnsFromBeginning = (p - 1) / 2
	} else {
		turnsFromBeginning = p / 2
	}

	if isEndPageOdd {
		if isTargetOdd {
			turnsFromEnd = (n - p) / 2
		} else {
			turnsFromEnd = (n - p - 1) / 2
		}
	} else {
		if isTargetOdd {
			turnsFromEnd = (n - p + 1) / 2
		} else {
			turnsFromEnd = (n - p) / 2
		}
	}

	if turnsFromBeginning < turnsFromEnd {
		return turnsFromBeginning
	}

	return turnsFromEnd
}

//func main() {
//	result := pageCount(6, 2)
//	fmt.Printf("%d\n", result)
//}
