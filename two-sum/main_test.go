package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCases(t *testing.T) {
	var tests = []struct {
		input          []int
		target         int
		expectedResult []int
	}{
		{[]int{0, 1}, 1, []int{0, 1}},
		{[]int{0, 3}, 3, []int{0, 1}},
		{[]int{0, 1, 3}, 4, []int{1, 2}},
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{-1, -2, -3, -4, -5}, -8, []int{2, 4}},
	}
	for _, tt := range tests {
		actualResult := twoSum(tt.input, tt.target)
		assert.Equal(t, tt.expectedResult, actualResult)
	}
}
