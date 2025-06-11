package main

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for i, val := range nums {
		remainder := target - val
		if secondIndex, ok := indexMap[remainder]; ok {
			return []int{secondIndex, i}
		}

		indexMap[val] = i
	}

	return []int{}
}

//func main() {
//	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
//}
