package main

/*
 * Complete the getMoneySpent function below.
 */
func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	maxBuy := int32(-1)
	for x := 0; x < len(keyboards); x++ {
		for y := 0; y < len(drives); y++ {
			expense := keyboards[x] + drives[y]
			if expense <= b && maxBuy < expense {
				maxBuy = expense
			}
		}
	}

	return maxBuy

}

//func main() {
//	keyboards := []int32{3, 1}
//	drives := []int32{5, 2, 8}
//	moneySpent := getMoneySpent(keyboards, drives, 10)
//	fmt.Printf("%d\n", moneySpent)
//}
