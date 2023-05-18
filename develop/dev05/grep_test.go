package dev05

import (
	"testing"
)

func TestGrep_Find(t *testing.T) {
	in := []string{"      41.14KK", "fafaf", "254u", "OOOOO", "   .MMMk1", "fafaf"}
	expect := []string{"1.       41.14KK", "3. 254u"}
	options := GrepOptions{Pattern: "3234refweg"}
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
