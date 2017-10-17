/*
we have one slice, nums, of n positive integers and another slice maxes, of m positive integers.
For each m in maxes, we want to know the total number of elements in nums which are less than or equal to m.
For example numb =[1,2,3] and maxes=[2,4], there are 2 elementes in nums that are <= m (which is 2)
and 3 elementes in nums that are <= m (which is 4). Store the respective answer in another array
answer=[2,3]
*/
package main

import "fmt"

func solution(num, max []int) []int {
	tmp := make([]int, len(max))
	count := 0
	for i, val := range max {
		count = 0
		for _, val2 := range num {
			if val2 <= val {
				count++
			}

			tmp[i] = count

		}

	}
	return tmp
}

func main() {
	num := []int{2, 10, 5, 4, 8}
	max := []int{3, 1, 7, 8}
	fmt.Println(solution(num, max))
}
