package object

func boolToFloat(b bool) float64 {
	if b {
		return 1.0
	}
	return 0.0
}

func boolToInt(b bool) int64 {
	if b {
		return 1
	}
	return 0
}

func floatToBool(f float64) bool {
	if f == 0.0 {
		return false
	}
	return true
}

func intToFloat(i int64) float64 {
	return float64(i)
}

func intToBool(i int64) bool {
	if i == 0 {
		return false
	}
	return true
}