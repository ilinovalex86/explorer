package explorer

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// ClearDir - Очищает папку
func ClearDir(path string) {
	paths, err := ioutil.ReadDir(path)
	check(err)
	for _, file := range paths {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			DelDir(filePath)
		} else {
			err = os.Remove(filePath)
			check(err)
		}
	}
}
