package explorer

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// DelDir - Удаляет папку
func DelDir(path string) {
	paths, err := ioutil.ReadDir(path)
	check(err)
	for _, file := range paths {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			DelDir(filePath)
		} else {
			err := os.Remove(filePath)
			check(err)
		}
	}
	err = os.Remove(path)
	check(err)
}
