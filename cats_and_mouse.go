package main

// Complete the catAndMouse function below.
func catAndMouse(x int32, y int32, z int32) string {
	catA := x - z
	catB := y - z

	if catA < 0 {
		catA = -1 * catA
	}

	if catB < 0 {
		catB = -1 * catB
	}

	if catA == catB {
		return "Mouse C"
	}

	if catA < catB {
		return "Cat A"
	}

	return "Cat B"

}

//func main() {
//	result := catAndMouse(1, 2, 3)
//	fmt.Printf("%s\n", result)
//
//	result = catAndMouse(1, 3, 2)
//	fmt.Printf("%s\n", result)
//}
