package main

func RemoveElement(nums []int, val int) int {
	currentEndIndex := len(nums) - 1
	finalElements := make([]int, 0)

	for i, element := range nums {
		if element == val {
			currentEndIndex = getLastValidElement(nums, currentEndIndex, val)

			if currentEndIndex < 0 {
				break
			}

			nums[i] = nums[currentEndIndex]
			currentEndIndex--
		} else {
			finalElements = append(finalElements, element)
		}
	}

	return len(finalElements)
}

func getLastValidElement(nums []int, currentLastValidIndex int, val int) int {
	for index := currentLastValidIndex; index >= 0; index-- {
		if nums[index] != val {
			return index
		}
	}

	return -1
}
