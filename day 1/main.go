package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func sortDictKeys(numbers []int) []int {
	sort.Ints(numbers)
	return numbers
}

// extract keys from maps to an array
func extractKeys(mapData map[int]string) []int {
	var keys []int = []int{}
	for key := range mapData {
		keys = append(keys, key)
	}
	return keys
}

func Order(sentence string) string {

	var sentenceData map[int]string = make(map[int]string)
	var sortedSentence string

	for _, word := range strings.Split(sentence, " ") {
		pattern := regexp.MustCompile("\\d")
		matched, _ := strconv.Atoi(pattern.FindString(word))
		sentenceData[matched] = word
	}

	mapKeys := extractKeys(sentenceData)
	sortedKeys := sortDictKeys(mapKeys)

	for _, num := range sortedKeys {
		word := sentenceData[num]
		sortedSentence += word + " "
	}

	sortedSentence = strings.TrimSpace(sortedSentence)
	fmt.Println(sortedSentence)

	return sortedSentence
}

func main() {
	Order("4of Fo1r pe6ople g3ood th5e the2")
}
