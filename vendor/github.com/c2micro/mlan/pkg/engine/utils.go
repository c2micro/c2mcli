package engine

import (
	"strings"
)

// stripHexPrefix вырезание префикса для хекса: 0x/0X
func stripHexPrefix(s string) string {
	s = strings.ReplaceAll(s, "0x", "")
	s = strings.ReplaceAll(s, "0X", "")
	return s
}

func floatToInt(f float64) int64 {
	return int64(f)
}

// boolToFloat bool -> float64
func boolToFloat(b bool) float64 {
	if b {
		return 1.0
	} else {
		return 0.0
	}
}

// boolToInt bool -> int64
func boolToInt(b bool) int64 {
	if b {
		return 1
	} else {
		return 0
	}
}

// intToFloat int64 -> float64
func intToFloat(i int64) float64 {
	return float64(i)
}
