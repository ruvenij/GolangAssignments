package main

func GetFactorial(n int) int {
	if n == 0 {
		return 1
	}

	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}

	return factorial
}

func GetFactorialUsingRecursion(n int) int {
	if n == 0 {
		return 1
	}

	return n * GetFactorialUsingRecursion(n-1)
}
