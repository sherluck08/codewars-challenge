package main

import "fmt"

func WordsToMarks(s string) int {
	alphabetAscii := []rune(s)
	var sum int
	for _, ascii_num := range alphabetAscii {
		sum += int(ascii_num) - 96
	}
	return sum
}

func main() {
	fmt.Println(WordsToMarks("attitude"))
}
