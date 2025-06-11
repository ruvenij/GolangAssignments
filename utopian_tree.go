package main

/*
 * Complete the 'utopianTree' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER n as parameter.
 */

func utopianTree(n int32) int32 {
	height := int32(1)
	for i := int32(1); i <= n; i++ {
		if i%2 == 0 {
			height++
		} else {
			height *= 2
		}
	}

	return height
}

//func main() {
//	n := int32(0)
//	result := utopianTree(n)
//	fmt.Printf("%d\n", result)
//
//	n = int32(1)
//	result = utopianTree(n)
//	fmt.Printf("%d\n", result)
//
//	n = int32(4)
//	result = utopianTree(n)
//	fmt.Printf("%d\n", result)
//}
