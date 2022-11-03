package main

import (
	"fmt"
	"sort"
)

func GetSum(a, b int) int {
	
	if a == b {
		return a
	}
	nums := []int{a, b}
	sort.Ints(nums)

	a, b = nums[0], nums[1]

	fmt.Println(nums)
	var sum int
	for ;a <= b; a++{
		sum += a
	}
	return sum
}

func main(){
	fmt.Println(GetSum(0, 1))
}