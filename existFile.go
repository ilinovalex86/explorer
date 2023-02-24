package explorer

import (
	"os"
)

// ExistFile - Проверяет существование файла
func ExistFile(filePath string) bool {
	path, err := os.Stat(filePath)
	if err != nil || path.IsDir() {
		return false
	}
	return true
}
