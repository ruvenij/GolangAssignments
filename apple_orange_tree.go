package main

import "fmt"

/*
 * Complete the 'countApplesAndOranges' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER s
 *  2. INTEGER t
 *  3. INTEGER a
 *  4. INTEGER b
 *  5. INTEGER_ARRAY apples
 *  6. INTEGER_ARRAY oranges
 */

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	applesFall := 0
	orangesFall := 0
	for _, apple := range apples {
		fallPoint := apple + a
		if fallPoint >= s && fallPoint <= t {
			applesFall++
		}
	}

	for _, orange := range oranges {
		fallPoint := orange + b
		if fallPoint >= s && fallPoint <= t {
			orangesFall++
		}
	}

	fmt.Println(applesFall)
	fmt.Println(orangesFall)
}

//func main() {
//	s := int32(7)
//	t := int32(11)
//	a := int32(5)
//	b := int32(15)
//	apples := []int32{-2, 2, 1}
//	oranges := []int32{5, -6}
//
//	countApplesAndOranges(s, t, a, b, apples, oranges)
//}
