package main

import "sort"

func GroupAnagram(words []string) [][]string {
	groupWithResult := map[string][]string{}

	for _, word := range words {
		wordRunes := []rune(word)
		sort.Slice(wordRunes, func(i int, j int) bool {
			return wordRunes[i] < wordRunes[j]
		})

		sortedWord := string(wordRunes)

		_, ok := groupWithResult[sortedWord]
		if !ok {
			groupWithResult[sortedWord] = []string{}
		}

		groupWithResult[sortedWord] = append(groupWithResult[sortedWord], word)
	}

	results := [][]string{}
	for _, result := range groupWithResult {
		results = append(results, result)
	}

	return results
}
