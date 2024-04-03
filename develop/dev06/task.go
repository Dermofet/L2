package main

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type InputArgs struct {
	Fields    string
	Delimiter string
	Separated bool
}

func main() {
	args := GetArgs()
	Cut(os.Stdin, os.Stdout, args)
}

func GetArgs() *InputArgs {
	fields := flag.String("f", "", "Select fields")
	delimiter := flag.String("d", "\t", "Use a different delimiter")
	separated := flag.Bool("s", false, "Only output lines containing delimiter")

	flag.Parse()

	return &InputArgs{
		Fields:    *fields,
		Delimiter: *delimiter,
		Separated: *separated,
	}
}

func Cut(in io.Reader, out io.Writer, args *InputArgs) {
	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		line := scanner.Text()

		if !args.Separated || strings.Contains(line, args.Delimiter) {
			fields := strings.Split(line, args.Delimiter)
			selectedFields := make([]string, 0)

			for _, field := range strings.Split(args.Fields, ",") {
				index := parseFieldIndex(field)

				if index >= 0 && index < len(fields) {
					selectedFields = append(selectedFields, fields[index])
				}
			}

			fmt.Fprintf(out, "%s\n", strings.Join(selectedFields, args.Delimiter))
		}
	}
}

func parseFieldIndex(field string) int {
	index := -1
	fmt.Sscanf(field, "%d", &index)
	return index - 1
}
