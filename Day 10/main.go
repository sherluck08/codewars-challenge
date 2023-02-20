package main

import (
	"fmt"
	"strings"
)

// In this Kata, you will be given an array of strings and your task is to remove all consecutive duplicate letters from each string in the array.

// For example:

// dup(["abracadabra","allottee","assessee"]) = ["abracadabra","alote","asese"].

// dup(["kelless","keenness"]) = ["keles","kenes"].

// Strings will be lowercase only, no spaces. See test cases for more examples.

func Dup(arr []string) []string {
	var withoutDuplicates []string
	for _, word := range arr {
		var words []string
		for index, alphabet := range word {
			if index != 0 {
				if word[index-1] != word[index] {
					words = append(words, string(alphabet))
				}
			} else {
				words = append(words, string(alphabet))
			}
		}
		withoutDuplicates = append(withoutDuplicates, strings.Join(words, ""))
	}
	return withoutDuplicates
}

func main() {
	words := []string{"abracadabra", "allottee", "assessee"}
	duplicate := Dup(words)
	fmt.Println(duplicate)
}
