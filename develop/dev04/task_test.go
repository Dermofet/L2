package main

import (
	"reflect"
	"testing"
)

func TestFindAnagramSets(t *testing.T) {
	tests := []struct {
		words    []string
		expected [][]string
	}{
		{
			words: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			expected: [][]string{
				{"пятак", "пятка", "тяпка"},
				{"листок", "слиток", "столик"},
			},
		},
		{
			words: []string{"Стол", "Лотс", "соЛт", "тСол"},
			expected: [][]string{
				{"лотс", "солт", "стол", "тсол"},
			},
		},
		{
			words: []string{"бар", "раб", "бра"},
			expected: [][]string{
				{"бар", "бра", "раб"},
			},
		},
	}

	for _, test := range tests {
		result := findAnagramSets(&test.words)

		// Проверка соответствия ожидаемого и фактического результата
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Для входных данных %v ожидается результат %v, но получено %v", test.words, test.expected, result)
		}
	}
}
