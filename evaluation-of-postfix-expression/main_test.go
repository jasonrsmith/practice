package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCases(t *testing.T) {
	var tests = []struct {
		input          string
		expectedResult int
	}{
		{"21+", 3},
		{"231*+9-", -4},
		{"123+*8-", -3},
	}
	for _, tt := range tests {
		actualResult := postfix(tt.input)
		assert.Equal(t, tt.expectedResult, actualResult)
	}
}
