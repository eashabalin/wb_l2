package dev03

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func Main() {
	options := parseArgs()

	err := options.Validate()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lines, err := readLines(options.FileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if options.IgnoreLeadingBlanks {
		removeLeadingBlanks(lines)
	}

	if options.Check {
		isSorted := check(lines, options)
		if isSorted {
			fmt.Println("sorted")
			return
		}
		fmt.Println("unsorted")
		return
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

	flag.IntVar(&options.Column, "k", 0, "Column index (starting from 1) to sort.")
	flag.BoolVar(&options.Unique, "u", false, "Remove duplicate lines.")
	flag.BoolVar(&options.Reverse, "r", false, "Reverse the sort order.")
	flag.BoolVar(&options.Numeric, "n", false, "Sort numerically.")
	flag.BoolVar(&options.IgnoreLeadingBlanks, "b", false, "Ignore leading blanks.")
	flag.BoolVar(&options.Check, "c", false, "Check if input is sorted.")
	flag.BoolVar(&options.SISuffix, "h", false, "Sort numerically with SI suffixes. For example: 10K < 10M < 10T. All supported suffixes: \"EPTGMK/kcmunpfa\" (descending order).")
	flag.BoolVar(&options.Month, "M", false, "Sort by month abbreviations. Unknown strings are considered bigger than the month names.")
	flag.StringVar(&options.OutputFileName, "o", "", "Write output to file instead of stdout.")
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
			if err1 == nil {
				return true
			}
			return false
		}
		if options.Month {
			m1, err1 := strToMonth(col1)
			m2, err2 := strToMonth(col2)
			if err1 == nil && err2 == nil {
				if options.Reverse {
					return m1 > m2
				}
				return m1 < m2
			}
			return false
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

func strToMonth(s string) (time.Month, error) {
	layout := "Jan"
	t, err := time.Parse(layout, s)
	if err != nil {
		return 0, err
	}
	return t.Month(), nil
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

func removeLeadingBlanks(lines []string) {
	for i := range lines {
		lines[i] = strings.TrimLeft(lines[i], " ")
	}
}

func check(lines []string, options SortOptions) bool {
	before := make([]string, len(lines))
	copy(before, lines)
	sortLines(lines, options)
	if len(before) != len(lines) {
		return false
	}
	for i, v := range before {
		if v != lines[i] {
			return false
		}
	}
	return true
}
