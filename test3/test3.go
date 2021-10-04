package test3

import "strings"

func FindFirstStringInBracket(str string) string {
	if str == "" {
		return ""
	}

	indexFirstBracketFound := strings.Index(str, "(")
	if indexFirstBracketFound < 0 {
		return ""
	}

	indexClosingBracketFound := strings.Index(str[indexFirstBracketFound:], ")")
	if indexClosingBracketFound < 0 {
		return ""
	}

	wordsAfterFirstBracket := str[indexFirstBracketFound:]
	return string(wordsAfterFirstBracket[1:indexClosingBracketFound])
}
