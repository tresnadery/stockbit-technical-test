package test4

import (
	"sort"
	"strings"
)

func GroupWordsByAnagram(words []string) map[string][]string {
	if len(words) == 0 {
		return nil
	}

	groupByAnagram := make(map[string][]string)
	for _, word := range words {
		wordToSlice := strings.Split(word, "")
		sort.Strings(wordToSlice)
		groupByAnagram[strings.Join(wordToSlice, "")] = append(groupByAnagram[strings.Join(wordToSlice, "")], word)
	}
	return groupByAnagram
}
