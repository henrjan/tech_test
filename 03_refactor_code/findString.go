package main

import "strings"

func findFirstStringInBracket(str string) string {
	if len(str) == 0 {
		return ""
	}

	indexFirstBracketFound := strings.Index(str, "(")
	if indexFirstBracketFound == -1 {
		return ""
	}

	wordsAfterFirstBracket := str[indexFirstBracketFound+1:]
	indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
	if indexClosingBracketFound == -1 {
		return ""
	}

	return wordsAfterFirstBracket[:indexClosingBracketFound]
}
