package main

/*
 * Complete the 'countingValleys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER steps
 *  2. STRING path
 */

func countingValleys(steps int32, path string) int32 {
	runes := []rune(path)
	level := 0
	valleyCount := int32(0)
	for _, r := range runes {
		if r == 'U' {
			level++
		} else if r == 'D' {
			if level == 0 {
				valleyCount++
			}

			level--
		}
	}

	return valleyCount
}

//func main() {
//	result := countingValleys(8, "UDDDUDUU")
//
//	fmt.Printf("%d\n", result)
//}
