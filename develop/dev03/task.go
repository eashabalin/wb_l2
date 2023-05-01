package dev03

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type SortOptions struct {
	FileName       string
	OutputFileName string
	Column         int
	Unique         bool
	Reverse        bool
	Numeric        bool
}

func Main() {
	options := parseArgs()

	lines, err := readLines(options.FileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sortLines(lines, options)

	if options.Unique {
		lines = removeDuplicate(lines)
	}

	if options.OutputFileName != "" {
		err = writeLines(lines, options.OutputFileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}

	for _, s := range lines {
		fmt.Println(s)
	}
}

func parseArgs() SortOptions {
	var options SortOptions

	flag.IntVar(&options.Column, "k", 0, "column index (starting from 1) to sort")
	flag.BoolVar(&options.Unique, "u", false, "remove duplicate lines")
	flag.BoolVar(&options.Reverse, "r", false, "reverse the sort order")
	flag.BoolVar(&options.Numeric, "n", false, "sort numerically")
	flag.StringVar(&options.OutputFileName, "o", "", "write output to file instead of stdout")
	flag.Parse()

	options.FileName = flag.Arg(0)

	return options
}

func readLines(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeLines(lines []string, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err = writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

func sortLines(lines []string, options SortOptions) {
	sort.Slice(lines, func(i, j int) bool {
		var col1 string
		var col2 string
		if options.Column > 0 {
			cols1 := strings.Fields(lines[i])
			cols2 := strings.Fields(lines[j])
			if len(cols1) < options.Column || len(cols2) < options.Column {
				return false
			}
			col1 = cols1[options.Column-1]
			col2 = cols2[options.Column-1]
		} else {
			col1 = lines[i]
			col2 = lines[j]
		}
		if options.Numeric {
			num1, err1 := strconv.ParseFloat(col1, 64)
			num2, err2 := strconv.ParseFloat(col2, 64)
			if err1 == nil && err2 == nil {
				if options.Reverse {
					return num1 > num2
				}
				return num1 < num2
			}
		}
		if col1 != col2 {
			if options.Reverse {
				return col1 > col2
			}
			return col1 < col2
		}
		return false
	})
}

func removeDuplicate(lines []string) []string {
	uniqueLines := make([]string, 0, len(lines))
	seen := make(map[string]bool)
	for _, line := range lines {
		if !seen[line] {
			seen[line] = true
			uniqueLines = append(uniqueLines, line)
		}
	}
	return uniqueLines
}
