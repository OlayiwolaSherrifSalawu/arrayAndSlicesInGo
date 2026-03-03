package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	text := "Go is an is is open source programming language that makes it easy to build simple reliable and efficient software"
	words := strings.Fields(strings.ToLower(text))

	wordCounts := make(map[string]int)

	for _, word := range words {
		wordCounts[word]++
	}
	sortArray := make([]string, 0, len(wordCounts))
	for word := range wordCounts {
		sortArray = append(sortArray, word)
	}
	sort.Strings(sortArray)
	for _, val := range sortArray {
		fmt.Printf("%s:%d \n", val, wordCounts[val])
	}
}
