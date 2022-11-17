package main

import (
	"fmt"
)

// You are given an array (which will have a length of at least 3, but could be very large) containing integers. The array is either entirely comprised of odd integers or entirely comprised of even integers except for a single integer N. Write a method that takes the array as an argument and returns this "outlier" N.

// Examples
// [2, 4, 0, 100, 4, 11, 2602, 36]
// Should return: 11 (the only odd number)

// [160, 3, 1719, 19, 11, 13, -21]
// Should return: 160 (the only even number)

func FindOutlier(integers []int) int {
	oddNum := []int{}
	evenNum := []int{}

	for _, num := range integers {
		if num%2 == 0 {
			evenNum = append(evenNum, num)
			if len(evenNum) > 1 && len(oddNum) > 0 {
				return oddNum[0]
			}
		} else {
			oddNum = append(oddNum, num)
		}
	}
	if len(evenNum) > 1 && len(oddNum) > 0 {
		return oddNum[0]
	} else if len(oddNum) > 1 && len(evenNum) > 0 {
		return evenNum[0]
	}

	return -1
}

func main() {
	nums := []int{206847684, 1056521, 7, 17, 1901, 21104421, 7, 1, 35521, 1, 7781}
	fmt.Println(FindOutlier(nums))
}
