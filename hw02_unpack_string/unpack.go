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

	length := len(runes)

	for i < length {
		char := runes[i]

		if i+1 < length && unicode.IsDigit(runes[i+1]) {
			if unicode.IsDigit(char) || (length > i+2 && unicode.IsDigit(runes[i+2])) {
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
	return string(char)
}
