package main

import (
	"fmt"
)

//Solution .
func Solution(A []int) int {
	m := make(map[int]int)

	b := false

	positive := 1

	for _, val := range A {
		if val > 0 {
			m[val] = val
		}
	}

	for !b {
		if _, ok := m[positive]; !ok {
			b = true
		} else {
			positive++
		}
	}

	return positive
}

func main() {
	A := []int{1, 2, 3}
	fmt.Println(Solution(A))
}
