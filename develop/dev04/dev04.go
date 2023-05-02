package dev04

import (
	"reflect"
	"strings"
)

func Anagrams(words []string) *map[string][]string {
	m := make(map[string][]string, len(words))
	added := make(map[string]bool, len(words))

	for i, w := range words {
		if w == "" {
			continue
		}
		w = strings.ToLower(w)
		letterMap1 := make(map[rune]int)
		for _, l := range w {
			letterMap1[l] += 1
		}
		for j := i + 1; j < len(words); j++ {
			w2 := strings.ToLower(words[j])
			if added[w2] == true || w == w2 {
				continue
			}
			letterMap2 := make(map[rune]int)
			for _, l := range w2 {
				letterMap2[l] += 1
			}
			if reflect.DeepEqual(letterMap1, letterMap2) {
				m[w] = append(m[w], w2)
				added[w] = true
				added[w2] = true
			}
		}
	}

	return &m
}
