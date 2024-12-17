package utils

func StrInSlice(line string, list []string) bool {
	for _, v := range list {
		if line == v {
			return true
		}
	}
	return false
}
