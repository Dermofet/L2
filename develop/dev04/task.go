package main

import (
	"sort"
	"strings"
)

// Функция для поиска всех множеств анаграмм по словарю
func findAnagramSets(words *[]string) [][]string {
	anagramSets := make(map[string][]string)

	// Проход по каждому слову в словаре
	for _, word := range *words {
		// Приведение слова к нижнему регистру и сортировка его букв
		word = strings.ToLower(word)
		sortedWord := sortWord(word)

		// Добавление слова в соответствующее множество анаграмм
		anagramSets[sortedWord] = append(anagramSets[sortedWord], word)
	}

	// Удаление множеств из одного элемента и сортировка элементов в множествах
	var result [][]string
	for _, value := range anagramSets {
		if len(value) > 1 {
			sort.Strings(value)
			result = append(result, value)
		}
	}

	return result
}

// Функция для сортировки букв в слове
func sortWord(word string) string {
	letters := strings.Split(word, "")
	sort.Strings(letters)
	return strings.Join(letters, "")
}

func main() {
	// Пример использования функции
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	anagramSets := findAnagramSets(&words)
	for _, set := range anagramSets {
		println("Множество анаграмм:", strings.Join(set, ", "))
	}
}
