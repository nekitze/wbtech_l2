package dev04

import (
	"slices"
	"sort"
	"strings"
)

func FindAnagrams(dictionary []string) map[string][]string {
	sets := make(map[string][]string)

	for _, word := range dictionary {
		word = strings.ToLower(word)

		setFound := false
		for key, _ := range sets {
			if isAnagram(word, key) {
				sets[key] = append(sets[key], word)
				setFound = true
				break
			}
		}

		if !setFound {
			sets[word] = append(sets[word], word)
		}
	}

	for word, _ := range sets {
		if len(sets[word]) == 1 {
			delete(sets, word)
		}
		sort.Strings(sets[word])
	}

	return sets
}

func isAnagram(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}

	sourceArray := []rune(a)
	sort.Slice(sourceArray, func(i, j int) bool {
		return sourceArray[i] < sourceArray[j]
	})
	targetArray := []rune(b)
	sort.Slice(targetArray, func(i, j int) bool {
		return targetArray[i] < targetArray[j]
	})

	return slices.Equal(sourceArray, targetArray)
}
