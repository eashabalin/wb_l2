package dev04

import (
	"reflect"
	"testing"
)

func TestAnagram(t *testing.T) {
	words := []string{"пятак", "стОлик", "тяпка", "листоК", "Слиток", "Пятка", "Бумага", "влад", "", "", "листок", "влад", "сТолик"}
	res := Anagrams(words)
	correct := make(map[string][]string)
	correct["пятак"] = []string{"тяпка", "пятка"}
	correct["столик"] = []string{"листок", "слиток"}
	if !reflect.DeepEqual(*res, correct) {
		t.Errorf("got: %v, expected: %v", res, correct)
	}
}
