package utils

import (
	"unicode"
)

var (
	ranger = []*unicode.RangeTable{unicode.Latin, unicode.Cyrillic, unicode.ASCII_Hex_Digit, unicode.Punct, unicode.White_Space}
)

func StrInSlice(line string, list []string) bool {
	for _, v := range list {
		if line == v {
			return true
		}
	}
	return false
}

func IsAsciiPrintable(s string) bool {
	if len(s) > 1024 {
		s = s[:1024]
	}
	for _, r := range []rune(s) {
		if !unicode.IsOneOf(ranger, r) {
			return false
		}
	}
	return true
}
