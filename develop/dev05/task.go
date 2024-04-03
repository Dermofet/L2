package main

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type InputArgs struct {
	After       int
	Before      int
	Context     int
	Count       bool
	IgnoreCase  bool
	Invert      bool
	Fixed       bool
	LineNumbers bool
	Pattern     string
	Filenames   []string
}

func main() {
	args := GetArgs()
	Grep(args)
}

func GetArgs() *InputArgs {
	after := flag.Int("A", 0, "Print +N lines after the match")
	before := flag.Int("B", 0, "Print +N lines before the match")
	context := flag.Int("C", 0, "Print ±N lines around the match")
	count := flag.Bool("c", false, "Count the number of lines")
	ignoreCase := flag.Bool("i", false, "Ignore case")
	invert := flag.Bool("v", false, "Invert the match")
	fixed := flag.Bool("F", false, "Fixed string match")
	lineNumbers := flag.Bool("n", false, "Print line numbers")
	flag.Parse()

	args := &InputArgs{
		After:       *after,
		Before:      *before,
		Context:     *context,
		Count:       *count,
		IgnoreCase:  *ignoreCase,
		Invert:      *invert,
		Fixed:       *fixed,
		LineNumbers: *lineNumbers,
		Pattern:     flag.Arg(0),
		Filenames:   flag.Args()[1:],
	}

	return args
}

func Grep(args *InputArgs) {
	for _, filename := range args.Filenames {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", filename, err)
			continue
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		lineNumber := 0
		printed := false

		for scanner.Scan() {
			line := scanner.Text()
			lineNumber++

			match := strings.Contains(line, args.Pattern)
			if args.IgnoreCase {
				match = strings.Contains(strings.ToLower(line), strings.ToLower(args.Pattern))
			}

			if args.Fixed {
				// println(line)
				// println(args.Pattern)
				// println(line == args.Pattern)
				match = line == args.Pattern
				if args.IgnoreCase {
					match = strings.EqualFold(line, args.Pattern)
				}
			}

			if args.Invert {
				match = !match
			}

			if match {
				printed = true
				PrintMatchedLine(args, filename, line, lineNumber)
			} else {
				printed = false
			}
		}

		if !printed && args.Count {
			fmt.Printf("0\n")
		}
	}
}

func PrintMatchedLine(args *InputArgs, filename, line string, lineNumber int) {
	if args.Count {
		return
	}

	if args.LineNumbers {
		fmt.Printf("%s:%d:", filename, lineNumber)
	}

	fmt.Println(line)
}
