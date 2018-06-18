package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClosestPalindrome(t *testing.T) {
	var cases = map[string]struct{ input, expected string }{
		"singleDigit":  {"1", "0"},
		"singleDigit0": {"0", "1"},
		"twoDigits":    {"11", "9"},
		"twoDigits22":  {"19", "22"},
		"1009":         {"1009", "1001"},
		"1001":         {"1001", "999"},
		"99":           {"99", "101"},
		"long":         {"23094587234590", "23094588549032"},
		"long2":        {"807045053224792883", "807045053350540708"},
	}

	for k, tc := range cases {
		actual := ClosestPalindrome(tc.input)
		assert.Equal(t, tc.expected, actual, k)
	}
}

func TestDecNumArr(t *testing.T) {
	var cases = map[string]struct{ input, expected []int }{
		"singleEl": {[]int{1}, []int{0}},
		"10":       {[]int{1, 0}, []int{9}},
		"11":       {[]int{1, 1}, []int{1, 0}},
	}

	for k, tc := range cases {
		actual := decNumArr(tc.input)
		assert.Equal(t, tc.expected, actual, k)
	}
}

func TestIncNumArr(t *testing.T) {
	var cases = map[string]struct{ input, expected []int }{
		"singleEl": {[]int{0}, []int{1}},
		"9":        {[]int{9}, []int{1, 0}},
		"21":       {[]int{2, 1}, []int{2, 2}},
	}

	for k, tc := range cases {
		actual := incNumArr(tc.input)
		assert.Equal(t, tc.expected, actual, k)
	}
}

func TestIsPalindrome(t *testing.T) {
	var cases = map[string]struct {
		input    []int
		expected bool
	}{
		"singleEl": {[]int{0}, true},
		"9":        {[]int{9}, true},
		"11":       {[]int{1, 1}, true},
		"21":       {[]int{2, 1}, false},
		"990":      {[]int{9, 9, 0}, false},
		"999":      {[]int{9, 9, 9}, true},
		"989":      {[]int{9, 8, 9}, true},
		"1000":     {[]int{1, 0, 0, 0}, false},
		"1001":     {[]int{1, 0, 0, 1}, true},
	}

	for k, tc := range cases {
		actual := isPalindrome(tc.input)
		assert.Equal(t, tc.expected, actual, k)
	}
}
