package main

func maxArea(height []int) int {
	maxAreaRes := 0
	start, end := 0, len(height)-1

	for start < end {
		width := end - start
		currentHeight := height[start]
		if height[end] < currentHeight {
			currentHeight = height[end]
		}

		if maxAreaRes < width*currentHeight {
			maxAreaRes = width * currentHeight
		}

		if height[end] < height[start] {
			end--
		} else {
			start++
		}
	}

	return maxAreaRes
}

//func main() {
//	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
//}
