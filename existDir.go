package explorer

import (
	"os"
)

// ExistDir - Проверяет существование папки
func ExistDir(folderPath string) bool {
	path, err := os.Stat(folderPath)
	if err != nil || !path.IsDir() {
		return false
	}
	return true
}
