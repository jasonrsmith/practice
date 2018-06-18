package main

import "strconv"

func postfix(exp string) int {
	expb := []byte(exp)
	result := make([]int, 0)
	for _, v := range expb {
		if v == "+" {
		}
		r, err := strconv.Atoi(string(v))
	}
	if len(result) != 1 {
		panic("invalid postfix expression")
	}
	return result[0]
}

func main() {
}
