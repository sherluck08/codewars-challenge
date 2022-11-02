package main

import (
	"fmt"
	"sort"
)

func TwoOldestAges(ages []int) [2]int {
	sort.Ints(ages)
	sortedAges := [2]int{ages[len(ages)-2], ages[len(ages)-1]}
	return sortedAges
}

func main() {
	numbers := []int{1, 2, 10, 8}
	lastTwo := TwoOldestAges(numbers)
	fmt.Println(lastTwo)
}
