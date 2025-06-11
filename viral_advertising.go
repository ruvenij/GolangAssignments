package main

/*
 * Complete the 'viralAdvertising' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER n as parameter.
 */

func viralAdvertising(n int32) int32 {
	peopleForNextDay := int32(5)
	totalLikes := int32(0)

	for i := int32(0); i < n; i++ {
		likesForCurrentDay := peopleForNextDay / 2
		totalLikes += likesForCurrentDay
		peopleForNextDay = likesForCurrentDay * 3
	}

	return totalLikes
}

//func main() {
//	n := int32(5)
//	result := viralAdvertising(n)
//	fmt.Printf("%d\n", result)
//}
