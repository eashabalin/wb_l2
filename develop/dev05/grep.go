package dev05

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Grep is an object that helps to find lines by pattern
type Grep struct {
	options GrepOptions
	regex   *regexp.Regexp
}

func NewGrep(options GrepOptions) (*Grep, error) {
	if options.IgnoreCase {
		options.Pattern = strings.ToLower(options.Pattern)
	}
	regex, err := regexp.Compile(options.Pattern)
	if err != nil {
		return nil, errors.New("can't compile regexp")
	}
	return &Grep{options: options, regex: regex}, nil
}

func (g *Grep) Find(lines []string) []string {
	if g.options.Pattern == "" {
		return lines
	}
	res := make([]string, 0, len(lines))
	appended := make([]bool, len(lines))
	for i, l := range lines {
		if g.options.IgnoreCase {
			l = strings.ToLower(l)
		}
		match := false
		if g.options.Fixed {
			match = l == g.options.Pattern
		} else {
			match = g.regex.MatchString(l)
		}
		if match != g.options.Invert {
			insertWithContext(&res, &lines, i, g.options, &appended)
		}
	}
	if g.options.Count {
		return []string{strconv.Itoa(len(res))}
	}
	return res
}

func insertWithContext(res, lines *[]string, i int, options GrepOptions, appended *[]bool) {
	from := i
	to := i
	if options.Context > 0 {
		from = i - options.Context
		to = i + options.Context
	} else if options.Before > 0 {
		from = i - options.Before
		to = i
	} else if options.After > 0 {
		from = i
		to = i + options.After
	} else {
		from = i
		to = i
	}
	if from < 0 {
		from = 0
	}
	if to >= len(*lines) {
		to = len(*lines) - 1
	}
	for j := from; j <= to; j++ {
		if !(*appended)[j] {
			l := (*lines)[j]
			if options.LineNum {
				l = fmt.Sprintf("%d. %s", j+1, l)
			}
			*res = append(*res, l)
		}
		(*appended)[j] = true
	}
}
