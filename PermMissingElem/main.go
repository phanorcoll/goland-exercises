/*
A zero-indexed array A consisting of N different integers is given. The array contains integers in the range [1..(N + 1)], which means that exactly one element is missing.

Your goal is to find that missing element.
For example, given array A such that:

  A[0] = 2
  A[1] = 3
  A[2] = 1
  A[3] = 5
the function should return 4, as it is the missing element.

Assume that:
-N is an integer within the range [0..100,000];
-the elements of A are all distinct;
-each element of array A is an integer within the range [1..(N + 1)].

Complexity:
-expected worst-case time complexity is O(N);
-expected worst-case space complexity is O(1), beyond input storage (not counting the storage required for input arguments).
-Elements of input arrays can be modified.


*/
package main

import (
	"fmt"
)

//Solution .
func Solution(A []int) int {

	tsum := 0
	min := 1
	max := 2
	var result float64
	for _, val := range A {

		if val > max {
			max = val
		}

		tsum = tsum + val
	}

	if max == len(A) {
		max++
	}

	//we get the sum of the arithmetic sequence (n/2)(min+max)
	sum := (float64((len(A) + 1)) / float64(2)) * (float64(min) + float64(max))

	result = sum - float64(tsum)

	if len(A) == 0 {
		return 1
	}
	return int(result)

}

func main() {
	A := []int{2, 3, 1, 5}
	//A := []int{9, 5, 8, 6}
	//A := []int{5}
	//A := []int{1, 2, 3, 4}
	//A := []int{1, 2}
	fmt.Println(Solution(A))
}
