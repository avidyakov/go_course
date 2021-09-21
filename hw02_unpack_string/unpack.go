package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(rawString string) (result string, err error) {
	backslash := false
	theLastCharacterIsNumber := false
	for _, character := range rawString {
		if string(character) == "\\" && !backslash {
			backslash = true
			continue
		}
		if unicode.IsNumber(character) && !backslash {
			if theLastCharacterIsNumber {
				return "", ErrInvalidString
			}
			number, _ := strconv.Atoi(string(character))
			lastCharacterIndex := len(result) - 1
			if lastCharacterIndex < 0 {
				return "", ErrInvalidString
			}
			result = result[:lastCharacterIndex] + strings.Repeat(string(result[lastCharacterIndex]), number)
			theLastCharacterIsNumber = true
		} else {
			result += string(character)
			theLastCharacterIsNumber = false
		}
		backslash = false
	}
	return
}
