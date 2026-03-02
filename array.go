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

	for word, count := range wordCounts {
		fmt.Printf("%s :  %d \n", word, count)
	}
	sortArray:=
}
