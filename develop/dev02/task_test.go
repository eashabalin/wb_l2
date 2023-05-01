package dev02

import "testing"

func TestUnpackStr(t *testing.T) {
	// Тесты корректных строк
	cases := []struct {
		in, want string
	}{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"", ""},
		{"2", ""},
		{`\\` + `"`, `\`},
		{`a\5b`, "a5b"},
	}
	for _, c := range cases {
		got := UnpackStr(c.in)
		if got != c.want {
			t.Errorf("UnpackStr(%q) == %q, want %q", c.in, got, c.want)
		}
	}

	//Тест на строку с ошибкой
	invalidCase := "45"
	got := UnpackStr(invalidCase)
	if got != "" {
		t.Errorf("UnpackStr(%q) == %q, want \"\"", invalidCase, got)
	}
}
