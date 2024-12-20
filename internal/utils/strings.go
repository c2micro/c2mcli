package utils

import (
	"fmt"
	"unicode"
)

func StrInSlice(line string, list []string) bool {
	for _, v := range list {
		if line == v {
			return true
		}
	}
	return false
}

// TODO
func IsAsciiPrintable(s string) bool {
	for _, r := range s {
		if (r > unicode.MaxASCII || !unicode.IsPrint(r)) && r != '\n' {
			fmt.Println(r)
			return false
		}
	}
	return true
}
