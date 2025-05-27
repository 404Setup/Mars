package utils

import (
	"strings"
	"unicode"
)

func IsValidProjectName(name string) bool {
	return strings.IndexFunc(name, func(r rune) bool {
		isLetter := unicode.IsLetter(r)
		isDigit := unicode.IsDigit(r)
		isExtended := r > 127
		return !(isLetter || isDigit || isExtended)
	}) == -1
}
