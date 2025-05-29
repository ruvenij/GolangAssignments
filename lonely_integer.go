package main

/*
 * Complete the 'lonelyinteger' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func lonelyinteger(a []int32) int32 {
	elementMap := make(map[int32]int32)
	for _, v := range a {
		elementMap[v]++
	}

	for k, v := range elementMap {
		if v == 1 {
			return k
		}
	}

	return 0
}

//func main() {
//	a := []int32{1, 2, 3, 4, 3, 2, 1}
//	result := lonelyinteger(a)
//	fmt.Printf("%d\n", result)
//}
