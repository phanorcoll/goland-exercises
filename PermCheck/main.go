/*
A non-empty zero-indexed array A consisting of N integers is given.

A permutation is a sequence containing each element from 1 to N once, and only once.

For example, array A such that:

    A[0] = 4
    A[1] = 1
    A[2] = 3
    A[3] = 2
is a permutation, but array A such that:

    A[0] = 4
    A[1] = 1
    A[2] = 3
is not a permutation, because value 2 is missing.

The goal is to check whether array A is a permutation.

Write a function:

func Solution(A []int) int

that, given a zero-indexed array A, returns 1 if array A is a permutation and 0 if it is not.

For example, given array A such that:

    A[0] = 4
    A[1] = 1
    A[2] = 3
    A[3] = 2
the function should return 1.

Given array A such that:

    A[0] = 4
    A[1] = 1
    A[2] = 3
the function should return 0.

Assume that:
-N is an integer within the range [1..100,000];
-each element of array A is an integer within the range [1..1,000,000,000].

Complexity:
-expected worst-case time complexity is O(N);
-expected worst-case space complexity is O(N), beyond input storage (not counting the storage required for input arguments).

Elements of input arrays can be modified.
*/
package main

import "fmt"

//Solution .
func Solution(A []int) int {
	m := make(map[int]int)
	var answer int
	min := A[0]
	max := A[0]
	//counts the recurrences of a N integer in the slice
	for _, value := range A {

		if value < min {
			min = value
		}

		if value > max {
			max = value
		}

		_, exist := m[value]
		if exist {
			m[value]++
		} else {
			m[value] = 1
		}

	}

	for _, val := range m {
		if val > 1 {
			answer = 0
			return answer
		}

	}

	//verify if the number is more than once, returns 0 if it finds a duplicate N integer
	for i := min; i <= max; i++ {
		if m[i] == 0 {
			answer = 0
			return answer
		}

	}

	if len(A) <= 2 {
		answer = 1
		return answer
	}

	return 1
}

func main() {
	//A := []int{9, 8, 10, 12, 1000000}
	A := []int{1, 2}
	fmt.Println(Solution(A))
}
