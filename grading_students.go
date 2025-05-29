package main

/*
 * Complete the 'gradingStudents' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY grades as parameter.
 */

func gradingStudents(grades []int32) []int32 {
	result := make([]int32, 0)
	for _, oldGrade := range grades {
		newGrade := oldGrade
		if oldGrade >= 38 {
			val := oldGrade / 5
			diff := ((val + 1) * 5) - oldGrade

			if diff < 3 {
				newGrade = (val + 1) * 5
			}
		}

		result = append(result, newGrade)
	}

	return result
}

//func main() {
//	grades := []int32{4, 73, 67, 38, 33}
//	result := gradingStudents(grades)
//
//	for i, resultItem := range result {
//		fmt.Printf("%d", resultItem)
//
//		if i != len(result)-1 {
//			fmt.Printf("\n")
//		}
//	}
//
//	fmt.Printf("\n")
//}
