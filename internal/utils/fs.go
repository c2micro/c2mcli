package utils

import "path/filepath"

// получение абсолютного пути
func GetAbsPath(path string) (string, error) {
	return filepath.Abs(path)
}
