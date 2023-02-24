package explorer

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// ClearDir - Очищает папку
func ClearDir(path string) error {
	paths, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range paths {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			err = DelDir(filePath)
			if err != nil {
				return err
			}
		} else {
			err = os.Remove(filePath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
