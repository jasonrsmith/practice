package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCases(t *testing.T) {
	var tests = []struct {
		l1             *ListNode
		l2             *ListNode
		expectedResult *ListNode
	}{
		{
			&ListNode{1, nil},
			&ListNode{2, nil},
			&ListNode{3, nil},
		},
		{
			&ListNode{1, &ListNode{3, nil}},
			&ListNode{2, nil},
			&ListNode{3, &ListNode{3, nil}},
		},
		{
			&ListNode{6, nil},
			&ListNode{7, nil},
			&ListNode{3, &ListNode{1, nil}},
		},
		{
			&ListNode{5, nil},
			&ListNode{5, nil},
			&ListNode{0, &ListNode{1, nil}},
		},
	}
	for _, tt := range tests {
		actualResult := addTwoNumbers(tt.l1, tt.l2)
		assert.Equal(t, tt.expectedResult, actualResult)
	}
}
