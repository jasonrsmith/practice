package main

import "fmt"

func nbonacci(n, m int) []int {
	r := []int{0, 0, 1}
	for i := 0; i < m-3; i++ {
		var next int
		for j := len(r) - 1; j > 0 && j > len(r)-n-1; j-- {
			next += r[j]
		}
		r = append(r, next)
	}
	return r[:m]
}

func main() {
	fmt.Println("vim-go")
}

/**
N-bonacci Numbers
You are given two Integers N and M, and print all the terms of the series upto M-terms of the N-bonacci Numbers. For example, when N = 2, the sequence becomes Fibonacci, when n = 3, sequence becomes Tribonacci.

In general, in N-bonacci sequence, we use sum of preceding N numbers from the next term. For example, a 3-bonacci sequence is the following:
0, 0, 1, 1, 2, 4, 7, 13, 24, 44, 81

The Fibonacci sequence is a set of numbers that starts with one or zero, followed by a one, and proceeds based on the rule that each number is equal to the sum of preceding two numbers 0, 1, 1, 2, 3, 5, 8…..

Examples:

Input : N = 3, M = 8
Output : 0, 0, 1, 1, 2, 4, 7, 13
We need to print first M terms.
First three terms are 0, 0 and 1.
Fourth term is 0 + 0 + 1 = 1
Fifth term is 0 + 1 + 1 = 2
Sixth terms is 1 + 1 + 2 = 4
Seventh term is 7 (1 + 2 + 4) and eighth
term is 13 (7 + 4 + 2).

Input : N = 4, M = 10
Output : 0, 0, 0, 1, 1, 2, 4, 8, 15, 29
*/
