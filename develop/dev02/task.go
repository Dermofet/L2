package main

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// UnpackString осуществляет распаковку строки.
func UnpackString(s string) (string, error) {
	var result strings.Builder    // Создаем новый буфер строк для накопления результата
	var nextIsDigit bool          // Переменная для отслеживания наличия цифры в предыдущем символе
	var prevWasEscape bool = true // Переменная для отслеживания наличия обратной косой черты в предыдущем символе
	var prev rune                 // Переменная для хранения предыдущего символа (для обработки повторений)

	for i, r := range s {
		// Проверка на корректность входной строки
		if i == 0 && unicode.IsDigit(r) {
			return "", errors.New("некорректная строка")
		}

		// Проверка что следующая руна - цифра
		if r == '\\' {
			// Проверка на ввод '\\'
			if r == prev {
				nextIsDigit = false
				continue
			}

			nextIsDigit = true
			if !prevWasEscape {
				result.WriteRune(prev)
			}
			prev = r
			continue
		}

		// Проверка на ввод цифры
		if unicode.IsDigit(r) && nextIsDigit {
			prev = r
			nextIsDigit = false
			continue
		}

		// Вывод escape последовательности
		if unicode.IsDigit(r) && !nextIsDigit {
			count, _ := strconv.Atoi(string(r)) // Преобразуем символ цифры в число
			for j := 0; j < count; j++ {        // Повторяем предыдущий символ указанное количество раз
				result.WriteRune(prev)
			}
			prevWasEscape = true
			continue
		}

		// Вывод букв
		if unicode.IsLetter(r) {
			if !prevWasEscape {
				result.WriteRune(prev)
			}
			prev = r
			prevWasEscape = false
		} else {
			return "", errors.New("некорректная строка")
		}
	}

	// Вывод последнего символа
	if !prevWasEscape {
		result.WriteRune(prev)
	}

	return result.String(), nil // Возвращаем полученный результат в виде строки и nil в качестве ошибки
}

func main() {
	UnpackString("a4bc2d5e")
}
