package dev06

import (
	"strings"
)

type Cut struct {
	options CutOptions
}

func NewCut(options CutOptions) *Cut {
	return &Cut{options: options}
}

func (c *Cut) Cut(lines []string) [][]string {
	sep := "	"
	if c.options.Delimiter != "" {
		sep = c.options.Delimiter
	}

	table := make([][]string, 0, len(lines))
	res := make([][]string, 0, len(lines))

	counter := 0

	for i := 0; i < len(lines); i++ {
		pieces := strings.Split(lines[i], sep)
		if len(pieces) == 1 && c.options.Separated {
			continue
		}

		table = append(table, []string{})

		for _, v := range pieces {
			table[counter] = append(table[counter], v)
		}
		counter++
	}

	for i := 0; i < len(table); i++ {
		res = append(res, []string{})
		for _, f := range c.options.Fields {
			if len(table[i]) == 1 && len(res[i]) == 0 {
				res[i] = append(res[i], table[i][0])
			} else if len(table[i]) >= f {
				res[i] = append(res[i], table[i][f-1])
			}
		}
	}

	return res
}
