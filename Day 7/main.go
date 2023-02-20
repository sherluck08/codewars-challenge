package main

import "fmt"

// ## Day 7 challenge

// Complete the method which returns the number which is most frequent in the given input array. If there is a tie for most frequent number, return the largest number among them.

// Note: no empty arrays will be given.

// ### Examples

// `[12, 10, 8, 12, 7, 6, 4, 10, 12] --> 12`
// `[12, 10, 8, 12, 7, 6, 4, 10, 12, 10] --> 12`
// `[12, 10, 8, 8, 3, 3, 3, 3, 2, 4, 10, 12, 10] --> 3`

func HighestRank(nums []int) int {

	numbers := make(map[int]int)
	for _, num := range nums {
		numbers[num] = numbers[num] + 1
	}
	var maxValue int
	var highestCount int
	var finalHighest int
	for key, value := range numbers {
		if value > maxValue {
			maxValue = value
			highestCount = key
		}
		if numbers[highestCount] == value {
			if highestCount != key {
				if highestCount > key {
					finalHighest = highestCount
				} else if highestCount < key {
					finalHighest = key
				}

			} else {
				finalHighest = highestCount
			}
		}
	}
	return finalHighest
}

func main() {
	nums := []int{12, 10, 8, 12, 7, 6, 4, 10, 10, 12}
	fmt.Println(HighestRank(nums))
}
