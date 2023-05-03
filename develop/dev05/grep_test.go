package dev05

import (
	"testing"
)

func TestGrep_Find(t *testing.T) {
	in := []string{"      41.14KK", "      14545M", "254u", "   .1MMMk", "fafaf"}
	expect := []string{"      41.14KK", "      14545M", "   .1MMMk"}
	options := GrepOptions{Pattern: "     14545m", LineNum: true, Fixed: false, IgnoreCase: true}
	grep, err := NewGrep(options)
	if err != nil {
		t.Errorf("error creating grep object")
	}
	res := grep.Find(in)

	for i, v := range res {
		if v != expect[i] {
			t.Fail()
		}
	}
	t.Logf("got: %v, expected: %v", res, expect)
}
