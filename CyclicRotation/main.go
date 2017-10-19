/*
Each element the slice will be shifted to the right by N indexes
For example, given array
A = [3, 8, 9, 7, 6] and K = 3,
the function should return [9, 7, 6, 3, 8].
*/
package main

import "fmt"

//Solution Rotates A K indexes to the right
func Solution(A []int, K int) []int {

	pos := 0
	tmp := make([]int, len(A))
	a := 0
	for i := range A {
		pos = i + K
		if pos >= len(A) {
			a = (pos % len(A))
			tmp[a] = A[i]
		} else {
			tmp[pos] = A[i]
		}
	}
	return tmp
}

func main() {
	t := []int{3, 8, 9, 7, 6}
	p := 13
	fmt.Printf("Original %x \npositions to the right K=%d \nFinal Result %x\n", t, p, Solution(t, p))
	//fmt.Println(Solution(t, p))
}
