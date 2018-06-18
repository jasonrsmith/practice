package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCases(t *testing.T) {
	var tests = []struct {
		n              int
		m              int
		expectedResult []int
	}{
		{1, 1, []int{0}},
		{2, 3, []int{0, 0, 1}},
		{2, 4, []int{0, 0, 1, 1}},
		{2, 5, []int{0, 0, 1, 1, 2}},
		{2, 6, []int{0, 0, 1, 1, 2, 3}},
		{3, 8, []int{0, 0, 1, 1, 2, 4, 7, 13}},
	}
	for _, tt := range tests {
		actualResult := nbonacci(tt.n, tt.m)
		assert.Equal(t, tt.expectedResult, actualResult)
	}
}
