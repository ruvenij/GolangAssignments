package main

import (
	"errors"
	"slices"
)

func FindMinMaxWithSort(input []int) (int, int, error) {
	if len(input) == 0 {
		return -1, -1, errors.New("Empty slice ")
	}

	slices.Sort(input)
	return input[0], input[len(input)-1], nil
}

func FindMinMaxWithRange(input []int) (int, int, error) {
	if len(input) == 0 {
		return -1, -1, errors.New("Empty slice ")
	}

	minSlice := input[0]
	maxSlice := input[0]
	for _, element := range input {
		if element < minSlice {
			minSlice = element
		}

		if element > maxSlice {
			maxSlice = element
		}
	}

	return minSlice, maxSlice, nil
}
