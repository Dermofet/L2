package main

import (
	"testing"
)

func TestUnpackString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{"", "", false},
		{"qwe\\4\\5", "qwe45", false},
		{"qwe\\45", "qwe44444", false},
		{"qwe\\\\5", "qwe\\\\\\\\\\", false},
	}

	for _, test := range tests {
		result, err := UnpackString(test.input)

		if test.err && err == nil {
			t.Errorf("Для входной строки '%s' ожидалась ошибка, но получен результат '%s'", test.input, result)
		}

		if !test.err && err != nil {
			t.Errorf("Для входной строки '%s' ожидался результат '%s', но получена ошибка: %v", test.input, test.expected, err)
		}

		if result != test.expected {
			t.Errorf("Для входной строки '%s' ожидался результат '%s', но получен результат '%s'", test.input, test.expected, result)
		}
	}
}
