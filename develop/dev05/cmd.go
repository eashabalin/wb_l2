package dev05

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type CmdGrepOptions struct {
	InputFile  string
	OutputFile string
	GrepOptions
}

func Run() {
	options := parseArgs()

	lines, err := readLines(options.InputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	grep, err := NewGrep(options.GrepOptions)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res := grep.Find(lines)

	if options.OutputFile != "" {
		err = writeLines(res, options.OutputFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}

	for _, v := range res {
		fmt.Println(v)
	}
}

func parseArgs() CmdGrepOptions {
	var options CmdGrepOptions

	flag.StringVar(&options.OutputFile, "o", "", "Write output to file instead of stdout.")
	flag.StringVar(&options.Pattern, "p", "", "Regular expression or exact text pattern")
	flag.BoolVar(&options.Count, "c", false, "Print count of matches.")
	flag.BoolVar(&options.IgnoreCase, "i", false, "Ignore text case.")
	flag.BoolVar(&options.Invert, "v", false, "Invert output. Includes all lines that don't match pattern.")
	flag.BoolVar(&options.Fixed, "F", false, "Include exact matches with string.")
	flag.BoolVar(&options.LineNum, "n", false, "Print line number with \". \" separator. "+
		"For example for line \"I am the line\": \"1. I am the line\".")
	flag.IntVar(&options.After, "A", 0, "Print N strings after.")
	flag.IntVar(&options.Before, "B", 0, "Print N strings before.")
	flag.IntVar(&options.Context, "C", 0, "Print N strings before and after.")
	flag.Parse()

	options.InputFile = flag.Arg(0)

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
