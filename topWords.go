package main

import (
	"fmt"
	"sort"
	"strings"
)

type WordCount struct {
	word  string
	count int
}

func topWords(s string, n int) []string {
	str := []string{s}
	words := strings.Split(strings.Join(str, ""), " ")

	// count same words
	m := make(map[string]int)
	for _, word := range words {
		if _, ok := m[word]; ok {
			m[word]++
		} else {
			m[word] = 1
		}
	}

	// create and fill slice
	wordCounts := make([]WordCount, 0, len(m))
	for key, val := range m {
		wordCounts = append(wordCounts, WordCount{word: key, count: val})
	}

	// sort wordCount slice by decreasing count number
	sort.Slice(wordCounts, func(i, j int) bool {
		return wordCounts[i].count > wordCounts[j].count
	})

	myslice := make([]string, len(m))

	for i := 0; i < len(wordCounts) && i < n; i++ {
		myslice = append(myslice, wordCounts[i].word)
	}
	return myslice
}

func main() {
	fmt.Println(topWords("this is my is is my wow so is is my my so hey no yes", 10))
}
