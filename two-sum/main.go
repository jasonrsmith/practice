package main

func twoSum(n []int, target int) []int {
	var i, j int
	for i = 0; i < len(n)-1; i++ {
		for j = i + 1; j < len(n); j++ {
			if n[i]+n[j] == target {
				return []int{i, j}
			}
		}
	}
	panic("Sum not found")
}

func main() {
}
