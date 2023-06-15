package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}

	if unicode.IsDigit(rune(s[0])) {
		return "", ErrInvalidString
	}

	var result strings.Builder
	runes := []rune(s)
	i := 0

	for i < len(runes) {
		char := runes[i]
		if i+1 < len(runes) && unicode.IsDigit(runes[i+1]) {
			if unicode.IsDigit(char) || unicode.IsDigit(runes[i+2]) {
				return "", ErrInvalidString
			}
			count, _ := strconv.Atoi(string(runes[i+1]))
			if count != 0 {
				result.WriteString(strings.Repeat(getRepeatCharacter(char), count))
			}
			i += 2
		} else {
			result.WriteString(getRepeatCharacter(char))
			i++
		}
	}

	return result.String(), nil
}

func getRepeatCharacter(char rune) string {
	if char == '\n' {
		return "\n"
	}
	return string(char)
}
