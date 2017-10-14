/*find the unpair element in the slice
The array contains an odd number of elements, and each element of the array can be paired with another element that has the same value,
except for one element that is left unpaired.

Assume that:
  -N is an odd integer within the range [1..1,000,000];
  -each element of array A is an integer within the range [1..1,000,000,000];
  -all but one of the values in A occur an even number of times.


*/
package main

import "fmt"

/*
Solution takes a slice of int and returns the unpaired int
 - first we check that item exists in m
 - if it exists, increase counter by one
 - if it doesnt exist start count at one
 - finally we search the map for the value equals to one and return it
*/
func Solution(A []int) int {

	m := make(map[int]int)
	unpaired := 0

	for _, value := range A {
		_, exist := m[value]
		if exist {
			m[value]++
		} else {
			m[value] = 1
		}
	}
	for i, val := range m {
		if val == 1 {
			unpaired = i
		}
	}
	return unpaired
}

func main() {
	t := []int{9, 3, 9, 3, 9, 7, 9}
	fmt.Println(Solution(t))
}
