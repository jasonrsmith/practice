package main

import (
	"bytes"
	"strconv"
)

func ClosestPalindrome(numStr string) string {
	if numStr == "0" {
		return "1"
	}

	numsInc := strToNumArr(numStr)
	numsDec := make([]int, len(numsInc))
	copy(numsDec, numsInc)

	for {
		numsDec = decNumArr(numsDec)
		if isPalindrome(numsDec) {
			return numArrToStr(numsDec)
		}
		numsInc = incNumArr(numsInc)
		if isPalindrome(numsInc) {
			return numArrToStr(numsInc)
		}
	}
}

func isPalindrome(numArr []int) bool {
	halfLen := len(numArr) / 2
	j := len(numArr) - 1
	for i := 0; i < halfLen; i++ {
		if numArr[i] != numArr[j] {
			return false
		}
		j--
	}
	return true
}

func decNumArr(numArr []int) []int {
	numArr[len(numArr)-1]--
	for i := len(numArr) - 1; i > 0; i-- {
		if numArr[i] == -1 {
			numArr[i] = 9
			numArr[i-1]--
		}
	}
	if numArr[0] == 0 && len(numArr) > 1 {
		numArr = numArr[1:]
	}

	return numArr

}

func incNumArr(numArr []int) []int {
	numArr[len(numArr)-1]++
	for i := len(numArr) - 1; i > 0; i-- {
		if numArr[i] == 10 {
			numArr[i] = 0
			numArr[i-1]++
		}
	}
	if numArr[0] == 10 {
		numArr[0] = 0
		numArr = append([]int{1}, numArr...)
	}

	return numArr
}

func strToNumArr(numStr string) []int {
	nums := make([]int, 0)
	for _, c := range numStr {
		number, _ := strconv.Atoi(string(c))
		nums = append(nums, number)
	}
	return nums
}

func numArrToStr(nums []int) string {
	if len(nums) == 1 {
		return strconv.Itoa(nums[0])
	}

	var buf bytes.Buffer
	leadingZero := true
	for _, n := range nums {
		if !(n == 0 && leadingZero) {
			buf.WriteString(strconv.Itoa(n))
			leadingZero = false
		}
	}
	return buf.String()
}

func main() {
}

/*
Given a number N. our task is to find the closest Palindrome number whose absolute difference with given number is minimum.

Input:
The first line of the input contains integer T denoting the number of test cases. Each test case contains a  number N.

Output:
For each test case, the print the closest palindrome number.
Note:  If the difference of two closest palindromes numbers is equal then we print smaller number as output.

Constraints:
1<=T<=1000
1<=n<=10^14

Input :
2
9
489

output:
8
484

Explanation :

Test Case 1: closest palindrome number is 8.
*/
