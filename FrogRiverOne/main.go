/*
A small frog wants to get to the other side of a river. The frog is initially located on one bank of the river (position 0) and wants to get to the opposite bank (position X+1). Leaves fall from a tree onto the surface of the river.

You are given a zero-indexed array A consisting of N integers representing the falling leaves. A[K] represents the position where one leaf falls at time K, measured in seconds.

The goal is to find the earliest time when the frog can jump to the other side of the river. The frog can cross only when leaves appear at every position across the river from 1 to X (that is, we want to find the earliest moment when all the positions from 1 to X are covered by leaves). You may assume that the speed of the current in the river is negligibly small, i.e. the leaves do not change their positions once they fall in the river.

For example, you are given integer X = 5 and array A such that:

  A[0] = 1
  A[1] = 3
  A[2] = 1
  A[3] = 4
  A[4] = 2
  A[5] = 3
  A[6] = 5
  A[7] = 4
In second 6, a leaf falls into position 5. This is the earliest time when leaves appear in every position across the river.

Write a function:

func Solution(X int, A []int) int

that, given a non-empty zero-indexed array A consisting of N integers and integer X, returns the earliest time when the frog can jump to the other side of the river.

If the frog is never able to jump to the other side of the river, the function should return −1.

For example, given X = 5 and array A such that:

  A[0] = 1
  A[1] = 3
  A[2] = 1
  A[3] = 4
  A[4] = 2
  A[5] = 3
  A[6] = 5
  A[7] = 4
the function should return 6, as explained above.

Assume that:
-N and X are integers within the range [1..100,000];
-each element of array A is an integer within the range [1..X].

Complexity:
-expected worst-case time complexity is O(N);
-expected worst-case space complexity is O(X), beyond input storage (not counting the storage required for input arguments).

Elements of input arrays can be modified.
*/
package main

import "fmt"

//Solution .
func Solution(X int, A []int) int {
	m := make(map[int]int)
	var answer int
	min := 1

	//create a map from the slice
	for i, val := range A {
		_, exist := m[val]
		if !exist {
			m[val] = i
		}

	}

	max := m[1]

	for i := min; i <= X; i++ {

		if m[i] > max {
			max = m[i]
		}

		_, exist := m[i]

		if !exist {
			answer = -1
			return answer
		}
	}

	answer = max
	//end
	fmt.Println(m)

	return answer
}

func main() {
	//A := []int{1, 3, 1, 4, 2, 3, 5, 4}
	//X := 5
	A := []int{2, 2, 2, 2, 2}
	X := 2
	// A := []int{1}
	// X := 1
	// A := []int{1, 3, 1, 3, 2, 1, 3}
	// X := 3
	fmt.Println(Solution(X, A))
}
