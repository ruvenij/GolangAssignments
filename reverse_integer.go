package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func reverse(x int) int {
	isNegative := false
	if x < 0 {
		isNegative = true
		x = -x
	}

	s := fmt.Sprintf("%d", x)
	result := strings.Builder{}
	for i := len(s) - 1; i >= 0; i-- {
		result.WriteByte(s[i])
	}

	res, err := strconv.Atoi(result.String())
	if err != nil {
		return 0
	}

	if isNegative {
		res = -res
	}

	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}

	return res
}

//func main() {
//	fmt.Println(reverse(123))
//}
