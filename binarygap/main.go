//BinaryGap - Find longest sequence of zeros in binary representation of an integer.
package main

import (
	"fmt"
	"strconv"
	"strings"
)

//Solution2 gets
func Solution2(N int) int {

	n := int64(N)
	tmp := strconv.FormatInt(n, 2)

	max := 0
	temp := 0
	for _, r := range tmp {
		if r == '0' {
			temp++
		} else {
			if temp > max {
				max = temp

			}
			temp = 0
		}
	}

	return max

}

//Solution converts int into binary
func Solution(N int) int {
	n := int64(N)
	tmp := strconv.FormatInt(n, 2)
	s := strings.Split(tmp, "1")

	max := 0

	for i := range s {
		if len(s[i]) > max {
			max = len(s[i])
		}
	}
	return max
}

func main() {
	//fmt.Println(Solution(561892))
	fmt.Println(Solution2(1376796946))
}
