package main

func StrStr(haystack string, needle string) int {
	//return strings.Index(haystack, needle)

	if len(needle) > len(haystack) {
		return -1
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		subStr := haystack[i : i+len(needle)]
		if subStr == needle {
			return i
		}
	}

	return -1
}

type A struct {
	Name string
	Age  int
}

func doSomething() {
	varA := A{
		Name: "Alice",
		Age:  30,
	}

	varA.Age = 18

	sliceA := make([]A, 0)
	var sliceAB [10]A

	sliceA = append(sliceA, varA)
	sliceAB[0] = varA
}
