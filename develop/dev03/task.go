package main

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

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

type InputArgs struct {
	Filename       string
	Column         int
	Numeric        bool
	Reverse        bool
	Unique         bool
	MonthSort      bool
	IgnoreTrailing bool
	CheckSorted    bool
	HumanNumeric   bool
}

func main() {
	Sort(GetArgs())
}

func GetArgs() *InputArgs {
	// Определение флагов командной строки
	column := flag.Int("k", -1, "указание колонки для сортировки")
	numeric := flag.Bool("n", false, "сортировать по числовому значению")
	reverse := flag.Bool("r", false, "сортировать в обратном порядке")
	unique := flag.Bool("u", false, "не выводить повторяющиеся строки")
	monthSort := flag.Bool("M", false, "сортировка по названию месяца")
	ignoreTrailing := flag.Bool("b", false, "игнорировать хвостовые пробелы")
	checkSorted := flag.Bool("c", false, "проверить, отсортированы ли данные")
	humanNumeric := flag.Bool("h", false, "сортировать по числовому значению с учетом суффиксов")

	flag.Parse()

	// Проверка, что указан файл для обработки
	if flag.NArg() != 1 {
		fmt.Println("Usage: go run main.go [options] filename")
		flag.PrintDefaults()
		os.Exit(1)
	}

	return &InputArgs{
		Filename:       flag.Arg(0),
		Column:         *column,
		Numeric:        *numeric,
		Reverse:        *reverse,
		Unique:         *unique,
		MonthSort:      *monthSort,
		IgnoreTrailing: *ignoreTrailing,
		CheckSorted:    *checkSorted,
		HumanNumeric:   *humanNumeric,
	}
}

func Sort(args *InputArgs) {
	// Открытие файла на чтение
	file, err := os.Open(args.Filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Считывание строк из файла
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Сортировка строк в соответствии с заданными параметрами
	sort.SliceStable(lines, func(i, j int) bool {
		return compare(
			lines[i],
			lines[j],
			args.Column,
			args.Numeric,
			args.MonthSort,
			args.IgnoreTrailing,
			args.HumanNumeric,
			args.Reverse,
		)
	})

	// Удаление дубликатов, если указан флаг -u
	if args.Unique {
		lines = removeDuplicates(lines)
	}

	// Проверка, отсортированы ли данные
	if args.CheckSorted {
		if isSorted(
			lines,
			args.Column,
			args.Numeric,
			args.MonthSort,
			args.IgnoreTrailing,
			args.HumanNumeric,
			args.Reverse,
		) {
			fmt.Println("Data is sorted.")
		} else {
			fmt.Println("Data is not sorted.")
		}
	}

	// Открытие файла на запись
	outputFile, err := os.Create(args.Filename)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	// Запись отсортированных строк в файл
	writer := bufio.NewWriter(outputFile)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			os.Exit(1)
		}
	}
	if err := writer.Flush(); err != nil {
		fmt.Printf("Error flushing buffer: %v\n", err)
		os.Exit(1)
	}
}

// Сравнение двух строк в соответствии с заданными параметрами
func compare(a, b string, column int, numeric, monthSort, ignoreTrailing, humanNumeric, reverse bool) bool {
	if ignoreTrailing {
		a = strings.TrimSpace(a)
		b = strings.TrimSpace(b)
	}

	if monthSort {
		aMonth, errA := time.Parse("January", a)
		bMonth, errB := time.Parse("January", b)
		if errA == nil && errB == nil {
			return aMonth.Before(bMonth) != reverse
		}
	}

	if numeric {
		aNum, errA := strconv.ParseFloat(a, 64)
		bNum, errB := strconv.ParseFloat(b, 64)
		if errA == nil && errB == nil {
			return aNum < bNum != reverse
		}
	}

	if humanNumeric {
		aNum, errA := strconv.ParseFloat(humanToNumeric(a), 64)
		bNum, errB := strconv.ParseFloat(humanToNumeric(b), 64)
		if errA == nil && errB == nil {
			return aNum < bNum != reverse
		}
	}

	if column > -1 {
		aParts := strings.Fields(a)
		bParts := strings.Fields(b)

		if column < len(aParts) && column < len(bParts) {
			return aParts[column] < bParts[column] != reverse
		}
	}

	return a < b != reverse
}

// Преобразование строки с числовым суффиксом в числовое значение
func humanToNumeric(s string) string {
	suffixes := map[string]float64{
		"k":  1e3,
		"m":  1e6,
		"g":  1e9,
		"t":  1e12,
		"p":  1e15,
		"e":  1e18,
		"z":  1e21,
		"y":  1e24,
		"K":  1e3,
		"M":  1e6,
		"G":  1e9,
		"T":  1e12,
		"P":  1e15,
		"E":  1e18,
		"Z":  1e21,
		"Y":  1e24,
		"Ki": 1 << 10,
		"Mi": 1 << 20,
		"Gi": 1 << 30,
		"Ti": 1 << 40,
		"Pi": 1 << 50,
		"Ei": 1 << 60,
	}

	for suffix, value := range suffixes {
		if strings.HasSuffix(s, suffix) {
			return strconv.FormatFloat(value*parseHuman(s[:len(s)-len(suffix)]), 'f', -1, 64)
		}
	}

	return s
}

// Парсинг числового суффикса
func parseHuman(s string) float64 {
	if len(s) == 0 {
		return 0
	}

	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}

	return num
}

// Удаление дубликатов из среза строк
func removeDuplicates(lines []string) []string {
	seen := make(map[string]struct{})
	result := make([]string, 0)

	for _, line := range lines {
		if _, ok := seen[line]; !ok {
			seen[line] = struct{}{}
			result = append(result, line)
		}
	}

	return result
}

// Проверка, отсортированы ли строки в соответствии с заданными параметрами
func isSorted(lines []string, column int, numeric, monthSort, ignoreTrailing, humanNumeric, reverse bool) bool {
	for i := 1; i < len(lines); i++ {
		if compare(lines[i-1], lines[i], column, numeric, monthSort, ignoreTrailing, humanNumeric, reverse) {
			return false
		}
	}
	return true
}
