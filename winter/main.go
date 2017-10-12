//Returns the lenght of winter from a slice of temperatures represented by int
package main

import "fmt"

func solution(T []int) int {

	maxInv := T[0]
	var indexStartSummer int
	maxTemp := T[0]

	for i := range T {
		if maxTemp < T[i] {
			maxTemp = T[i]
		}

		if maxInv >= T[i] {
			indexStartSummer = -1
			maxInv = maxTemp
		} else {
			if indexStartSummer == -1 {
				indexStartSummer = i
			}
		}

	}

	return indexStartSummer
}

func main() {
	//t := []int{5, -2, 3, 6, 8} return 3
	//t := []int{-5, -5, -5, -42, 6, 12} return 4
	//t := []int{-5, -5, -5, -42, 7, 6, 8}
	/*return 4*/
	t := []int{-10, 5, -12, 3, 8, 9}
	fmt.Println("Length of winter ", solution(t))

}
