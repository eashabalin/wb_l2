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

// TODO: вынести всё из Find() в отдельные функции

func (g *Grep) Find(lines []string) []string {
	if g.options.Pattern == "" {
		return lines
	}
	res := make([]string, 0, len(lines))
	for i, l := range lines {
		originalLine := l
		if g.options.IgnoreCase {
			l = strings.ToLower(l)
		}
		if g.options.Fixed {
			if l == g.options.Pattern {
				if g.options.LineNum {
					l = fmt.Sprintf("%d: %s", i+1, l)
				}
				res = append(res, originalLine)
			}
		} else if g.regex.MatchString(l) {
			if g.options.LineNum {
				l = fmt.Sprintf("%d: %s", i+1, l)
			}
			res = append(res, originalLine)
		}
	}
	if g.options.Count {
		return []string{strconv.Itoa(len(res))}
	}
	return res
}
