package main

func RemoveDuplicates(nums []int) int {
	existingElements := make(map[int]interface{})
	finalElements := make([]int, 0)
	currentIndex := 0

	for _, val := range nums {
		if _, ok := existingElements[val]; !ok {
			finalElements = append(finalElements, val)
			nums[currentIndex] = val
			currentIndex++
			existingElements[val] = struct{}{}
		}
	}

	return len(finalElements)
}
