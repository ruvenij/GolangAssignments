package main

/*
 * Complete the 'designerPdfViewer' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY h
 *  2. STRING word
 */

func designerPdfViewer(h []int32, word string) int32 {
	runes := []rune(word)
	maxHeight := int32(0)
	for _, r := range runes {
		runeIndex := r - 'a'
		if maxHeight < h[runeIndex] {
			maxHeight = h[runeIndex]
		}
	}

	return maxHeight * int32(len(runes))
}

//func main() {
//	h := []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}
//	word := "abc"
//	result := designerPdfViewer(h, word)
//	fmt.Printf("%d\n", result)
//}
