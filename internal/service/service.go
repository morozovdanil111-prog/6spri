package service

import (
	"strings"
	"errors"
	"github.com/yourusername/morse"
)

// AutoConvert определяет, является ли переданный текст кодом Морзе или обычным текстом.
// Возвращает преобразованный результат.
func AutoConvert(input string) (string, error) {
	// Убираем все пробелы в начале и в конце строки
	input = strings.TrimSpace(input)

	// Если строка состоит только из точек, тире и пробелов, вероятно, это код Морзе
	if isMorse(input) {
		return morse.ToText(input), nil
	}

	// Если строка содержит обычный текст, преобразуем её в код Морзе
	return morse.ToMorse(input), nil
}

// isMorse проверяет, является ли строка кодом Морзе
func isMorse(input string) bool {
	for _, c := range input {
		if !(c == '.' || c == '-' || c == ' ' || c == '\n') {
			return false
		}
	}
	return true
}