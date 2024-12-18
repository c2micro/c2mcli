package utils

func BoolToFloat(b bool) float64 {
	if b {
		return 1.0
	}
	return 0.0
}

func BoolToInt(b bool) int64 {
	if b {
		return 1
	}
	return 0
}

func FloatToBool(f float64) bool {
	if f == 0.0 {
		return false
	}
	return true
}

func IntToFloat(i int64) float64 {
	return float64(i)
}

func IntToBool(i int64) bool {
	if i == 0 {
		return false
	}
	return true
}

func FloatToInt(f float64) int64 {
	return int64(f)
}
