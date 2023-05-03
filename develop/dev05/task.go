package dev05

import (
	"bufio"
	"flag"
	"os"
)

func Main() {

}

func parseArgs() GrepOptions {
	var options GrepOptions

	//flag.IntVar(&options.Column, "k", 0, "Column index (starting from 1) to sort.")
	//flag.BoolVar(&options.Unique, "u", false, "Remove duplicate lines.")
	//flag.BoolVar(&options.Reverse, "r", false, "Reverse the sort order.")
	//flag.BoolVar(&options.Numeric, "n", false, "Sort numerically.")
	//flag.BoolVar(&options.IgnoreLeadingBlanks, "b", false, "Ignore leading blanks.")
	//flag.BoolVar(&options.Check, "c", false, "Check if input is sorted.")
	//flag.BoolVar(&options.SISuffix, "h", false, "Sort numerically with SI suffixes. For example: 10K < 10M < 10T. All supported suffixes: \"EPTGMKcmunpfa\" (descending order).")
	//flag.BoolVar(&options.Month, "M", false, "Sort by month abbreviations. Unknown strings are considered bigger than the month names.")
	//flag.StringVar(&options.OutputFileName, "o", "", "Write output to file instead of stdout.")
	flag.Parse()

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
