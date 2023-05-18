package dev05

import (
	"testing"
)

func TestGrep_Find(t *testing.T) {
	in := []string{"      41.14KK", "fafaf", "254u", "OOOOO", "   .MMMk1", "fafaf"}
	expect := []string{"      41.14KK", "254u"}
	options := GrepOptions{Pattern: "[1-4].+", LineNum: false, Fixed: false, IgnoreCase: true, Invert: false, Context: 0}
	grep, err := NewGrep(options)
	if err != nil {
		t.Errorf("error creating grep object")
	}
	res := grep.Find(in)

	if len(res) != len(expect) {
		t.Fail()
	} else {
		for i, v := range res {
			if v != expect[i] {
				t.Fail()
			}
		}
	}
	t.Logf("got: %v, expected: %v", res, expect)
}
