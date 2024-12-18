package utils

import "strings"

// вырезание префикса для хекса: 0x/0X
func StripHexPrefix(s string) string {
	s = strings.ReplaceAll(s, "0x", "")
	s = strings.ReplaceAll(s, "0X", "")
	return s
}
