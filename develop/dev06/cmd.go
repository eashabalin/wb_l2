package dev06

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	options, err := parseArgs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing args: %v\n", err)
		os.Exit(1)
	}

	c := NewCut(options.CutOptions)

	lines, err := readLines(options.InputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
		os.Exit(1)
	}

	res := c.Cut(lines)

	if options.OutputFile != "" {
		resLines := make([]string, 0, len(res))
		for i := 0; i < len(res); i++ {
			resLines = append(resLines, strings.Join(res[i], options.Delimiter))
		}
		err = writeLines(resLines, options.OutputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error writing file: %v\n", err)
			os.Exit(1)
		}
		return
	}

	printTable(res)
}

func printTable(table [][]string) {
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			fmt.Printf("%v ", table[i][j])
		}
		fmt.Println()
	}
}

type CmdCutOptions struct {
	InputFile  string
	OutputFile string
	CutOptions
}

func parseArgs() (*CmdCutOptions, error) {
	var options CmdCutOptions

	flag.StringVar(&options.OutputFile, "o", "", "Write output to file instead of stdout.")
	fieldsStr := flag.String("f", "", "Fields to cut out.")
	flag.BoolVar(&options.Separated, "s", false, "Include lines with separators only.")
	flag.StringVar(&options.Delimiter, "d", "	", "Custom delimiter.")
	flag.Parse()

	fields := strings.Split(*fieldsStr, ",")
	for _, v := range fields {
		f, err := strconv.Atoi(v)
		if f == 0 {
			return nil, errors.New("list: values may not include zero")
		}
		if err != nil {
			return nil, errors.New("list: illegal list value")
		}
		options.Fields = append(options.Fields, f)
	}

	options.InputFile = flag.Arg(0)

	return &options, nil
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
