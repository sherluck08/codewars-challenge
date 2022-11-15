package main

import (
	"fmt"
	"strings"
)

// ## Day 6 challenge

// Your task is to remove all duplicate words from a string, leaving only single (first) words entries.

// Example:

// Input:

// 'alpha beta beta gamma gamma gamma delta alpha beta beta gamma gamma gamma delta'

// Output:

// 'alpha beta gamma delta'

func RemoveDuplicateWords(str string) string {
	s := strings.Split(str, " ")
	wordsMap := make(map[string]int)
	var nonDuplicate string

	for _, word := range s {
		if wordsMap[word] > 0 {
			continue
		}
		wordsMap[word] = wordsMap[word] + 1
		nonDuplicate += word + " "

	}
	return strings.TrimSpace(nonDuplicate)
}

func main() {
	fmt.Println(RemoveDuplicateWords("my cat is my cat fat"))
}
