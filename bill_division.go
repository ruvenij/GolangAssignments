package main

import (
	"fmt"
)

/*
 * Complete the 'bonAppetit' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY bill
 *  2. INTEGER k
 *  3. INTEGER b
 */

func bonAppetit(bill []int32, k int32, b int32) {
	totalShared := int32(0)
	for i := int32(0); i < int32(len(bill)); i++ {
		if i != k {
			totalShared += bill[i]
		}
	}

	annaShouldPay := totalShared / 2
	if annaShouldPay == b {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(b - annaShouldPay)
	}
}

//func main() {
//	k := int32(1)
//	bill := []int32{3, 10, 2, 9}
//	b := int32(12)
//
//	bonAppetit(bill, k, b)
//}
