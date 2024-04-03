package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

// Тест для проверки сортировки строк в обычном порядке без дополнительных параметров
func TestSortStrings(t *testing.T) {
	input := "c\na\nb"
	expected := "a\nb\nc"

	// Создание временного файла с тестовыми данными
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	// Запись тестовых данных во временный файл
	if _, err := tmpfile.Write([]byte(input)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	args := &InputArgs{
		Filename:       tmpfile.Name(),
		Column:         1,
		Numeric:        false,
		Reverse:        false,
		Unique:         false,
		MonthSort:      false,
		IgnoreTrailing: false,
		CheckSorted:    false,
		HumanNumeric:   false,
	}

	// Запуск программы с сортировкой
	Sort(args)

	// Считывание отсортированных строк из временного файла
	file, err := os.Open(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	// Сравнение отсортированных строк с ожидаемым результатом
	actual := strings.Join(result, "\n")
	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}

// Тест для проверки сортировки строк в обратном порядке с использованием ключа -r
func TestSortStringsReverse(t *testing.T) {
	input := "c\na\nb"
	expected := "c\nb\na"

	// Создание временного файла с тестовыми данными
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	// Запись тестовых данных во временный файл
	if _, err := tmpfile.Write([]byte(input)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	args := &InputArgs{
		Filename:       tmpfile.Name(),
		Column:         0,
		Numeric:        false,
		Reverse:        true,
		Unique:         false,
		MonthSort:      false,
		IgnoreTrailing: false,
		CheckSorted:    false,
		HumanNumeric:   false,
	}

	// Запуск программы с сортировкой
	Sort(args)

	// Считывание отсортированных строк из временного файла
	file, err := os.Open(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	// Сравнение отсортированных строк с ожидаемым результатом
	actual := strings.Join(result, "\n")
	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}

// Тест для проверки сортировки строк по числовому значению с использованием ключа -n
func TestSortStringsNumeric(t *testing.T) {
	input := "10\n2\n1"
	expected := "1\n2\n10"

	// Создание временного файла с тестовыми данными
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	// Запись тестовых данных во временный файл
	if _, err := tmpfile.Write([]byte(input)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	args := &InputArgs{
		Filename:       tmpfile.Name(),
		Column:         0,
		Numeric:        true,
		Reverse:        false,
		Unique:         false,
		MonthSort:      false,
		IgnoreTrailing: false,
		CheckSorted:    false,
		HumanNumeric:   false,
	}

	// Запуск программы с сортировкой по числовому значению с ключом -n
	Sort(args)

	// Считывание отсортированных строк из временного файла
	file, err := os.Open(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	// Сравнение отсортированных строк с ожидаемым результатом
	actual := strings.Join(result, "\n")
	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}

// Тест для проверки сортировки строк по названию месяца с использованием ключа -M
func TestSortStringsByMonth(t *testing.T) {
	input := "February\nJanuary\nMarch"
	expected := "January\nFebruary\nMarch"

	// Создание временного файла с тестовыми данными
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	// Запись тестовых данных во временный файл
	if _, err := tmpfile.Write([]byte(input)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	args := &InputArgs{
		Filename:       tmpfile.Name(),
		Column:         0,
		Numeric:        false,
		Reverse:        false,
		Unique:         false,
		MonthSort:      true,
		IgnoreTrailing: false,
		CheckSorted:    false,
		HumanNumeric:   false,
	}

	// Запуск программы с сортировкой по названию месяца с ключом -M
	Sort(args)

	// Считывание отсортированных строк из временного файла
	file, err := os.Open(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	// Сравнение отсортированных строк с ожидаемым результатом
	actual := strings.Join(result, "\n")
	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}

// Тест для проверки игнорирования хвостовых пробелов с использованием ключа -b
func TestIgnoreTrailingSpaces(t *testing.T) {
	input := "c\n a \nb"
	expected := " a \nb\nc"

	// Создание временного файла с тестовыми данными
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	// Запись тестовых данных во временный файл
	if _, err := tmpfile.Write([]byte(input)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	args := &InputArgs{
		Filename:       tmpfile.Name(),
		Column:         0,
		Numeric:        false,
		Reverse:        false,
		Unique:         false,
		MonthSort:      false,
		IgnoreTrailing: true,
		CheckSorted:    false,
		HumanNumeric:   false,
	}

	// Запуск программы с игнорированием хвостовых пробелов с ключом -b
	Sort(args)

	// Считывание отсортированных строк из временного файла
	file, err := os.Open(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	// Сравнение отсортированных строк с ожидаемым результатом
	actual := strings.Join(result, "\n")
	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}

// Тест для проверки сортировки строк с учетом числовых суффиксов с использованием ключа -h
func TestSortStringsWithSuffixes(t *testing.T) {
	input := "2M\n1K\n3G"
	expected := "1K\n2M\n3G"

	// Создание временного файла с тестовыми данными
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	// Запись тестовых данных во временный файл
	if _, err := tmpfile.Write([]byte(input)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	args := &InputArgs{
		Filename:       tmpfile.Name(),
		Column:         0,
		Numeric:        false,
		Reverse:        false,
		Unique:         false,
		MonthSort:      false,
		IgnoreTrailing: false,
		CheckSorted:    false,
		HumanNumeric:   true,
	}

	// Запуск программы с сортировкой с учетом числовых суффиксов с ключом -h
	Sort(args)

	// Считывание отсортированных строк из временного файла
	file, err := os.Open(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	// Сравнение отсортированных строк с ожидаемым результатом
	actual := strings.Join(result, "\n")
	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}

// Тест для проверки удаления дубликатов строк с использованием ключа -u
func TestRemoveDuplicates(t *testing.T) {
	input := "a\nb\na"
	expected := "a\nb"

	// Создание временного файла с тестовыми данными
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	// Запись тестовых данных во временный файл
	if _, err := tmpfile.Write([]byte(input)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	args := &InputArgs{
		Filename:       tmpfile.Name(),
		Column:         0,
		Numeric:        false,
		Reverse:        false,
		Unique:         true,
		MonthSort:      false,
		IgnoreTrailing: false,
		CheckSorted:    false,
		HumanNumeric:   false,
	}

	Sort(args)

	// Считывание отсортированных строк из временного файла
	file, err := os.Open(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	// Сравнение отсортированных строк с ожидаемым результатом
	actual := strings.Join(result, "\n")
	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}

// Тест для проверки сортировки с использованием стандартной функции сравнения строк
func TestSortWithoutArgs(t *testing.T) {
	input := "c\na\nb"
	expected := "a\nb\nc"

	// Создание временного файла с тестовыми данными
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	// Запись тестовых данных во временный файл
	if _, err := tmpfile.Write([]byte(input)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	args := &InputArgs{
		Filename:       tmpfile.Name(),
		Column:         -1,
		Numeric:        false,
		Reverse:        false,
		Unique:         false,
		MonthSort:      false,
		IgnoreTrailing: false,
		CheckSorted:    false,
		HumanNumeric:   false,
	}

	// Запуск программы с сортировкой
	Sort(args)

	// Считывание отсортированных строк из временного файла
	file, err := os.Open(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	// Сравнение отсортированных строк с ожидаемым результатом
	actual := strings.Join(result, "\n")
	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}
