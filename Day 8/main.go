package main

import (
	"fmt"
	"sync"
)

// Write a simple parser that will parse and run Deadfish.

// Deadfish has 4 commands, each 1 character long:

// i increments the value (initially 0)
// d decrements the value
// s squares the value
// o outputs the value into the return array
// Invalid characters should be ignored.

// Parse("iiisdoso") == []int{8, 64}

func starter(wg *sync.WaitGroup) {
	fmt.Println("This is the starter on call")
	wg.Done()
}

func follow() {
	fmt.Println("This is the follower on call")
}

func Parse(data string) []int {
	result := []int{}
	var total int

	for _, char := range data {
		ch := fmt.Sprintf("%c", char)
		if ch == "i" {
			total += 1
		} else if ch == "d" {
			total -= 1
		} else if ch == "s" {
			total *= total
		} else if ch == "o" {
			result = append(result, total)
		}
	}
	return result
}

func main() {
	// fmt.Println(Parse("iiisdoso"))
	var wg sync.WaitGroup
	wg.Add(1)
	go starter(&wg)
	follow()
	wg.Wait()
}
