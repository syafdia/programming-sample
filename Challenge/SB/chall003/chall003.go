package main

import (
	"strings"
)

// findFirstStringInBracketV1 is original code.
func findFirstStringInBracketV1(str string) string {
	if len(str) > 0 {
		indexFirstBracketFound := strings.Index(str, "(")
		if indexFirstBracketFound >= 0 {
			runes := []rune(str)
			wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
			if indexClosingBracketFound >= 0 {
				runes := []rune(wordsAfterFirstBracket)
				return string(runes[1 : indexClosingBracketFound-1])
			} else {
				return ""
			}
		} else {
			return ""
		}
	} else {
		return ""
	}

	return ""
}

// findFirstStringInBracketV2 is refactored code.
func findFirstStringInBracketV2(str string) string {
	if len(str) == 0 {
		return ""
	}

	indexFirstBracketFound := strings.Index(str, "(")
	if indexFirstBracketFound < 0 {
		return ""
	}

	wordsAfterFirstBracket := string(str[indexFirstBracketFound:])
	indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
	if indexClosingBracketFound < 0 {
		return ""
	}

	return string(wordsAfterFirstBracket[1 : indexClosingBracketFound-1])
}
