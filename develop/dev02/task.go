package dev02

import (
	"errors"
	"strings"
	"unicode"
)

const (
	escapeCharacter = '\\'
)

type packedCharacter struct {
	ch          rune
	repeatTimes int
}

func UnpackString(source string) (string, error) {
	var packedChars []packedCharacter

	var isEscaped bool
	var last int

	for i, ch := range []rune(source) {

		if ch == escapeCharacter && !isEscaped {
			isEscaped = true
			continue
		}

		if unicode.IsDigit(ch) && !isEscaped {
			if i == 0 {
				return "", errors.New("некорректная строка")
			}
			packedChars[last].repeatTimes *= 10
			packedChars[last].repeatTimes += runeToDigit(ch)
			continue
		}

		pChar := packedCharacter{ch, 0}
		packedChars = append(packedChars, pChar)
		last = len(packedChars) - 1

		isEscaped = false
	}

	result := unpackCharacters(packedChars)

	return result, nil
}

func unpackCharacters(chars []packedCharacter) string {
	builder := strings.Builder{}

	for _, pChar := range chars {
		if pChar.repeatTimes > 0 {
			builder.WriteString(strings.Repeat(string(pChar.ch), pChar.repeatTimes))
		} else {
			builder.WriteRune(pChar.ch)
		}
	}

	return builder.String()
}

func runeToDigit(r rune) int {
	return int(r - '0')
}
